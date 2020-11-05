package ucp

import (
	"github.com/Mirantis/mcc/pkg/api"
	"github.com/Mirantis/mcc/pkg/phase"
	log "github.com/sirupsen/logrus"
)

// Reset - reinstall
func (u *UCP) Reset() error {
	log.Debugf("loaded cluster cfg: %+v", u.ClusterConfig)

	phaseManager := phase.NewManager(&u.ClusterConfig)

	phaseManager.AddPhases(
		&phase.Connect{},
		&phase.GatherFacts{},
		&phase.RunHooks{Stage: "Before", Action: "Reset", StepListFunc: func(h *api.Host) *[]string { return h.Hooks.Reset.Before }},
		// begin DTR phases
		&phase.UninstallDTR{},
		// end DTR phases
		&phase.UninstallUCP{},
		&phase.DownloadInstaller{},
		&phase.UninstallEngine{},
		&phase.CleanUp{},
		&phase.RunHooks{Stage: "After", Action: "Reset", StepListFunc: func(h *api.Host) *[]string { return h.Hooks.Reset.After }},
		&phase.Disconnect{},
	)

	return phaseManager.Run()
}
