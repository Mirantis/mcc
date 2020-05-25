package exec

import (
	"bufio"
	"io"
	"regexp"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

type host interface {
	IsWindows() bool
	Name() string
	SshSession() (*ssh.Session, error)
}

type Option func(*Options)

type Options struct {
	Stdin      string
	LogInfo    bool
	LogDebug   bool
	LogError   bool
	LogCommand bool
	Redact     string
	Output     *string
}

var defaultOptions = Options{
	Stdin:      "",
	LogInfo:    false,
	LogCommand: true,
	LogDebug:   true,
	LogError:   true,
	Redact:     "",
	Output:     nil,
}

// Stdin exec option for sending data to the command through stdin
func Stdin(t string) Option {
	return func(o *Options) {
		o.Stdin = t
	}
}

// StreamOutput exec option for sending the command output to info log
func StreamOutput() Option {
	return func(o *Options) {
		o.LogInfo = true
	}
}

// HideCommand exec option for hiding the command-string and stdin contents from the logs
func HideCommand() Option {
	return func(o *Options) {
		o.LogCommand = true
	}
}

// HideOutput exec option for hiding the command output from logs
func HideOutput() Option {
	return func(o *Options) {
		o.LogDebug = false
	}
}

// Sensitive exec option for disabling all logging of the command
func Sensitive() Option {
	return func(o *Options) {
		o.LogDebug = false
		o.LogInfo = false
		o.LogError = false
		o.LogCommand = false
	}
}

// Redact exec option for defining a redact regexp pattern that will be replaced with [REDACTED] in the logs
func Redact(s string) Option {
	return func(o *Options) {
		o.Redact = s
	}
}

func redactFunc(rs string) func(s string) string {
	if rs == "" {
		return func(s string) string {
			return s
		}
	} else {
		re := *regexp.MustCompile(rs)
		return func(s string) string {
			return re.ReplaceAllString(s, "[REDACTED]")
		}
	}
}

// Cmd runs a command on the host
func Cmd(h host, cmd string, opts ...Option) error {
	options := defaultOptions
	for _, o := range opts {
		o(&options)
	}

	redact := redactFunc(options.Redact)

	session, err := h.SshSession()
	if err != nil {
		return err
	}
	defer session.Close()

	if options.Stdin == "" && !h.IsWindows() {
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

	if options.LogCommand {
		log.Debugf("executing command: %s", redact(cmd))
	}

	if err := session.Start(cmd); err != nil {
		return err
	}

	if options.Stdin != "" {
		if options.LogCommand {
			log.Debugf("writing data to command stdin: %s", redact(options.Stdin))
		}

		go func() {
			defer stdinPipe.Close()
			io.WriteString(stdinPipe, options.Stdin)
		}()
	}

	multiReader := io.MultiReader(stdout, stderr)
	outputScanner := bufio.NewScanner(multiReader)

	for outputScanner.Scan() {
		text := outputScanner.Text()
		if options.Output != nil {
			*options.Output += text + "\n"
		}

		if options.LogInfo {
			log.Infof("%s:  %s", h.Name(), redact(text))
		} else if options.LogDebug {
			log.Debugf("%s:  %s", h.Name(), redact(text))
		}
	}

	if options.LogError {
		if err := outputScanner.Err(); err != nil {
			log.Errorf("%s:  %s", h.Name(), err.Error())
		}
	}

	return session.Wait()
}

// withOutput a helper to build a Stdin option
func withOutput(t *string) Option {
	return func(o *Options) {
		o.Output = t
	}
}

// CmdWithOutput runs a command on the host and returns the output as a string
func CmdWithOutput(h host, cmd string, opts ...Option) (string, error) {
	output := ""
	opts = append(opts, withOutput(&output))
	err := Cmd(h, cmd, opts...)
	return output, err
}