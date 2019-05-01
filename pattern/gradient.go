package pattern

import (
	"math"

	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/tuple"
)

type Gradient struct {
	A color.Color
	B color.Color
}

func (g Gradient) PatternAt(p tuple.Tuple) color.Color {
	distance := g.B.Subtract(g.A)
	fraction := p.X - math.Floor(p.X)
	return g.A.Add(distance.Multiply(fraction))
}

func NewGradient(a, b color.Color) Pattern {
	return New(Gradient{
		A: a,
		B: b,
	})
}
