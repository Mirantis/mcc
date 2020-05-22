package v1beta1

// ClusterSpec defines cluster spec
type ClusterSpec struct {
	Hosts  Hosts        `yaml:"hosts" validate:"required,dive,min=1"`
	Ucp    UcpConfig    `yaml:"ucp"`
	Engine EngineConfig `yaml:"engine"`
}

// Workers filters only the workers from the cluster config
func (c *ClusterSpec) Workers() []*Host {
	workers := make([]*Host, 0)
	for _, h := range c.Hosts {
		if h.Role == "worker" {
			workers = append(workers, h)
		}
	}

	return workers
}

// Managers filters only the manager nodes from the cluster config
func (c *ClusterSpec) Managers() []*Host {
	managers := make([]*Host, 0)
	for _, h := range c.Hosts {
		if h.Role == "manager" {
			managers = append(managers, h)
		}
	}

	return managers
}

// UnmarshalYAML sets in some sane defaults when unmarshaling the data from yaml
func (c *ClusterSpec) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type rawConfig ClusterSpec
	raw := rawConfig{
		Engine: NewEngineConfig(),
		Ucp:    NewUcpConfig(),
	}

	if err := unmarshal(&raw); err != nil {
		return err
	}

	*c = ClusterSpec(raw)
	return nil
}