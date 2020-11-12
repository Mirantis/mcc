package phase

import (
	"fmt"
	"strings"

	"github.com/Mirantis/mcc/pkg/api"
	"github.com/Mirantis/mcc/pkg/dtr"
	"github.com/Mirantis/mcc/pkg/exec"
	"github.com/Mirantis/mcc/pkg/phase"
	log "github.com/sirupsen/logrus"
)

// JoinDtrReplicas phase implementation
type JoinDtrReplicas struct {
	phase.Analytics
	DtrPhase
}

// Title for the phase
func (p *JoinDtrReplicas) Title() string {
	return "Join DTR Replicas"
}

// Run joins all the workers nodes to swarm if not already part of it.
func (p *JoinDtrReplicas) Run() error {
	dtrLeader := p.Config.Spec.DtrLeader()
	ucpFlags := dtr.BuildUCPFlags(p.Config)
	sequentialInt := 0

	for _, d := range p.Config.Spec.Dtrs() {
		sequentialInt++
		// Iterate through the Dtrs and determine which have DTR installed
		// on them, if one is found which is not yet in the cluster, perform
		// a join against dtrLeader
		if api.IsDtrInstalled(d) {
			log.Infof("%s: already a DTR node", d)
			continue
		}

		// Run the join with the appropriate flags taken from the install spec
		runFlags := []string{"--rm", "-i"}
		if dtrLeader.Configurer.SELinuxEnabled() {
			runFlags = append(runFlags, "--security-opt label=disable")
		}
		joinFlags := []string{
			fmt.Sprintf("--ucp-node %s", d.Metadata.LongHostname),
			fmt.Sprintf("--existing-replica-id %s", p.Config.Spec.Dtr.Metadata.DtrLeaderReplicaID),
		}
		if p.Config.Spec.Dtr.ReplicaConfig == "sequential" {
			// Assign the appropriate sequential replica value if set
			builtSeqInt := dtr.SequentialReplicaID(sequentialInt)
			log.Debugf("Joining replica with sequential replicaID: %s", builtSeqInt)
			joinFlags = append(joinFlags, fmt.Sprintf("--replica-id %s", builtSeqInt))
		}
		joinFlags = append(joinFlags, ucpFlags...)
		// We can't just append the installFlags to joinFlags because they
		// differ, so we have to selectively pluck the ones that are shared
		for _, f := range dtr.PluckSharedInstallFlags(p.Config.Spec.Dtr.InstallFlags, dtr.SharedInstallJoinFlags) {
			joinFlags = append(joinFlags, f)
		}

		joinCmd := dtrLeader.Configurer.DockerCommandf("run %s %s join %s", strings.Join(runFlags, " "), p.Config.Spec.Dtr.Metadata.InstalledBootstrapImage, strings.Join(joinFlags, " "))
		log.Debugf("%s: Joining DTR replica to cluster", d)
		err := dtrLeader.Exec(joinCmd, exec.StreamOutput())
		if err != nil {
			return fmt.Errorf("%s: failed to run DTR join: %s", dtrLeader, err.Error())
		}
	}
	return nil
}