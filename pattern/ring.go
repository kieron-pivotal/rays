package pattern

import (
	"math"

	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/tuple"
)

type Ring struct {
	A color.Color
	B color.Color
}

func (r Ring) PatternAt(p tuple.Tuple) color.Color {
	dist := math.Sqrt(p.X*p.X + p.Z*p.Z)
	if int(math.Floor(dist))%2 == 0 {
		return r.A
	}
	return r.B
}

func NewRing(a, b color.Color) Pattern {
	return New(Ring{A: a, B: b})
}
