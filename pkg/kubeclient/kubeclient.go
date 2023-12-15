package kubeclient

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Mirantis/mcc/pkg/constant"
	log "github.com/sirupsen/logrus"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeClient struct {
	Namespace string

	client         kubernetes.Interface
	extendedClient apiextensionsclientset.Interface
	config         *rest.Config
}

// NewFromBundle returns a new instance of KubeClient from
// a given bundle directory.
func NewFromBundle(bundleDir, namespace string) (*KubeClient, error) {
	f := filepath.Join(bundleDir, constant.KubeConfigFile)

	configBytes, err := os.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read %q: %w", f, err)
	}

	config, err := clientcmd.RESTConfigFromKubeConfig(configBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse kubeconfig: %w", err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("could not initialize kubernetes client: %w", err)
	}

	extendedClientSet, err := apiextensionsclientset.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize apiextensions clientset: %w", err)
	}

	kc := &KubeClient{
		Namespace:      namespace,
		client:         clientSet,
		extendedClient: extendedClientSet,
		config:         config,
	}

	return kc, nil
}

// CRDReady verifies that the CRD and Deployment is available.
// This is the equivalent of running `kubectl get crd crdName`.
func (kc *KubeClient) crdReady(ctx context.Context, name string) error {
	_, err := kc.extendedClient.ApiextensionsV1().CustomResourceDefinitions().Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get %q crd: %w", name, err)
	}

	return nil
}

// deploymentReady verifies that the Deployment is available.
// This is the equivalent of running `kubectl list deployment -l labels`.
func (kc *KubeClient) deploymentReady(ctx context.Context, labels string) error {
	d, err := kc.client.AppsV1().Deployments(kc.Namespace).List(ctx, metav1.ListOptions{
		LabelSelector: labels,
	})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return fmt.Errorf("Deployment with %q labels not found, ensure the deployment exists", labels)
		}

		return err
	}

	if len(d.Items) < 1 {
		return fmt.Errorf("Deployment with %q labels not found, ensure the deployment exists", labels)
	}

	if d.Items[0].Status.ReadyReplicas < 1 {
		return fmt.Errorf("Deployment with %q labels was found, but is not yet ready", labels)
	}

	return nil
}

func (kc *KubeClient) crIsReady(ctx context.Context, obj *unstructured.Unstructured, resourceClient dynamic.ResourceInterface) (bool, error) {
	crdObj, err := resourceClient.Get(ctx, obj.GetName(), metav1.GetOptions{})
	if err != nil {
		return false, fmt.Errorf("failed to get resource: %q: %w", obj.GetName(), err)
	}

	// Check readiness condition
	conditions, found, err := unstructured.NestedSlice(crdObj.Object, "status", "conditions")
	if err != nil || !found {
		return false, fmt.Errorf("cannot find status condition: %w", err)
	}

	for _, cond := range conditions {
		condMap, ok := cond.(map[string]interface{})
		if !ok {
			continue
		}

		if condType, ok := condMap["type"].(string); ok && condType == "Ready" {
			if readyStatus, ok := condMap["status"].(string); ok && readyStatus == "True" {
				return true, nil
			}
		}
	}

	return false, nil
}

// SetStorageClassDefault configures the given storageclass name as the default,
// ensuring that no other storageclass has the default annotation.
func (kc *KubeClient) SetStorageClassDefault(ctx context.Context, name string) error {
	log.Debugf("setting: %s as default StorageClass", name)

	// Ensure no other StorageClass is already set to default.
	storageClasses, err := kc.client.StorageV1().StorageClasses().List(ctx, metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list StorageClasses: %w", err)
	}

	var needsUpdate bool

	for _, sc := range storageClasses.Items {
		if sc.Name == name {
			// Apply the default annotation to the named StorageClass.
			if _, ok := sc.ObjectMeta.Annotations[constant.DefaultStorageClassAnnotation]; !ok {
				log.Debugf("setting default StorageClass annotation on: %s", name)

				sc.ObjectMeta.Annotations[constant.DefaultStorageClassAnnotation] = "true"
			}
			needsUpdate = true
		} else {
			// Strip the default annotation if found from a different StorageClass.
			if _, ok := sc.ObjectMeta.Annotations[constant.DefaultStorageClassAnnotation]; ok {
				log.Debugf("found existing default StorageClass: %s, removing default annotation", sc.Name)

				delete(sc.ObjectMeta.Annotations, constant.DefaultStorageClassAnnotation)
				needsUpdate = true
			}
		}

		if needsUpdate {
			// Apply the annotation modifications if they were made.
			if _, err := kc.client.StorageV1().StorageClasses().Update(ctx, &sc, metav1.UpdateOptions{}); err != nil {
				return fmt.Errorf("failed to update StorageClass: %q annotations: %w", name, err)
			}
		}
	}

	return nil
}

// DeleteService deletes service by name.
func (kc *KubeClient) DeleteService(ctx context.Context, name string) error {
	return kc.client.CoreV1().Services(kc.Namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

// ExposeLoadBalancer creates a new service of Type: LoadBalancer, its
// the equivalent of 'kubectl expose'.
func (kc *KubeClient) ExposeLoadBalancer(ctx context.Context, url string) error {
	return nil
}
