package graph

import "github.com/draganm/clicker/stats"

type Point struct {
	X float64
	Y float64
}

type Line []Point

// type Graph struct {
// 	Width  float64
// 	Height float64
// 	Line
// }

func ToLine(d stats.Data, width, height float64) Line {
	n := d.Normalize()
	bucketWidth := width / float64(len(n))

	l := Line{}

	for i, v := range n {
		x := bucketWidth * float64(i)
		y := v * float64(height)
		l = append(l, Point{x, y})
	}
	return l
}
