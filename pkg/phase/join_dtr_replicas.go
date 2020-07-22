package phase

import (
	"fmt"
	"strings"

	api "github.com/Mirantis/mcc/pkg/apis/v1beta2"
	"github.com/Mirantis/mcc/pkg/dtr"
	log "github.com/sirupsen/logrus"
)

// JoinDtrReplicas phase implementation
type JoinDtrReplicas struct {
	Analytics
}

// Title for the phase
func (p *JoinDtrReplicas) Title() string {
	return "Join DTR Replicas"
}

// Run joins all the workers nodes to swarm if not already part of it.
func (p *JoinDtrReplicas) Run(config *api.ClusterConfig) error {
	dtrLeader := config.Spec.DtrLeader()
	ucpFlags := dtr.BuildUcpFlags(config)
	sequentialInt := 0

	for _, d := range config.Spec.Dtrs() {
		sequentialInt++
		// Iterate through the Dtrs and determine which have DTR installed
		// on them, if one is found which is not yet in the cluster, perform
		// a join against dtrLeader
		if api.IsDtrInstalled(d) {
			log.Infof("%s: already a DTR node", d.Address)
			continue
		}

		// Run the join with the appropriate flags taken from the install spec
		runFlags := []string{"--rm", "-i"}
		if dtrLeader.Configurer.SELinuxEnabled() {
			runFlags = append(runFlags, "--security-opt label=disable")
		}
		joinFlags := []string{
			fmt.Sprintf("--ucp-node %s", d.Metadata.Hostname),
			fmt.Sprintf("--existing-replica-id %s", config.Spec.Dtr.Metadata.DtrLeaderReplicaID),
		}
		if config.Spec.Dtr.ReplicaConfig == "sequential" {
			// Assign the appropriate sequential replica value if set
			builtSeqInt := dtr.SequentialReplicaID(sequentialInt)
			log.Debugf("Joining replica with sequential replicaID: %s", builtSeqInt)
			joinFlags = append(joinFlags, fmt.Sprintf("--replica-id %s", builtSeqInt))
		}
		joinFlags = append(joinFlags, ucpFlags...)
		// We can't just append the installFlags to joinFlags because they
		// differ, so we have to selectively pluck the ones that are shared
		for _, f := range dtr.PluckSharedInstallFlags(config.Spec.Dtr.InstallFlags, dtr.SharedInstallJoinFlags) {
			joinFlags = append(joinFlags, f)
		}

		joinCmd := dtrLeader.Configurer.DockerCommandf("run %s %s join %s", strings.Join(runFlags, " "), config.Spec.Dtr.GetBootstrapperImage(), strings.Join(joinFlags, " "))
		log.Debugf("%s: Joining DTR replica to cluster", d.Address)
		err := dtrLeader.ExecCmd(joinCmd, "", true, false)
		if err != nil {
			return NewError("Failed to run DTR join")
		}
	}
	return nil
}