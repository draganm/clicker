package graph

import (
	"fmt"
	"strings"

	"github.com/draganm/clicker/stats"
)

type Point struct {
	X float64
	Y float64
}

type Line []Point

func (l Line) String() string {
	parts := make([]string, len(l))
	for i, p := range l {
		parts[i] = fmt.Sprintf("%.2f,%.2f", p.X, p.Y)
	}
	return strings.Join(parts, " ")
}

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
