package ray

import (
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/tuple"
)

type Ray struct {
	Origin    tuple.Tuple
	Direction tuple.Tuple
}

func New(origin, direction tuple.Tuple) Ray {
	r := Ray{
		Origin:    origin,
		Direction: direction,
	}
	return r
}

func (r Ray) Position(t float64) tuple.Tuple {
	return r.Origin.Add(r.Direction.Multiply(t))
}

func (r Ray) Transform(t matrix.Matrix) Ray {
	return Ray{
		Origin:    t.TupleMultiply(r.Origin),
		Direction: t.TupleMultiply(r.Direction),
	}
}
