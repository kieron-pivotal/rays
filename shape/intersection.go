package shape

import "sort"

type Intersections struct {
	list []*Intersection
}

type Intersection struct {
	T     float64
	Shape Object
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

func (i *Intersections) Add(t float64, s *Sphere) {
	i.list = append(i.list, &Intersection{T: t, Shape: s})
}

func (i *Intersections) Hit() *Intersection {
	sort.Slice(i.list, func(a, b int) bool {
		return i.list[a].T < i.list[b].T
	})
	for _, x := range i.list {
		if x.T >= 0 {
			return x
		}
	}
	return nil
}
