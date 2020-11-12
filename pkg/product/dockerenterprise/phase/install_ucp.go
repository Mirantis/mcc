package phase

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Mirantis/mcc/pkg/api"
	"github.com/Mirantis/mcc/pkg/exec"
	"github.com/Mirantis/mcc/pkg/phase"
	"github.com/Mirantis/mcc/pkg/ucp"
	"github.com/Mirantis/mcc/pkg/util"
	log "github.com/sirupsen/logrus"
)

const configName string = "com.docker.ucp.Config"

// InstallUCP is the phase implementation for running the actual UCP installer container
type InstallUCP struct {
	phase.Analytics
	phase.BasicPhase
	SkipCleanup bool
}

// Title prints the phase title
func (p *InstallUCP) Title() string {
	return "Install UCP components"
}

// Run the installer container
func (p *InstallUCP) Run() (err error) {
	swarmLeader := p.Config.Spec.SwarmLeader()

	if !p.SkipCleanup {
		defer func() {
			if err != nil {
				log.Println("Cleaning-up")
				if cleanupErr := cleanupUcp(swarmLeader); cleanupErr != nil {
					log.Warnln("Error while cleaning-up resources")
				}
			}
		}()
	}

	p.EventProperties = map[string]interface{}{
		"ucp_version": p.Config.Spec.Ucp.Version,
	}

	if p.Config.Spec.Ucp.Metadata.Installed {
		log.Infof("%s: UCP already installed at version %s, not running installer", swarmLeader, p.Config.Spec.Ucp.Metadata.InstalledVersion)
		return nil
	}

	image := fmt.Sprintf("%s/ucp:%s", p.Config.Spec.Ucp.ImageRepo, p.Config.Spec.Ucp.Version)
	installFlags := p.Config.Spec.Ucp.InstallFlags

	if p.Config.Spec.Ucp.CACertData != "" && p.Config.Spec.Ucp.CertData != "" && p.Config.Spec.Ucp.KeyData != "" {
		err := p.installCertificates(p.Config)
		if err != nil {
			return err
		}
		installFlags.AddUnlessExist("--external-server-cert")
	}

	if p.Config.Spec.Ucp.ConfigData != "" {
		defer func() {
			err := swarmLeader.Exec(swarmLeader.Configurer.DockerCommandf("config rm %s", configName))
			if err != nil {
				log.Warnf("Failed to remove the temporary UCP installer configuration %s : %s", configName, err)
			}
		}()

		installFlags.AddUnlessExist("--existing-config")
		log.Info("Creating UCP configuration")
		configCmd := swarmLeader.Configurer.DockerCommandf("config create %s -", configName)
		err := swarmLeader.Exec(configCmd, exec.Stdin(p.Config.Spec.Ucp.ConfigData))
		if err != nil {
			return err
		}
	}

	if licenseFilePath := p.Config.Spec.Ucp.LicenseFilePath; licenseFilePath != "" {
		log.Debugf("Installing UCP with LicenseFilePath: %s", licenseFilePath)
		licenseFlag, err := util.SetupLicenseFile(p.Config.Spec.Ucp.LicenseFilePath)
		if err != nil {
			return fmt.Errorf("error while reading license file %s: %v", licenseFilePath, err)
		}
		installFlags.AddUnlessExist(licenseFlag)
	}

	if p.Config.Spec.Ucp.Cloud != nil {
		if p.Config.Spec.Ucp.Cloud.Provider != "" {
			installFlags.AddUnlessExist("--cloud-provider " + p.Config.Spec.Ucp.Cloud.Provider)
		}
		if p.Config.Spec.Ucp.Cloud.ConfigData != "" {
			applyCloudConfig(p.Config)
		}
	}

	if api.IsCustomImageRepo(p.Config.Spec.Ucp.ImageRepo) {
		// In case of custom repo, don't let UCP check the images
		installFlags.AddUnlessExist("--pull never")
	}
	runFlags := []string{"--rm", "-i", "-v /var/run/docker.sock:/var/run/docker.sock"}
	if swarmLeader.Configurer.SELinuxEnabled() {
		runFlags = append(runFlags, "--security-opt label=disable")
	}

	if p.Config.Spec.Ucp.AdminUsername != "" {
		installFlags.AddUnlessExist("--admin-username " + p.Config.Spec.Ucp.AdminUsername)
	}

	if p.Config.Spec.Ucp.AdminPassword != "" {
		installFlags.AddUnlessExist("--admin-password " + p.Config.Spec.Ucp.AdminPassword)
	}

	installCmd := swarmLeader.Configurer.DockerCommandf("run %s %s install %s", strings.Join(runFlags, " "), image, strings.Join(installFlags, " "))
	output, err := swarmLeader.ExecWithOutput(installCmd, exec.StreamOutput(), exec.RedactString(p.Config.Spec.Ucp.AdminUsername, p.Config.Spec.Ucp.AdminPassword))
	if err != nil {
		return fmt.Errorf("%s: failed to run UCP installer: %s", swarmLeader, err.Error())
	}

	if installFlags.GetValue("--admin-password") == "" {
		re := regexp.MustCompile(`msg="Generated random admin password: (.+?)"`)
		md := re.FindStringSubmatch(output)
		if len(md) > 0 && md[1] != "" {
			log.Warnf("Using an automatically generated password for UCP admin user: %s -- you will have to set it to Spec.Ucp.AdminPassword for any subsequent launchpad runs.", md[1])
			p.Config.Spec.Ucp.AdminPassword = md[1]
			if p.Config.Spec.Ucp.AdminUsername == "" {
				log.Debugf("defaulting to ucp admin username 'admin'")
				p.Config.Spec.Ucp.AdminUsername = "admin"
			}
		}
	}

	err = ucp.CollectFacts(swarmLeader, p.Config.Spec.Ucp.Metadata)
	if err != nil {
		return fmt.Errorf("%s: failed to collect existing UCP details: %s", swarmLeader, err.Error())
	}

	return nil
}

// installCertificates installs user supplied UCP certificates
func (p *InstallUCP) installCertificates(config *api.ClusterConfig) error {
	log.Infof("Installing UCP certificates")
	managers := config.Spec.Managers()
	err := managers.ParallelEach(func(h *api.Host) error {
		err := h.Exec(h.Configurer.DockerCommandf("volume inspect ucp-controller-server-certs"))
		if err != nil {
			log.Infof("%s: creating ucp-controller-server-certs volume", h)
			err := h.Exec(h.Configurer.DockerCommandf("volume create ucp-controller-server-certs"))
			if err != nil {
				return err
			}
		}

		dir, err := h.ExecWithOutput(h.Configurer.DockerCommandf(`volume inspect ucp-controller-server-certs --format "{{ .Mountpoint }}"`))
		if err != nil {
			return err
		}

		log.Infof("%s: installing certificate files to %s", h, dir)
		err = h.Configurer.WriteFile(fmt.Sprintf("%s/ca.pem", dir), config.Spec.Ucp.CACertData, "0600")
		if err != nil {
			return err
		}
		err = h.Configurer.WriteFile(fmt.Sprintf("%s/cert.pem", dir), config.Spec.Ucp.CertData, "0600")
		if err != nil {
			return err
		}
		err = h.Configurer.WriteFile(fmt.Sprintf("%s/key.pem", dir), config.Spec.Ucp.KeyData, "0600")
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func applyCloudConfig(config *api.ClusterConfig) error {
	configData := config.Spec.Ucp.Cloud.ConfigData
	provider := config.Spec.Ucp.Cloud.Provider

	var destFile string
	if provider == "azure" {
		destFile = "/etc/kubernetes/azure.json"
	} else if provider == "openstack" {
		destFile = "/etc/kubernetes/openstack.conf"
	} else {
		return fmt.Errorf("Spec.Cloud.configData is only supported with Azure and OpenStack cloud providers")
	}

	err := phase.RunParallelOnHosts(config.Spec.Hosts, config, func(h *api.Host, c *api.ClusterConfig) error {
		log.Infof("%s: copying cloud provider (%s) config to %s", h, provider, destFile)
		return h.Configurer.WriteFile(destFile, configData, "0700")
	})

	return err
}

func cleanupUcp(h *api.Host) error {
	containersToRemove, err := h.ExecWithOutput(h.Configurer.DockerCommandf("ps -aq --filter name=ucp-"))
	if err != nil {
		return err
	}
	if strings.Trim(containersToRemove, " ") == "" {
		log.Debugf("No containers to remove")
		return nil
	}
	containersToRemove = strings.ReplaceAll(containersToRemove, "\n", " ")
	if err := h.Exec(h.Configurer.DockerCommandf("rm -f %s", containersToRemove)); err != nil {
		return err
	}

	return nil
}