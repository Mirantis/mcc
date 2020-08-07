package v1beta2

import (
	"fmt"

	"github.com/Mirantis/mcc/pkg/constant"
	"github.com/Mirantis/mcc/pkg/util"
)

// UcpConfig has all the bits needed to configure UCP during installation
type UcpConfig struct {
	Version         string    `yaml:"version"`
	ImageRepo       string    `yaml:"imageRepo,omitempty"`
	InstallFlags    []string  `yaml:"installFlags,omitempty,flow"`
	ConfigFile      string    `yaml:"configFile,omitempty" validate:"omitempty,file"`
	ConfigData      string    `yaml:"configData,omitempty"`
	LicenseFilePath string    `yaml:"licenseFilePath,omitempty" validate:"omitempty,file"`
	Cloud           *UcpCloud `yaml:"cloud,omitempty"`

	Metadata *UcpMetadata `yaml:"-"`
}

// UcpMetadata has the "runtime" discovered metadata of already existing installation.
type UcpMetadata struct {
	Installed        bool
	InstalledVersion string
	ClusterID        string
}

// UcpCloud has the cloud provider configuration
type UcpCloud struct {
	Provider   string `yaml:"provider,omitempty" validate:"required"`
	ConfigFile string `yaml:"configFile,omitempty" validate:"omitempty,file"`
	ConfigData string `yaml:"configData,omitempty"`
}

// UnmarshalYAML sets in some sane defaults when unmarshaling the data from yaml
func (c *UcpConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type rawUcpConfig UcpConfig
	config := NewUcpConfig()
	raw := rawUcpConfig(config)
	if err := unmarshal(&raw); err != nil {
		return err
	}

	if raw.ConfigFile != "" {
		configData, err := util.LoadExternalFile(raw.ConfigFile)
		if err != nil {
			return err
		}
		raw.ConfigData = string(configData)
	}

	if raw.Cloud != nil && raw.Cloud.ConfigFile != "" {
		cloudConfigData, err := util.LoadExternalFile(raw.Cloud.ConfigFile)
		if err != nil {
			return err
		}
		raw.Cloud.ConfigData = string(cloudConfigData)
	}

	*c = UcpConfig(raw)
	return nil
}

// NewUcpConfig creates new config with sane defaults
func NewUcpConfig() UcpConfig {
	return UcpConfig{
		Version:   constant.UCPVersion,
		ImageRepo: constant.ImageRepo,
	}
}

// GetBootstrapperImage combines the bootstrapper image name based on user given config
func (c *UcpConfig) GetBootstrapperImage() string {
	return fmt.Sprintf("%s/ucp:%s", c.ImageRepo, c.Version)
}
