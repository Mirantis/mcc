package v12

import (
	"github.com/Mirantis/mcc/pkg/config/migration"
	log "github.com/sirupsen/logrus"
)

// Migrate migrates an v1 format configuration into the v1.1 api format and replaces the contents of the supplied data byte slice
func Migrate(plain map[string]interface{}) error {
	plain["apiVersion"] = "launchpad.mirantis.com/mke/v1.3"

	if spec, ok := plain["spec"].(map[interface{}]interface{}); ok {
		if hosts, ok := spec["hosts"]; ok {
			hslice := hosts.([]interface{})
			for _, h := range hslice {
				host, ok := h.(map[interface{}]interface{})
				if ok {
					if addr, ok := host["address"].(string); ok {
						if ssh, ok := host["ssh"].(map[interface{}]interface{}); ok {
							ssh["address"] = addr
							log.Debugf("migrated v1.2 spec.hosts[*].address to v1.3 spec.hosts[*].ssh.address")
						} else if winrm, ok := host["winRM"].(map[interface{}]interface{}); ok {
							winrm["address"] = addr
							log.Debugf("migrated v1.2 spec.hosts[*].address to v1.3 spec.hosts[*].winrm.address")
						} else {
							ssh := map[string]interface{}{"address": addr}
							host["ssh"] = ssh
							log.Debugf("migrated v1.2 spec.hosts[*].address to v1.3 spec.hosts[*].ssh.address")
						}
						delete(host, "address")
					}
					if lh, ok := host["localhost"].(bool); ok && lh {
						local := map[string]interface{}{"enabled": true}
						host["localhost"] = local
						delete(host, "address")
						delete(host, "ssh")
						delete(host, "winRM")
						log.Debugf("migrated v1.2 spec.hosts[*].localhost to v1.3 spec.hosts[*].localhost.enabled")
					}
				}
			}
		}
	}

	log.Debugf("migrated configuration from launchpad.mirantis.com/v1.1 to launchpad.mirantis.com/mke/v1.2")
	log.Infof("Note: The configuration has been migrated from a previous version")
	log.Infof("      to see the migrated configuration use: launchpad describe config")
	return nil
}

func init() {
	migration.Register("launchpad.mirantis.com/mke/v1.2", Migrate)
}
