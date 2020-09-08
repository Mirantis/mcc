package phase

import (
	api "github.com/Mirantis/mcc/pkg/apis/v1beta3"
)

// CleanUp phase implementation does all the prep work we need for the hosts
type CleanUp struct {
	Analytics
	BasicPhase
}

// Title for the phase
func (p *CleanUp) Title() string {
	return "Clean up"
}

// Run does all the prep work on the hosts in parallel
func (p *CleanUp) Run() error {
	err := runParallelOnHosts(p.config.Spec.Hosts, p.config, p.cleanupEnv)
	if err != nil {
		return err
	}

	return nil
}

func (p *CleanUp) cleanupEnv(host *api.Host, c *api.ClusterConfig) error {
	if len(host.Environment) > 0 {
		return host.Configurer.CleanupEnvironment()
	}
	return nil
}
