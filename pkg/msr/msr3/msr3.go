package msr3

import (
	"context"
	"fmt"
	"time"

	"github.com/Mirantis/mcc/pkg/constant"
	"github.com/Mirantis/mcc/pkg/helm"
	"github.com/Mirantis/mcc/pkg/kubeclient"
	"github.com/Mirantis/mcc/pkg/product/mke/api"
	msrv1 "github.com/Mirantis/msr-operator/api/v1"
	"github.com/mitchellh/mapstructure"
	"helm.sh/helm/v3/pkg/release"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func CollectFacts(ctx context.Context, msrName string, kc *kubeclient.KubeClient, h *helm.Helm) (*api.MSRMetadata, error) {
	obj, err := kc.GetMSRCR(ctx, msrName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return &api.MSRMetadata{Installed: false}, nil
		}
		return nil, err
	}

	if err := kc.WaitForMSRCRReady(ctx, obj); err != nil {
		// If we've failed to validate the MSR CR is ready then we cannot
		// reliably determine whether it is installed or not.
		return nil, err
	}

	var m msrv1.MSR

	if err := mapstructure.Decode(obj.Object, &m); err != nil {
		return nil, fmt.Errorf("failed to decode msr CR: %w", err)
	}

	filter := constant.MSROperator + "|" + constant.PostgresOperator + "|" + constant.RethinkDBOperator + "|" + constant.CertManager

	releases, err := h.List(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to list helm releases: %w", err)
	}

	if len(releases) == 0 {
		return nil, fmt.Errorf("failed to find any installed helm dependencies")
	}

	installedDeps := make(map[string]helm.ChartDetails)

	for _, rel := range releases {
		installedDeps[rel.Name] = helm.ChartDetails{
			ReleaseName: rel.Name,
			ChartName:   rel.Chart.Name(),
			Version:     rel.Chart.Metadata.Version,
			Installed: func() bool {
				return rel.Info.Status.String() == string(release.StatusDeployed)
			}(),
		}
	}

	version := m.Spec.Tag

	return &api.MSRMetadata{
		Installed:        true,
		InstalledVersion: version,
		MSR3: &api.MSR3Metadata{
			InstalledDependencies: installedDeps,
		},
	}, nil
}

// ApplyCRD applies the MSR CRD to the cluster then waits for it to be ready.
func ApplyCRD(ctx context.Context, msr *msrv1.MSR, kc *kubeclient.KubeClient) error {
	result := make(map[string]interface{})

	d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &result,
		TagName: "json",
	})
	if err != nil {
		return fmt.Errorf("failed to create decoder: %w", err)
	}

	if err := d.Decode(msr); err != nil {
		return fmt.Errorf("failed to decode MSR CR into map: %w", err)
	}

	obj := &unstructured.Unstructured{Object: result}

	// Set specific fields to ensure the object is valid and remove the TypeMeta
	// field, this is to workaround an issue with mapstructure decoding "inline"
	// tagged fields into map, we're effectively rebuilding the inlined TypeMeta
	// fields here.
	obj.SetKind(msr.Kind)
	obj.SetAPIVersion(msr.APIVersion)
	obj.SetCreationTimestamp(metav1.Time{Time: time.Now()})
	delete(obj.Object, "TypeMeta")

	if err := kc.ApplyMSRCR(ctx, obj); err != nil {
		return err
	}

	return kc.WaitForMSRCRReady(ctx, obj)
}
