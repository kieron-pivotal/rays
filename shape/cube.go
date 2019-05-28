package shape

import (
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/tuple"
)

type Cube struct {
}

func (c Cube) Name() string {
	return "Cube"
}

func (c Cube) LocalIntersect(ray.Ray) []float64 {
	return []float64{0}
}

func (c Cube) LocalNormalAt(tuple.Tuple) tuple.Tuple {
	return tuple.Vector(0, 0, 0)
}
