package phase

import (
	"fmt"
	"math"
	"sync"

	"github.com/Mirantis/mcc/pkg/analytics"
	api "github.com/Mirantis/mcc/pkg/apis/v1beta1"
	retry "github.com/avast/retry-go"
	"github.com/gammazero/workerpool"
	log "github.com/sirupsen/logrus"
)

// InstallEngine phase implementation
type InstallEngine struct {
	Analytics
}

// Title for the phase
func (p *InstallEngine) Title() string {
	return "Install Docker EE Engine on the hosts"
}

// Run installs the engine on each host
func (p *InstallEngine) Run(c *api.ClusterConfig) error {
	props := analytics.NewAnalyticsEventProperties()
	props["engine_version"] = c.Spec.Engine.Version
	p.EventProperties = props
	err := p.upgradeEngines(c)
	if err != nil {
		return err
	}
	newHosts := []*api.Host{}
	for _, h := range c.Spec.Hosts {
		if h.Metadata.EngineVersion == "" {
			newHosts = append(newHosts, h)
		}
	}
	return runParallelOnHosts(newHosts, c, p.installEngine)
}

// Upgrades host docker engines, first managers (one-by-one) and then ~10% rolling update to workers
// TODO: should we drain?
func (p *InstallEngine) upgradeEngines(c *api.ClusterConfig) error {
	for _, h := range c.Spec.Managers() {
		if h.Metadata.EngineVersion != "" && h.Metadata.EngineVersion != c.Spec.Engine.Version {
			err := p.installEngine(h, c)
			if err != nil {
				return err
			}
			if c.Spec.Ucp.Metadata.Installed {
				err = retry.Do( // wait for ucp api to be healthy if UCP is already installed before engine upgrade
					func() error {
						return h.Exec("curl -k -f https://localhost/_ping")
					},
					retry.Attempts(10), // last attempt should wait 1min30s, should be long enough
				)
				if err != nil {
					return err
				}
			}
		}
	}

	workers := []*api.Host{}
	for _, h := range c.Spec.Workers() {
		if h.Metadata.EngineVersion != "" && h.Metadata.EngineVersion != c.Spec.Engine.Version {
			workers = append(workers, h)
		}
	}

	// sacrifice 10% of workers for upgrade gods
	concurrentUpgrades := int(math.Floor(float64(len(workers)) * 0.10))
	if concurrentUpgrades == 0 {
		concurrentUpgrades = 1
	}
	wp := workerpool.New(concurrentUpgrades)
	mu := sync.Mutex{}
	installErrors := &Error{}
	for _, host := range workers {
		if host.Metadata.EngineVersion != "" {
			h := host
			wp.Submit(func() {
				err := p.installEngine(h, c)
				if err != nil {
					mu.Lock()
					installErrors.Errors = append(installErrors.Errors)
					mu.Unlock()
				}
			})
		}
	}
	wp.StopWait()
	if installErrors.Count() > 0 {
		return installErrors
	}
	return nil
}

func (p *InstallEngine) installEngine(host *api.Host, c *api.ClusterConfig) error {
	newInstall := host.Metadata.EngineVersion == ""
	prevVersion := resolveEngineVersion(host)
	err := retry.Do(
		func() error {
			if newInstall {
				log.Infof("%s: installing engine (%s)", host.Address, c.Spec.Engine.Version)
			} else {
				log.Infof("%s: updating engine (%s -> %s)", host.Address, prevVersion, c.Spec.Engine.Version)
			}
			return host.Configurer.InstallEngine(&c.Spec.Engine)
		},
	)
	if err != nil {
		if newInstall {
			log.Errorf("%s: failed to install engine -> %s", host.Address, err.Error())
		} else {
			log.Errorf("%s: failed to update engine -> %s", host.Address, err.Error())
		}

		return err
	}

	currentVersion := resolveEngineVersion(host)
	if !newInstall && currentVersion == prevVersion {
		err = host.Configurer.RestartEngine()
		if err != nil {
			return NewError(fmt.Sprintf("%s: failed to restart engine", host.Address))
		}
	}

	log.Printf("%s: engine version %s installed", host.Address, c.Spec.Engine.Version)
	return nil
}
