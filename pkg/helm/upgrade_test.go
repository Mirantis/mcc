package helm

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/utils/ptr"
)

func TestUpgrade(t *testing.T) {
	h := NewHelmTestClient(t)

	t.Run("Upgrade success", func(t *testing.T) {
		rd, uninstallFunc := InstallRethinkDBOperatorChart(t, h)
		t.Cleanup(uninstallFunc)
		rd.Values = map[string]interface{}{
			"controllerManager": map[string]interface{}{
				"manager": map[string]interface{}{
					"tag": "1.0.1",
				},
			},
		}

		rel, err := h.Upgrade(context.Background(), &Options{
			ReleaseDetails: rd,
			Timeout:        ptr.To(DefaultTimeout),
		})
		if assert.NoError(t, err) {
			assert.Equal(t, rd.ChartName, rel.Chart.Metadata.Name)
			assert.ObjectsAreEqualValues(rd.Values, rel.Chart.Values)
		}
	})

	t.Run("Upgrade, reuse values", func(t *testing.T) {
		rd, uninstallFunc := InstallRethinkDBOperatorChart(t, h)
		t.Cleanup(uninstallFunc)

		rd.Values = map[string]interface{}{
			"rethinkdb": map[string]interface{}{
				"name": "cooler-rethinkdb",
			},
		}

		rel, err := h.Upgrade(context.Background(), &Options{
			ReleaseDetails: rd,
			ReuseValues:    true,
			Timeout:        ptr.To(DefaultTimeout),
		})
		if assert.NoError(t, err) {
			assert.Equal(t, rd.ChartName, rel.Chart.Metadata.Name)
			// ReuseValues should not change the values, but reuse the previous
			// ones.
			assert.NotEqual(t, "cooler-rethinkdb", rel.Chart.Values["rethinkdb"].(map[string]interface{})["name"])
			assert.ObjectsAreEqualValues(rd.Values, rel.Chart.Values)
		}
	})

	t.Run("Upgrade failure, empty release details", func(t *testing.T) {
		_, err := h.Upgrade(context.Background(), &Options{
			ReleaseDetails: ReleaseDetails{},
			Timeout:        ptr.To(DefaultTimeout),
		})
		assert.ErrorContains(t, err, "failed to upgrade")
	})

}