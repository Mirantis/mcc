package api

import (
	"github.com/Mirantis/mcc/pkg/constant"
	validator "github.com/go-playground/validator/v10"
)

// ClusterMeta defines cluster metadata
type ClusterMeta struct {
	Name string `yaml:"name" validate:"required"`
}

// ClusterConfig describes launchpad.yaml configuration
type ClusterConfig struct {
	APIVersion string       `yaml:"apiVersion" validate:"eq=launchpad.mirantis.com/v1.1"`
	Kind       string       `yaml:"kind" validate:"eq=DockerEnterprise"`
	Metadata   *ClusterMeta `yaml:"metadata"`
	Spec       *ClusterSpec `yaml:"spec"`
}

// UnmarshalYAML sets in some sane defaults when unmarshaling the data from yaml
func (c *ClusterConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Metadata = &ClusterMeta{
		Name: "launchpad-de",
	}
	c.Spec = &ClusterSpec{}

	type spec ClusterConfig
	yc := (*spec)(c)

	if err := unmarshal(yc); err != nil {
		return err
	}

	return nil
}

// Validate validates that everything in the config makes sense
// Currently we do only very "static" validation using https://github.com/go-playground/validator
func (c *ClusterConfig) Validate() error {
	validator := validator.New()
	validator.RegisterStructValidation(requireManager, ClusterSpec{})
	return validator.Struct(c)
}

func requireManager(sl validator.StructLevel) {
	hosts := sl.Current().Interface().(ClusterSpec).Hosts
	if hosts.Count(func(h *Host) bool { return h.Role == "manager" }) == 0 {
		sl.ReportError(hosts, "Hosts", "", "manager required", "")
	}
}

// Init returns an example of configuration file contents
func Init() *ClusterConfig {
	var dtrConfig *DtrConfig
	dtrConfig = &DtrConfig{
		Version:       constant.DTRVersion,
		ReplicaConfig: "sequential",
	}
	return &ClusterConfig{
		APIVersion: "launchpad.mirantis.com/v1",
		Kind:       "DockerEnterprise",
		Metadata: &ClusterMeta{
			Name: "my-ucp-cluster",
		},
		Spec: &ClusterSpec{
			Engine: EngineConfig{
				Version: constant.EngineVersion,
			},
			Ucp: UcpConfig{
				Version: constant.UCPVersion,
			},
			Dtr: dtrConfig,
			Hosts: []*Host{
				{
					Address: "10.0.0.1",
					Role:    "manager",
					SSH: &SSH{
						User:    "root",
						Port:    22,
						KeyPath: "~/.ssh/id_rsa",
					},
				},
				{
					Address: "10.0.0.2",
					Role:    "worker",
					SSH:     DefaultSSH(),
				},
				{
					Address: "10.0.0.3",
					Role:    "dtr",
					SSH:     DefaultSSH(),
				},
			},
		},
	}
}
