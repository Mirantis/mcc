package phase

import (
	"fmt"

	"github.com/Mirantis/mcc/pkg/api"
	"github.com/Mirantis/mcc/pkg/docker"
	log "github.com/sirupsen/logrus"
)

// PullUCPImages phase implementation
type PullUCPImages struct {
	Analytics
	BasicPhase
}

// Title for the phase
func (p *PullUCPImages) Title() string {
	return "Pull UCP images"
}

// Run pulls images in parallel across nodes via a workerpool of 5
func (p *PullUCPImages) Run() error {
	images, err := p.ListImages()
	if err != nil {
		return err
	}
	log.Debugf("loaded images list: %v", images)

	imageRepo := p.config.Spec.Ucp.ImageRepo
	if api.IsCustomImageRepo(imageRepo) {
		pullList := docker.AllToRepository(images, imageRepo)
		// In case of custom image repo, we need to pull and retag all the images on all hosts
		return runParallelOnHosts(p.config.Spec.Hosts, p.config, func(h *api.Host, c *api.ClusterConfig) error {
			if err := docker.PullImages(h, pullList); err != nil {
				return err
			}
			return docker.RetagAllToRepository(h, pullList, images[0].Repository)
		})
	}

	// Normally we pull only on managers, let workers pull needed stuff on-demand
	return runParallelOnHosts(p.config.Spec.Managers(), p.config, func(h *api.Host, c *api.ClusterConfig) error {
		return docker.PullImages(h, images)
	})
}

// ListImages obtains a list of images from UCP
func (p *PullUCPImages) ListImages() ([]*docker.Image, error) {
	manager := p.config.Spec.SwarmLeader()
	bootstrap := docker.NewImage(p.config.Spec.Ucp.GetBootstrapperImage())

	if !bootstrap.Exist(manager) {
		if err := bootstrap.Pull(manager); err != nil {
			return []*docker.Image{}, err
		}
	}
	output, err := manager.ExecWithOutput(manager.Configurer.DockerCommandf("run --rm %s images %s --list", bootstrap))
	if err != nil {
		return []*docker.Image{}, fmt.Errorf("%s: failed to get UCP image list", manager)
	}

	return docker.AllFromString(output), nil
}
