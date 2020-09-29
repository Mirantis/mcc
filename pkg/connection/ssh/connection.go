package ssh

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"sync"

	ssh "golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"

	"github.com/Mirantis/mcc/pkg/exec"
	util "github.com/Mirantis/mcc/pkg/util"
	log "github.com/sirupsen/logrus"
)

// Connection describes an SSH connection
type Connection struct {
	Address string
	User    string
	Port    int
	KeyPath string

	isWindows bool
	client    *ssh.Client
}

// Disconnect closes the SSH connection
func (c *Connection) Disconnect() {
	c.client.Close()
}

// SetWindows can be used to tell the SSH connection to consider the host to be running Windows
func (c *Connection) SetWindows(v bool) {
	c.isWindows = v
}

// IsWindows is true when SetWindows(true) has been used
func (c *Connection) IsWindows() bool {
	return c.isWindows
}

// Connect opens the SSH connection
func (c *Connection) Connect() error {
	key, err := util.LoadExternalFile(c.KeyPath)
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User:            c.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	address := fmt.Sprintf("%s:%d", c.Address, c.Port)

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
	c.client = client

	return nil
}

// Exec executes a command on the host
func (c *Connection) Exec(cmd string, opts ...exec.Option) error {
	o := exec.Build(opts...)
	session, err := c.client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	if o.Stdin == "" && !c.isWindows {
		// FIXME not requesting a pty for commands with stdin input for now,
		// as it appears the pipe doesn't get closed with stdinpipe.Close()
		modes := ssh.TerminalModes{}
		err = session.RequestPty("xterm", 80, 40, modes)
		if err != nil {
			return err
		}
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		return err
	}

	stdinPipe, err := session.StdinPipe()
	if err != nil {
		return err
	}

	o.LogCmd(c.Address, cmd)

	if err := session.Start(cmd); err != nil {
		return err
	}

	if o.Stdin != "" {
		o.LogStdin(c.Address)

		go func() {
			defer stdinPipe.Close()
			io.WriteString(stdinPipe, o.Stdin)
		}()
	}

	multiReader := io.MultiReader(stdout, stderr)
	outputScanner := bufio.NewScanner(multiReader)

	for outputScanner.Scan() {
		o.AddOutput(c.Address, outputScanner.Text()+"\n")
	}

	if err := outputScanner.Err(); err != nil {
		o.LogErrorf("%s:  %s", c.Address, err.Error())
	}

	return session.Wait()
}

// WriteFileLarge copies a larger file to the host.
// Use instead of configurer.WriteFile when it seems appropriate
func (c *Connection) WriteFileLarge(src, dstdir string) error {
	stat, err := os.Stat(src)
	if err != nil {
		return err
	}
	base := path.Base(src)

	log.Infof("%s: copying %d bytes to %s/%s", c.Address, stat.Size(), dstdir, base)

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	session, err := c.client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		hostIn, _ := session.StdinPipe()
		defer hostIn.Close()
		fmt.Fprintf(hostIn, "C0664 %d %s\n", stat.Size(), base)
		io.Copy(hostIn, in)
		fmt.Fprint(hostIn, "\x00")
		wg.Done()
	}()

	err = session.Run(fmt.Sprintf("/usr/bin/scp -t %s/", dstdir))
	wg.Wait()
	log.Debugf("%s: completed file copy", c.Address)
	return err
}
