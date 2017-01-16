package stats

import "math"

// Data represents a slice of data points
type Data []float64

// Min returns minimum of the data.
// If the data slice is empty, it returns 0.0.
func (d Data) Min() float64 {
	if len(d) == 0 {
		return 0.0
	}

	min := d[0]
	for _, v := range d[1:] {
		min = math.Min(min, v)
	}
	return min
}

// Max returns minimum of the data.
// If the data slice is empty, it returns 0.0.
func (d Data) Max() float64 {
	if len(d) == 0 {
		return 0.0
	}

	min := d[0]
	for _, v := range d[1:] {
		min = math.Max(min, v)
	}
	return min
}
