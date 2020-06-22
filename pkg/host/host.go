package host

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"strings"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type RemoteHost interface {
	Connect() error
	Disconnect() error
}

type Hosts []*Host

type Host struct {
	Address          string   `yaml:"address" validate:"required,hostname|ip"`
	User             string   `yaml:"user"`
	SSHPort          int      `yaml:"sshPort" validate:"gt=0,lte=65535"`
	SSHKeyPath       string   `yaml:"sshKeyPath" validate:"file"`
	Role             string   `yaml:"role" validate:"oneof=master worker"`
	ExtraArgs        []string `yaml:"extraArgs"`
	PrivateInterface string   `validate:"omitempty,gt=2"`

	sshClient *ssh.Client
}

// Normalize puts in the defaults
// FIXME Maybe better to handle this during yaml unmarshaling...
func (h *Host) Normalize() {
	if h.SSHKeyPath == "" {
		homeDir, _ := homedir.Dir()
		h.SSHKeyPath = path.Join(homeDir, ".ssh", "id_rsa")
	}

	if h.SSHPort == 0 {
		h.SSHPort = 22
	}
}

// Connect to the host
func (h *Host) Connect() error {
	key, err := ioutil.ReadFile(h.SSHKeyPath)
	if err != nil {
		return err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return err
	}
	config := ssh.ClientConfig{
		User: h.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	address := fmt.Sprintf("%s:%d", h.Address, h.SSHPort)

	client, err := ssh.Dial("tcp", address, &config)
	if err != nil {
		return err
	}
	h.sshClient = client

	return nil
}

// Exec a command on the host and streams the logs
func (h *Host) Exec(cmd string) error {
	session, err := h.sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		return err
	}

	logrus.Debugf("executing command: %s", cmd)
	if err := session.Start(cmd); err != nil {
		return err
	}

	multiReader := io.MultiReader(stdout, stderr)
	outputScanner := bufio.NewScanner(multiReader)

	for outputScanner.Scan() {
		logrus.Debugf("%s:  %s", h.Address, outputScanner.Text())
	}
	if err := outputScanner.Err(); err != nil {
		logrus.Errorf("%s:  %s", h.Address, err.Error())
	}

	return nil
}

// ExecWithOutput execs a command on the host and return output
func (h *Host) ExecWithOutput(cmd string) (string, error) {
	session, err := h.sshClient.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return trimOutput(output), err
	}

	return trimOutput(output), nil
}

func trimOutput(output []byte) string {
	if len(output) > 0 {
		return strings.TrimSpace(string(output))
	}

	return ""
}

func (h *Host) PullImage(name string) error {
	output, err := h.ExecWithOutput(fmt.Sprintf("sudo docker pull %s", name))
	if err != nil {
		log.Warnf("%s: failed to pull image %s: \n%s", h.Address, name, output)
		return err
	}
	return nil
}

func (h *Host) SwarmAddress() string {
	return fmt.Sprintf("%s:%d", h.Address, 2377)
}