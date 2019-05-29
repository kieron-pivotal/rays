package shape

import (
	"math"

	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/tuple"
)

type Cube struct {
}

func (c Cube) Name() string {
	return "Cube"
}

func (c Cube) LocalIntersect(r ray.Ray) []float64 {
	xtmin, xtmax := c.checkAxis(r.Origin.X, r.Direction.X)
	ytmin, ytmax := c.checkAxis(r.Origin.Y, r.Direction.Y)
	ztmin, ztmax := c.checkAxis(r.Origin.Z, r.Direction.Z)
	tmin := math.Max(xtmin, math.Max(ytmin, ztmin))
	tmax := math.Min(xtmax, math.Min(ytmax, ztmax))

	if tmin > tmax {
		return []float64{}
	}
	return []float64{tmin, tmax}
}

func (c Cube) checkAxis(origin, direction float64) (tmin, tmax float64) {
	inf := 1e100
	tmin_numerator := -1 - origin
	tmax_numerator := 1 - origin

	if math.Abs(direction) > tuple.EPSILON {
		tmin = tmin_numerator / direction
		tmax = tmax_numerator / direction
	} else {
		tmin = tmin_numerator * inf
		tmax = tmax_numerator * inf
	}

	if tmin > tmax {
		tmin, tmax = tmax, tmin
	}
	return
}

func (c Cube) LocalNormalAt(p tuple.Tuple) tuple.Tuple {
	maxc := math.Max(math.Abs(p.X), math.Max(math.Abs(p.Y), math.Abs(p.Z)))
	if maxc == math.Abs(p.X) {
		return tuple.Vector(p.X, 0, 0)
	}
	if maxc == math.Abs(p.Y) {
		return tuple.Vector(0, p.Y, 0)
	}
	return tuple.Vector(0, 0, p.Z)
}
