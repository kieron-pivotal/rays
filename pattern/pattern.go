package pattern

import (
	"math"

	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/tuple"
)

type Stripe struct {
	A         color.Color
	B         color.Color
	transform matrix.Matrix
}

func NewStripe(a, b color.Color) Stripe {
	return Stripe{
		A:         a,
		B:         b,
		transform: matrix.Identity(4, 4),
	}
}

func (s Stripe) StripeAt(p tuple.Tuple) color.Color {
	if int(math.Floor(p.X))%2 == 0 {
		return s.A
	}
	return s.B
}

func (s Stripe) StripeAtObject(t matrix.Matrix, p tuple.Tuple) color.Color {
	op := t.Inverse().TupleMultiply(p)
	pp := s.transform.Inverse().TupleMultiply(op)
	return s.StripeAt(pp)
}

func (s *Stripe) SetTransform(t matrix.Matrix) {
	s.transform = t
}
