package v1beta1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMigrationToV1Beta2Basic(t *testing.T) {
	b1 := []byte(`---
apiVersion: "launchpad.mirantis.com/v1beta1"
kind: UCP
spec:
  hosts:
    - address: "10.0.0.1"
      sshKeyPath: /tmp/tmp
      sshPort: 9022
      user: "admin"
      role: "manager"
`)
	// go's YAML marshal does not add the --- header
	b2 := []byte(`apiVersion: launchpad.mirantis.com/v1beta2
kind: UCP
spec:
  hosts:
  - address: 10.0.0.1
    role: manager
    ssh:
      keyPath: /tmp/tmp
      port: 9022
      user: admin
`)
	require.NoError(t, MigrateToV1Beta2(&b1))
	require.Equal(t, b1, b2)
}