package pattern

import (
	"math"

	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/tuple"
)

type Checker struct {
	A color.Color
	B color.Color
}

func (c Checker) PatternAt(p tuple.Tuple) color.Color {
	md := int64(math.Floor(p.X)) + int64(math.Floor(p.Y)) + int64(math.Floor(p.Z))
	if md%2 == 0 {
		return c.A
	}
	return c.B
}

func NewChecker(a, b color.Color) Pattern {
	return New(Checker{A: a, B: b})
}
