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

// Normalize returns data in the rage [0.0, 1.0].
func (d Data) Normalize() Data {

	normalized := make(Data, len(d))

	min := d.Min()
	max := d.Max()

	offset := -min
	scale := 0.0
	if min != max {
		scale = 1.0 / (max - min)
	}

	for i, v := range d {
		normalized[i] = (v + offset) * scale
	}
	return normalized
}
