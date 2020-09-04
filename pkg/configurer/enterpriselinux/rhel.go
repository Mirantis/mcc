package enterpriselinux

import (
	"github.com/Mirantis/mcc/pkg/api"
	"github.com/Mirantis/mcc/pkg/configurer"
)

// Rhel RedHat Enterprise Linux
type Rhel struct {
	Configurer
}

func resolveRedhatConfigurer(h *api.Host) api.HostConfigurer {
	if h.Metadata.Os.ID == "rhel" {
		return &Rhel{
			Configurer: Configurer{
				LinuxConfigurer: configurer.LinuxConfigurer{
					Host: h,
				},
			},
		}
	}

	return nil
}

func init() {
	api.RegisterHostConfigurer(resolveRedhatConfigurer)
}
