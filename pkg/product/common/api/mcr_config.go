package api

import (
	"github.com/Mirantis/mcc/pkg/constant"
)

// MCRConfig holds the Mirantis Container Runtime installation specific options
type MCRConfig struct {
	Version           string `yaml:"version"`
	RepoURL           string `yaml:"repoURL,omitempty"`
	InstallURLLinux   string `yaml:"installURLLinux,omitempty"`
	InstallURLWindows string `yaml:"installURLWindows,omitempty"`
	Channel           string `yaml:"channel,omitempty"`
}

// UnmarshalYAML puts in sane defaults when unmarshaling from yaml
func (c *MCRConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type mcrConfig MCRConfig
	yc := (*mcrConfig)(c)

	if err := unmarshal(yc); err != nil {
		return err
	}

	c.SetDefaults()

	return nil
}

// SetDefaults sets defaults on the object
func (c *MCRConfig) SetDefaults() {
	// Constants can't be used in tags, so yaml defaults can't be used here.
	if c.Version == "" {
		c.Version = constant.MCRVersion
	}

	if c.Channel == "" {
		c.Channel = constant.MCRChannel
	}

	if c.RepoURL == "" {
		c.RepoURL = constant.MCRRepoURL
	}

	if c.InstallURLLinux == "" {
		c.InstallURLLinux = constant.MCRInstallURLLinux
	}

	if c.InstallURLWindows == "" {
		c.InstallURLWindows = constant.MCRInstallURLWindows
	}
}