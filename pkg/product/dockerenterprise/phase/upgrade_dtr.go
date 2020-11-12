package phase

import (
	"fmt"
	"strings"

	"github.com/Mirantis/mcc/pkg/dtr"
	"github.com/Mirantis/mcc/pkg/exec"
	"github.com/Mirantis/mcc/pkg/phase"

	log "github.com/sirupsen/logrus"
)

// UpgradeDtr is the phase implementation for running the actual DTR upgrade container
type UpgradeDtr struct {
	phase.Analytics
	DtrPhase
}

// Title prints the phase title
func (p *UpgradeDtr) Title() string {
	return "Upgrade DTR components"
}

// Run the upgrade container
func (p *UpgradeDtr) Run() error {
	dtrLeader := p.Config.Spec.DtrLeader()

	err := p.Config.Spec.CheckUCPHealthRemote(dtrLeader)
	if err != nil {
		return fmt.Errorf("%s: failed to health check ucp, try to set `--ucp-url` installFlag and check connectivity", dtrLeader)
	}

	p.EventProperties = map[string]interface{}{
		"dtr_upgraded": false,
	}

	if p.Config.Spec.Dtr.Version == p.Config.Spec.Dtr.Metadata.InstalledVersion {
		log.Infof("%s: DTR cluster already at version %s, not running upgrade", dtrLeader, p.Config.Spec.Dtr.Version)
		return nil
	}

	runFlags := []string{"--rm", "-i"}
	if dtrLeader.Configurer.SELinuxEnabled() {
		runFlags = append(runFlags, "--security-opt label=disable")
	}
	upgradeFlags := []string{
		fmt.Sprintf("--existing-replica-id %s", p.Config.Spec.Dtr.Metadata.DtrLeaderReplicaID),
	}
	ucpFlags := dtr.BuildUCPFlags(p.Config)
	upgradeFlags = append(upgradeFlags, ucpFlags...)
	for _, f := range dtr.PluckSharedInstallFlags(p.Config.Spec.Dtr.InstallFlags, dtr.SharedInstallUpgradeFlags) {
		upgradeFlags = append(upgradeFlags, f)
	}

	upgradeCmd := dtrLeader.Configurer.DockerCommandf("run %s %s upgrade %s", strings.Join(runFlags, " "), p.Config.Spec.Dtr.GetBootstrapperImage(), strings.Join(upgradeFlags, " "))
	log.Debug("Running DTR upgrade via bootstrapper")
	err = dtrLeader.Exec(upgradeCmd, exec.StreamOutput())
	if err != nil {
		return fmt.Errorf("failed to run DTR upgrade: %s", err.Error())
	}

	dtrMeta, err := dtr.CollectFacts(dtrLeader)
	if err != nil {
		return fmt.Errorf("%s: failed to collect existing DTR details: %s", dtrLeader, err.Error())
	}

	// Check to make sure installedversion matches bootstrapperVersion
	if dtrMeta.InstalledVersion != p.Config.Spec.Dtr.Version {
		// If our newly collected facts do not match the version we upgraded to
		// then the upgrade has failed
		return fmt.Errorf("%s: upgraded DTR version: %s does not match intended upgrade version: %s", dtrLeader, dtrMeta.InstalledVersion, p.Config.Spec.Dtr.Version)
	}

	p.EventProperties["dtr_upgraded"] = true
	p.EventProperties["dtr_installed_version"] = p.Config.Spec.Dtr.Metadata.InstalledVersion
	p.EventProperties["dtr_upgraded_version"] = p.Config.Spec.Dtr.Version
	p.Config.Spec.Dtr.Metadata = dtrMeta

	return nil
}