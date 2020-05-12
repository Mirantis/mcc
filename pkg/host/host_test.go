package host

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHost_SwarmAddress(t *testing.T) {

	h := Host{
		Address: "1.2.3.4",
	}

	require.Equal(t, "1.2.3.4:2377", h.SwarmAddress())
}
