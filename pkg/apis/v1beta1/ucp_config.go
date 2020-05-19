package v1beta1

import (
	"fmt"

	"github.com/Mirantis/mcc/pkg/constant"
)

// UcpConfig has all the bits needed to configure UCP during installation
type UcpConfig struct {
	Version      string   `yaml:"version"`
	ImageRepo    string   `yaml:"imageRepo"`
	InstallFlags []string `yaml:"installFlags,flow"`
	ConfigFile   string   `yaml:"configFile" validate:"omitempty,file"`
	ConfigData   string   `yaml:"configData"`

	Metadata *UcpMetadata
}

// UcpMetadata has the "runtime" discovered metadata of already existing installation.
type UcpMetadata struct {
	Installed        bool
	InstalledVersion string
	ClusterID        string
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
		configData, err := loadExternalFile(raw.ConfigFile)
		if err != nil {
			return err
		}
		raw.ConfigData = string(configData)
	}

	*c = UcpConfig(raw)
	return nil
}

// NewUcpConfig creates new config with sane defaults
func NewUcpConfig() UcpConfig {
	return UcpConfig{
		Version:   constant.Version,
		ImageRepo: constant.ImageRepo,
	}
}

// GetBootstrapperImage combines the bootstrapper image name based on user given config
func (c *UcpConfig) GetBootstrapperImage() string {
	return fmt.Sprintf("%s/ucp:%s", c.ImageRepo, c.Version)
}