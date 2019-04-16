package light

import (
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/tuple"
)

type Point struct {
	Position  tuple.Tuple
	Intensity color.Color
}

func NewPoint(pos tuple.Tuple, intensity color.Color) Point {
	return Point{
		Position:  pos,
		Intensity: intensity,
	}
}
