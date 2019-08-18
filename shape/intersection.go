package shape

import (
	"container/list"
	"math"
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

type Computations struct {
	T          float64
	Object     *Object
	Point      tuple.Tuple
	OverPoint  tuple.Tuple
	UnderPoint tuple.Tuple
	EyeV       tuple.Tuple
	NormalV    tuple.Tuple
	ReflectV   tuple.Tuple
	N1         float64
	N2         float64
	Inside     bool
}

func (c Computations) Schlick() float64 {
	cos := c.EyeV.Dot(c.NormalV)

	if c.N1 > c.N2 {
		n := c.N1 / c.N2
		sin2T := n * n * (1 - cos*cos)
		if sin2T > 1 {
			return 1
		}
		cos = math.Sqrt(1 - sin2T)
	}
	r0 := math.Pow((c.N1-c.N2)/(c.N1+c.N2), 2)
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}

func NewIntersections() *Intersections {
	return &Intersections{
		list: make([]*Intersection, 0, 2),
	}
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

func (i *Intersection) PrepareComputations(r ray.Ray, xs *Intersections) Computations {
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
	c.UnderPoint = c.Point.Add(c.NormalV.Multiply(-tuple.EPSILON))
	c.ReflectV = r.Direction.Reflect(c.NormalV)

	containers := list.New()
	for j := 0; j < xs.Count(); j++ {
		x := xs.Get(j)
		if x == i {
			if containers.Len() == 0 {
				c.N1 = 1.0
			} else {
				c.N1 = containers.Back().Value.(*Object).Material().RefractiveIndex
			}
		}

		if el := find(x.Object, containers); el != nil {
			containers.Remove(el)
		} else {
			containers.PushBack(x.Object)
		}

		if x == i {
			if containers.Len() == 0 {
				c.N2 = 1.0
			} else {
				c.N2 = containers.Back().Value.(*Object).Material().RefractiveIndex
			}
			break
		}
	}

	return c
}

func find(obj *Object, containers *list.List) *list.Element {
	for e := containers.Front(); e != nil; e = e.Next() {
		if e.Value.(*Object) == obj {
			return e
		}
	}
	return nil
}
