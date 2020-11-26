package phase

import (
	"fmt"
	"strings"

	"github.com/Mirantis/mcc/pkg/api"
	"github.com/Mirantis/mcc/pkg/exec"
	"github.com/Mirantis/mcc/pkg/phase"
	"github.com/Mirantis/mcc/pkg/swarm"

	log "github.com/sirupsen/logrus"
)

// UninstallMKE is the phase implementation for running MKE uninstall
type UninstallMKE struct {
	phase.Analytics
	phase.BasicPhase
}

// Title prints the phase title
func (p *UninstallMKE) Title() string {
	return "Uninstall MKE components"
}

// Run the installer container
func (p *UninstallMKE) Run() error {
	swarmLeader := p.Config.Spec.SwarmLeader()
	if !p.Config.Spec.MKE.Metadata.Installed {
		log.Infof("%s: MKE is not installed, skipping", swarmLeader)
		return nil
	}

	image := fmt.Sprintf("%s/ucp:%s", p.Config.Spec.MKE.ImageRepo, p.Config.Spec.MKE.Version)
	args := fmt.Sprintf("--id %s", swarm.ClusterID(swarmLeader))
	runFlags := []string{"--rm", "-i", "-v /var/run/docker.sock:/var/run/docker.sock"}
	if swarmLeader.Configurer.SELinuxEnabled() {
		runFlags = append(runFlags, "--security-opt label=disable")
	}
	uninstallCmd := swarmLeader.Configurer.DockerCommandf("run %s %s uninstall-ucp %s", strings.Join(runFlags, " "), image, args)
	err := swarmLeader.Exec(uninstallCmd, exec.StreamOutput(), exec.RedactString(p.Config.Spec.MKE.InstallFlags.GetValue("--admin-username"), p.Config.Spec.MKE.InstallFlags.GetValue("--admin-password")))
	if err != nil {
		return fmt.Errorf("%s: failed to run MKE uninstaller: %s", swarmLeader, err.Error())
	}

	if p.Config.Spec.MKE.CertData != "" {
		managers := p.Config.Spec.Managers()
		managers.ParallelEach(func(h *api.Host) error {
			log.Infof("%s: removing ucp-controller-server-certs volume", h)
			err := h.Exec(h.Configurer.DockerCommandf("volume rm --force ucp-controller-server-certs"))
			if err != nil {
				log.Errorf("%s: failed to remove the volume", h)
			}
			return nil
		})
	}
	return nil
}