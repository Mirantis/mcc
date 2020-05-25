package v1beta1

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Mirantis/mcc/pkg/exec"
	"github.com/Mirantis/mcc/pkg/util"
	"github.com/creasty/defaults"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"

	log "github.com/sirupsen/logrus"
)

// RemoteHost interface defines the connection (ssh) related interface each remote host should implement
type RemoteHost interface {
	Connect() error
	Disconnect() error
}

// OsRelease host operating system info
type OsRelease struct {
	ID      string
	IDLike  string
	Name    string
	Version string
}

// HostMetadata resolved metadata for host
type HostMetadata struct {
	Hostname        string
	InternalAddress string
	EngineVersion   string
	Os              *OsRelease
}

// Hosts is the type alias for slice of Hosts
type Hosts []*Host

// Host contains all the needed details to work with hosts
type Host struct {
	Address          string   `yaml:"address" validate:"required,hostname|ip"`
	User             string   `yaml:"user" validate:"omitempty,gt=2" default:"root"`
	SSHPort          int      `yaml:"sshPort" default:"22" validate:"gt=0,lte=65535"`
	SSHKeyPath       string   `yaml:"sshKeyPath" validate:"file" default:"~/.ssh/id_rsa"`
	Role             string   `yaml:"role" validate:"oneof=manager worker"`
	ExtraArgs        []string `yaml:"extraArgs"`
	PrivateInterface string   `yaml:"privateInterface" default:"eth0" validate:"gt=2"`
	Metadata         *HostMetadata
	Configurer       HostConfigurer

	sshClient *ssh.Client
}

func (h *Host) SshSession() (*ssh.Session, error) {
	session, err := h.sshClient.NewSession()
	return session, err
}

func (h *Host) Name() string {
	return h.Address
}

// UnmarshalYAML sets in some sane defaults when unmarshaling the data from yaml
func (h *Host) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(h)
	// Need to expand possible ~... paths so validation will pass
	h.SSHKeyPath, _ = homedir.Expand(h.SSHKeyPath)
	type plain Host
	if err := unmarshal((*plain)(h)); err != nil {
		return err
	}

	return nil
}

// Connect to the host
func (h *Host) Connect() error {
	key, err := util.LoadExternalFile(h.SSHKeyPath)
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User:            h.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	address := fmt.Sprintf("%s:%d", h.Address, h.SSHPort)

	sshAgentSock := os.Getenv("SSH_AUTH_SOCK")
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil && sshAgentSock == "" {
		return err
	}
	if err == nil {
		config.Auth = append(config.Auth, ssh.PublicKeys(signer))
	}

	if sshAgentSock != "" {
		sshAgent, err := net.Dial("unix", sshAgentSock)
		if err != nil {
			return fmt.Errorf("cannot connect to SSH agent auth socket %s: %s", sshAgentSock, err)
		}
		log.Debugf("using SSH auth sock %s", sshAgentSock)
		config.Auth = append(config.Auth, ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers))
	}

	client, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return err
	}
	h.sshClient = client

	return nil
}

// Exec runs a command on the host
func (h *Host) Exec(cmd string, opts ...exec.Option) error {
	return exec.Cmd(h, cmd, opts...)
}

// ExecWithOutput runs a command on the host and returns the output as a string
func (h *Host) ExecWithOutput(cmd string, opts ...exec.Option) (string, error) {
	return exec.CmdWithOutput(h, cmd, opts...)
}

// WriteFile writes a file to the host with given contents
func (h *Host) WriteFile(path string, data string, permissions string) error {
	tempFile, _ := h.ExecWithOutput("mktemp")
	err := h.Exec(fmt.Sprintf("cat > %s && (sudo install -m %s %s %s || (rm %s; exit 1))", tempFile, permissions, tempFile, path, tempFile), exec.Stdin(data))
	if err != nil {
		return err
	}
	return nil
}

func trimOutput(output []byte) string {
	if len(output) > 0 {
		return strings.TrimSpace(string(output))
	}

	return ""
}

// AuthenticateDocker performs a docker login on the host using local REGISTRY_USERNAME
// and REGISTRY_PASSWORD when set
func (h *Host) AuthenticateDocker(server string) error {
	if user := os.Getenv("REGISTRY_USERNAME"); user != "" {
		pass := os.Getenv("REGISTRY_PASSWORD")
		if pass == "" {
			return fmt.Errorf("%s: REGISTRY_PASSWORD not set", h.Address)
		}
		log.Infof("%s: authenticating docker", h.Address)
		err := h.Exec(h.Configurer.DockerCommandf("login -u %s --password-stdin %s", user, server), exec.Stdin(pass), exec.Redact(pass))

		if err != nil {
			return fmt.Errorf("%s: failed to authenticate docker: %s", h.Address, err)
		}
	} else {
		log.Debugf("%s: REGISTRY_USERNAME not set, not authenticating", h.Address)
	}
	return nil
}

// PullImage pulls the named docker image on the host
func (h *Host) PullImage(name string) error {
	output, err := h.ExecWithOutput(h.Configurer.DockerCommandf("pull %s", name))
	if err != nil {
		log.Warnf("%s: failed to pull image %s: \n%s", h.Address, name, output)
		return err
	}
	return nil
}

// SwarmAddress determines the swarm address for the host
func (h *Host) SwarmAddress() string {
	return fmt.Sprintf("%s:%d", h.Metadata.InternalAddress, 2377)
}

// IsWindows returns true if host has been detected running windows
func (h *Host) IsWindows() bool {
	if h.Metadata == nil || h.Metadata.Os == nil {
		return false
	}
	return strings.HasPrefix(h.Metadata.Os.ID, "windows-")
}
