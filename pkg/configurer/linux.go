package configurer

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Mirantis/mcc/pkg/util"

	api "github.com/Mirantis/mcc/pkg/apis/v1beta2"
	log "github.com/sirupsen/logrus"
)

// LinuxConfigurer is a generic linux host configurer
type LinuxConfigurer struct {
	Host *api.Host
}

// InstallEngine install Docker EE engine on Linux
func (c *LinuxConfigurer) InstallEngine(engineConfig *api.EngineConfig) error {
	log.Debugf("engine config: %+v", c.Host.DaemonConfig)
	if len(c.Host.DaemonConfig) > 0 {
		daemonJSONData, err := json.Marshal(c.Host.DaemonConfig)
		if err != nil {
			return fmt.Errorf("failed to marshal daemon json config: %w", err)
		}

		cfg := "/etc/docker/daemon.json"
		if c.FileExist(cfg) {
			log.Debugf("deleting %s", cfg)
			if err := c.DeleteFile(cfg); err != nil {
				return err
			}
		}

		log.Debugf("writing %s", cfg)
		if err := c.WriteFile(cfg, string(daemonJSONData), "0700"); err != nil {
			return err
		}
	}

	if c.Host.Metadata.EngineVersion == engineConfig.Version {
		return nil
	}
	cmd := fmt.Sprintf("curl %s | DOCKER_URL=%s CHANNEL=%s VERSION=%s bash", engineConfig.InstallURL, engineConfig.RepoURL, engineConfig.Channel, engineConfig.Version)
	err := c.Host.Exec(cmd)
	if err != nil {
		return err
	}

	err = c.Host.Exec("sudo systemctl enable docker")
	if err != nil {
		return err
	}

	err = c.Host.Exec("sudo systemctl start docker")
	if err != nil {
		return err
	}
	return nil
}

// RestartEngine restarts Docker EE engine
func (c *LinuxConfigurer) RestartEngine() error {
	return c.Host.Exec("sudo systemctl restart docker")
}

// ResolveHostname resolves hostname
func (c *LinuxConfigurer) ResolveHostname() string {
	hostname, _ := c.Host.ExecWithOutput("hostname -s")

	return hostname
}

// ResolveInternalIP resolves internal ip from private interface
func (c *LinuxConfigurer) ResolveInternalIP() (string, error) {
	output, err := c.Host.ExecWithOutput(fmt.Sprintf("ip -o addr show dev %s scope global", c.Host.PrivateInterface))
	if err != nil {
		return "", fmt.Errorf("failed to find private interface with name %s: %s. Make sure you've set correct 'privateInterface' for the host in config", c.Host.PrivateInterface, output)
	}
	return c.ParseInternalIPFromIPOutput(output)
}

// ParseInternalIPFromIPOutput parses internal ip from ip command output
func (c *LinuxConfigurer) ParseInternalIPFromIPOutput(output string) (string, error) {
	lines := strings.Split(output, "\r\n")
	for _, line := range lines {
		items := strings.Fields(line)
		if len(items) < 4 {
			log.Debugf("not enough items in ip address line (%s), skipping...", items)
			continue
		}
		addrItems := strings.Split(items[3], "/")
		if addrItems[0] != c.Host.Address {
			if util.IsValidAddress(addrItems[0]) {
				return addrItems[0], nil
			}

			return "", fmt.Errorf("found address %s for interface %s but it does not seem to be valid address", addrItems[0], c.Host.PrivateInterface)
		}
	}
	// FIXME If we get this far should we just bail out with error!?!?
	return c.Host.Address, nil
}

// IsContainerized checks if host is actually a container
func (c *LinuxConfigurer) IsContainerized() bool {
	err := c.Host.Exec("grep 'container=docker' /proc/1/environ")
	if err != nil {
		return false
	}
	return true
}

// FixContainerizedHost configures host if host is containerized environment
func (c *LinuxConfigurer) FixContainerizedHost() error {
	if c.IsContainerized() {
		return c.Host.Exec("sudo mount --make-rshared /")
	}
	return nil
}

// DockerCommandf accepts a printf-like template string and arguments
// and builds a command string for running the docker cli on the host
func (c *LinuxConfigurer) DockerCommandf(template string, args ...interface{}) string {
	return fmt.Sprintf("sudo docker %s", fmt.Sprintf(template, args...))
}

// ValidateFacts validates all the collected facts so we're sure we can proceed with the installation
func (c *LinuxConfigurer) ValidateFacts() error {
	localAddresses, err := c.getHostLocalAddresses()
	if err != nil {
		return fmt.Errorf("failed to find host local addresses: %w", err)
	}

	if !util.StringSliceContains(localAddresses, c.Host.Metadata.InternalAddress) {
		return fmt.Errorf("discovered private address %s does not seem to be a node local address (%s). Make sure you've set correct 'privateInterface' for the host in config", c.Host.Metadata.InternalAddress, strings.Join(localAddresses, ","))
	}

	return nil
}

// SELinuxEnabled is SELinux enabled
func (c *LinuxConfigurer) SELinuxEnabled() bool {
	output, err := c.Host.ExecWithOutput("sudo getenforce")
	if err != nil {
		return false
	}
	return strings.ToLower(strings.TrimSpace(output)) == "enforcing"
}

func (c *LinuxConfigurer) getHostLocalAddresses() ([]string, error) {
	output, err := c.Host.ExecWithOutput("sudo hostname --all-ip-addresses")
	if err != nil {
		return nil, err
	}

	return strings.Split(output, " "), nil
}

// AuthenticateDocker performs a docker login on the host
func (c *LinuxConfigurer) AuthenticateDocker(user, pass, imageRepo string) error {
	return c.Host.ExecCmd(c.DockerCommandf("login -u %s --password-stdin %s", user, imageRepo), pass, false, true)
}

// WriteFile writes file to host with given contents. Do not use for large files.
func (c *LinuxConfigurer) WriteFile(path string, data string, permissions string) error {
	tempFile, err := c.Host.ExecWithOutput("mktemp")
	if err != nil {
		return err
	}

	err = c.Host.ExecCmd(fmt.Sprintf("cat > %s && (sudo install -D -m %s %s %s || (rm %s; exit 1))", tempFile, permissions, tempFile, path, tempFile), data, false, true)
	if err != nil {
		return err
	}
	return nil
}

// ReadFile reads a files contents from the host.
func (c *LinuxConfigurer) ReadFile(path string) (string, error) {
	return c.Host.ExecWithOutput(fmt.Sprintf("sudo cat \"%s\"", path))
}

// DeleteFile deletes a file from the host.
func (c *LinuxConfigurer) DeleteFile(path string) error {
	return c.Host.ExecCmd(fmt.Sprintf(`sudo rm -f "%s"`, path), "", false, false)
}

// FileExist checks if a file exists on the host
func (c *LinuxConfigurer) FileExist(path string) bool {
	return c.Host.ExecCmd(fmt.Sprintf(`sudo test -e "%s"`, path), "", false, false) == nil
}
