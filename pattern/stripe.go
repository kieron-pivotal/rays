package pattern

import (
	"math"

	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/tuple"
)

type Stripe struct {
	A color.Color
	B color.Color
}

func NewStripe(a, b color.Color) Pattern {
	s := Stripe{
		A: a,
		B: b,
	}
	return New(s)
}

func (s Stripe) PatternAt(p tuple.Tuple) color.Color {
	if int(math.Floor(p.X))%2 == 0 {
		return s.A
	}
	return s.B
}
