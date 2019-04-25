package shape

import (
	"sort"

	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/tuple"
)

type Intersections struct {
	list []*Intersection
}

type Intersection struct {
	T      float64
	Object *Object
}

func NewIntersections() *Intersections {
	return &Intersections{}
}

func (i Intersections) Count() int {
	return len(i.list)
}

func (i Intersections) Get(idx int) *Intersection {
	return i.list[idx]
}

func (i *Intersections) Add(t float64, s *Object) {
	i.list = append(i.list, &Intersection{T: t, Object: s})
	sort.Slice(i.list, func(a, b int) bool {
		return i.list[a].T < i.list[b].T
	})
}

func (i *Intersections) Hit() *Intersection {
	for _, x := range i.list {
		if x.T >= 0 {
			return x
		}
	}
	return nil
}

type Computations struct {
	T         float64
	Object    *Object
	Point     tuple.Tuple
	OverPoint tuple.Tuple
	EyeV      tuple.Tuple
	NormalV   tuple.Tuple
	Inside    bool
}

func (i *Intersection) PrepareComputations(r ray.Ray) Computations {
	c := Computations{}
	c.T = i.T
	c.Object = i.Object
	c.Point = r.Position(c.T)
	c.EyeV = r.Direction.Multiply(-1)
	c.NormalV = c.Object.NormalAt(c.Point)
	c.Inside = c.EyeV.Dot(c.NormalV) < 0
	if c.Inside {
		c.NormalV = c.NormalV.Multiply(-1)
	}
	c.OverPoint = c.Point.Add(c.NormalV.Multiply(tuple.EPSILON))
	return c
}
