package msr2

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	common "github.com/Mirantis/mcc/pkg/product/common/api"
	"github.com/Mirantis/mcc/pkg/product/mke/api"
)

func TestPluckSharedInstallFlags(t *testing.T) {

	t.Run("Install flags are shared with join", func(t *testing.T) {
		installFlags := []string{
			"--replica-http-port 8000",
			"--replica-https-port 4443",
			"--ucp-insecure-tls",
			"--nfs-storage-url nfs://nfs.acme.com",
		}
		expectedJoinFlags := []string{
			"--replica-http-port 8000",
			"--replica-https-port 4443",
			"--ucp-insecure-tls",
		}
		actualJoinFlags := PluckSharedInstallFlags(installFlags, SharedInstallJoinFlags)
		sort.Strings(actualJoinFlags)
		sort.Strings(expectedJoinFlags)
		if !reflect.DeepEqual(actualJoinFlags, expectedJoinFlags) {
			t.Fatalf("expected is not equal to actual\nexpected: %s\nactual: %s", expectedJoinFlags, actualJoinFlags)
		}
	})

	t.Run("Flags with multiple values can still be plucked", func(t *testing.T) {
		multiValueFlag := []string{
			"--fake-flag one two three",
		}
		testSharedFlags := []string{
			"--fake-flag",
		}
		expected := []string{
			"--fake-flag one two three",
		}
		actual := PluckSharedInstallFlags(multiValueFlag, testSharedFlags)
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("expected is not equal to actual\nexpected: %s\nactual: %s", expected, actual)
		}
	})
}

func TestBuildMKEFlags(t *testing.T) {
	config := &api.ClusterConfig{
		Spec: &api.ClusterSpec{
			MKE: &api.MKEConfig{
				AdminUsername: "admin",
				AdminPassword: "password1234",
				InstallFlags: []string{
					"--san ucp.acme.com",
				},
			},
			MSR2: &api.MSR2Config{},
		},
	}

	t.Run("MKE flags are built when --san is provided", func(t *testing.T) {
		actual := BuildMKEFlags(config)
		expected := common.Flags{
			"--ucp-url=\"ucp.acme.com\"",
			"--ucp-username=\"admin\"",
			"--ucp-password=\"password1234\"",
		}
		sort.Strings(actual)
		sort.Strings(expected)
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("expected is not equal to actual\nexpected: %s\nactual: %s", expected, actual)
		}
	})
}

func TestFormatReplicaID(t *testing.T) {
	require.Equal(t, "000000000001", FormatReplicaID(1))
	require.Equal(t, "00000000000a", FormatReplicaID(10))
	require.Equal(t, "000000000010", FormatReplicaID(16))
}

func TestSequentialReplicaIDs(t *testing.T) {
	config := &api.ClusterConfig{
		Spec: &api.ClusterSpec{
			Hosts: []*api.Host{
				{Role: "msr2"},
				{Role: "msr2", MSR2Metadata: &api.MSR2Metadata{
					ReplicaID: "00000000001f",
				}},
				{Role: "msr2"},
			},
			MSR2: &api.MSR2Config{ReplicaIDs: "sequential"},
		},
	}
	require.NoError(t, AssignSequentialReplicaIDs(config))
	require.Equal(t, "000000000020", config.Spec.Hosts[0].MSR2Metadata.ReplicaID)
	require.Equal(t, "00000000001f", config.Spec.Hosts[1].MSR2Metadata.ReplicaID)
	require.Equal(t, "000000000021", config.Spec.Hosts[2].MSR2Metadata.ReplicaID)
}