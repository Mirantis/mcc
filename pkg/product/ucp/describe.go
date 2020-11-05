package ucp

import (
	"github.com/Mirantis/mcc/pkg/phase"
	log "github.com/sirupsen/logrus"
)

// Describe - gets information about configured instance
func (u *UCP) Describe(reportName string) error {
	var dtr bool
	var ucp bool

	if reportName == "dtr" {
		dtr = true
	}

	if reportName == "ucp" {
		ucp = true
	}

	log.Debugf("loaded cluster cfg: %+v", u.ClusterConfig)

	phaseManager := phase.NewManager(&u.ClusterConfig)
	phaseManager.IgnoreErrors = true

	phaseManager.AddPhases(
		&phase.Connect{},
		&phase.GatherFacts{},
		&phase.Disconnect{},
		&phase.Describe{Ucp: ucp, Dtr: dtr},
	)

	return phaseManager.Run()
}
