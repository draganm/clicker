package stats_test

import (
	"testing"

	"github.com/draganm/clicker/stats"
	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	cases := []struct {
		data     stats.Data
		expected stats.Data
	}{
		{
			stats.Data{},
			stats.Data{},
		},
		{
			stats.Data{0.0},
			stats.Data{0.0},
		},
		{
			stats.Data{1.0, 2.0, 0.0},
			stats.Data{0.5, 1.0, 0.0},
		},
		{
			stats.Data{1.0, 2.0, 3.0},
			stats.Data{0.0, 0.5, 1.0},
		},
		{
			stats.Data{2.0, 4.0, -4.0},
			stats.Data{0.75, 1.0, 0.0},
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.expected, c.data.Normalize(), "Case %d failed", i+1)
	}

}
