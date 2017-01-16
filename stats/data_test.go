package stats_test

import (
	"testing"

	"github.com/draganm/clicker/stats"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	cases := []struct {
		data     stats.Data
		expected float64
	}{
		{
			stats.Data{},
			0.0,
		},
		{
			stats.Data{1.0, 2.0, 0.0},
			0.0,
		},
		{
			stats.Data{1.0},
			1.0,
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.expected, c.data.Min(), "Case %d failed", i+1)
	}

}

func TestMax(t *testing.T) {
	cases := []struct {
		data     stats.Data
		expected float64
	}{
		{
			stats.Data{},
			0.0,
		},
		{
			stats.Data{1.0, 2.0, 0.0},
			2.0,
		},
		{
			stats.Data{1.0},
			1.0,
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.expected, c.data.Max(), "Case %d failed", i+1)
	}

}
