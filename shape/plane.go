package shape

import (
	"math"

	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/tuple"
)

type Plane struct {
}

func NewPlane() *Object {
	return New(Plane{})
}

func (p Plane) Name() string {
	return "Plane"
}

func (p Plane) LocalIntersect(r ray.Ray) []float64 {
	if math.Abs(r.Direction.Y) < tuple.EPSILON {
		return []float64{}
	}

	return []float64{-r.Origin.Y / r.Direction.Y}
}

func (p Plane) LocalNormalAt(tuple.Tuple) tuple.Tuple {
	return tuple.Vector(0, 1, 0)
}
