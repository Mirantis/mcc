package kubeclient

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Mirantis/mcc/pkg/constant"
	"github.com/docker/dhe-deploy/gocode/pkg/pollutil"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

func (kc *KubeClient) GetMSRCR(ctx context.Context, name string) (*unstructured.Unstructured, error) {
	rc, err := kc.getMSRResourceClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get resource client for MSR CR: %w", err)
	}

	return rc.Get(ctx, name, metav1.GetOptions{})
}

func (kc *KubeClient) ValidateMSROperatorReady(ctx context.Context) error {
	if err := kc.crdReady(ctx, "msrs.msr.mirantis.com"); err != nil {
		return err
	}

	return kc.deploymentReady(ctx, constant.MSROperatorDeploymentLabels)
}

// WaitForMSRCRReady waits for CR object provided to be ready by polling the
// status obtained from the given object.
func (kc *KubeClient) WaitForMSRCRReady(ctx context.Context, obj *unstructured.Unstructured) error {
	pollCfg := pollutil.DefaultPollfConfig(log.InfoLevel, "waiting for %q CR Ready state for up to 10m0s", obj.GetName())

	// Wait for a maximum time of 10 minutes.
	pollCfg.Interval = 5 * time.Second
	pollCfg.NumRetries = 120

	rc, err := kc.getMSRResourceClient()
	if err != nil {
		return err
	}

	err = pollutil.Pollf(pollCfg)(func() error {
		ready, e := kc.crIsReady(ctx, obj, rc)
		if e != nil {
			return fmt.Errorf("failed to process MSR CR: %w", e)
		}
		if !ready {
			return errors.New("MSR CR is not ready")
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to obtain MSR CR Ready state after %d retries: %w", pollCfg.NumRetries, err)
	}

	return nil
}

// ApplyMSRCR applies the given MSR CR object to the cluster, reattempting
// the operation ever if it fails every 5 seconds for up to 30 seconds.
func (kc *KubeClient) ApplyMSRCR(ctx context.Context, obj *unstructured.Unstructured) error {
	name := obj.GetName()

	existingObj, err := kc.GetMSRCR(ctx, name)
	if err != nil {
		if apierrors.IsNotFound(err) {
			log.Infof("MSR CR %q not found, creating", name)
		} else {
			return fmt.Errorf("failed to get MSR CR: %w", err)
		}
	}

	rc, err := kc.getMSRResourceClient()
	if err != nil {
		return pollutil.Abort(fmt.Errorf("failed to get resource client for MSR CR: %w", err))
	}

	pollCfg := pollutil.DefaultPollfConfig(log.InfoLevel, "Applying resource YAML")
	pollCfg.Interval = 5 * time.Second
	pollCfg.NumRetries = 6

	err = pollutil.Pollf(pollCfg)(func() error {
		if existingObj == nil {
			log.Debugf("msr resource: %q does not yet exist, creating", name)

			_, err = rc.Create(ctx, obj, metav1.CreateOptions{})
			if err != nil {
				return err
			}
		} else {
			// Set the resource version to the existing object's resource version
			// if it already exists to ensure that the update succeeds.
			obj.SetResourceVersion(existingObj.GetResourceVersion())

			log.Debugf("msr resource: %q exists, updating", name)

			_, err = rc.Update(ctx, obj, metav1.UpdateOptions{})
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to apply resource YAML after %d retries: %w", pollCfg.NumRetries, err)
	}

	return nil
}

func (kc *KubeClient) DeleteMSRCR(ctx context.Context, name string) error {
	rc, err := kc.getMSRResourceClient()
	if err != nil {
		return fmt.Errorf("failed to get resource client for MSR CR: %w", err)
	}

	return rc.Delete(ctx, name, metav1.DeleteOptions{})
}

// PrepareNodeForMSR updates the given node name setting the MSRNodeSelector
// on the node and removing any found Kubernetes NoExecute taints added by MKE.
func (kc *KubeClient) PrepareNodeForMSR(ctx context.Context, name string) error {
	node, err := kc.client.CoreV1().Nodes().Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get node %q: %w", name, err)
	}

	node.Labels[constant.MSRNodeSelector] = "true"

	// Rebuild the taints list without the NoExecute taint if found, removing it.
	var taints []corev1.Taint
	for _, t := range node.Spec.Taints {
		if t.Key == constant.KubernetesOrchestratorTaint && t.Value == "NoExecute" {
			continue
		}
		taints = append(taints, t)
	}

	node.Spec.Taints = taints

	_, err = kc.client.CoreV1().Nodes().Update(ctx, node, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update node %q: %w", name, err)
	}

	return nil
}

// getMSRResourceClient returns a dynamic client for the MSR custom resource.
func (kc *KubeClient) getMSRResourceClient() (dynamic.ResourceInterface, error) {
	client, err := dynamic.NewForConfig(kc.config)
	if err != nil {
		return nil, fmt.Errorf("failed to create dynamic client: %w", err)
	}

	return client.Resource(schema.GroupVersionResource{
		Group:    "msr.mirantis.com",
		Version:  "v1",
		Resource: "msrs",
	}).Namespace(kc.Namespace), nil
}
