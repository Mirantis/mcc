package phase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mirantis/mcc/pkg/helm"
	"github.com/Mirantis/mcc/pkg/mke"
	"github.com/Mirantis/mcc/pkg/phase"
	"github.com/hashicorp/go-version"
	log "github.com/sirupsen/logrus"
	"k8s.io/utils/ptr"
)

// ConfigureDepsMSR3 phase implementation configures any Helm dependencies for
// msr-operator to be able to deploy and run an MSR CR.
type ConfigureDepsMSR3 struct {
	phase.Analytics
	phase.KubernetesPhase

	dependencyUpgrades []helm.ReleaseDetails
}

// Title for the phase.
func (p *ConfigureDepsMSR3) Title() string {
	return "Configuring MSR dependencies"
}

func (p *ConfigureDepsMSR3) Prepare(config interface{}) error {
	var err error

	p.Config, err = convertConfigToClusterConfig(config)
	if err != nil {
		return err
	}

	p.Kube, p.Helm, err = mke.KubeAndHelmFromConfig(p.Config)
	if err != nil {
		return fmt.Errorf("failed to get kube and helm clients: %w", err)
	}

	for _, releaseDetails := range p.Config.Spec.MSR.V3.Dependencies.List() {
		vers, err := version.NewSemver(releaseDetails.Version)
		if err != nil {
			// We should never get here, we should be parsing the version prior
			// to this phase during config validation.
			return fmt.Errorf("failed to parse version %q for dependency %q: %w", releaseDetails.Version, releaseDetails.ReleaseName, err)
		}

		needsUpgrade, err := p.Helm.ChartNeedsUpgrade(releaseDetails.ReleaseName, vers)
		if err != nil {
			// Log any errors that are different than NotFound, but try to
			// upgrade anyway.
			var notFoundErr helm.ReleaseNotFoundError

			if !errors.As(err, &notFoundErr) {
				log.Warnf("failed to check if dependency %q needs upgrade, will try to upgrade anyway: %s", releaseDetails.ReleaseName, err)
			}

			needsUpgrade = true
		}

		if needsUpgrade {
			// If the dependency needs upgrade, add it to the list of
			// dependencies to upgrade.
			p.dependencyUpgrades = append(p.dependencyUpgrades, *releaseDetails)
		}
	}

	return nil
}

func (p *ConfigureDepsMSR3) ShouldRun() bool {
	return p.Config.Spec.ContainsMSR3() && len(p.dependencyUpgrades) > 0
}

// Run configures the dependencies for an MSR CR to be able to deploy by
// installing cert-manager, postgres-operator, rethinkdb-operator and
// msr-operator.  If these are already installed, the phase is a no-op.
func (p *ConfigureDepsMSR3) Run() error {
	for _, releaseDetails := range p.dependencyUpgrades {
		_, err := p.Helm.Upgrade(context.Background(), &helm.Options{
			ReleaseDetails: releaseDetails,
			ReuseValues:    true,
			Wait:           true,
			Atomic:         true,
			Timeout:        ptr.To(helm.DefaultTimeout),
		})
		if err != nil {
			return fmt.Errorf("failed to install/upgrade Helm release %q: %w", releaseDetails.ReleaseName, err)
		}

		log.Infof("dependency %q installed/upgraded", releaseDetails.ReleaseName)
	}

	return nil
}