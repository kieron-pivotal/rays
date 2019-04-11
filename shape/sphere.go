package shape

import (
	"math"

	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/tuple"
)

type Sphere struct{}

func NewSphere() *Sphere {
	return &Sphere{}
}

func (s *Sphere) Intersect(ray ray.Ray) *Intersections {
	sphere_to_ray := ray.Origin.Subtract(tuple.Point(0, 0, 0))

	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphere_to_ray)
	c := sphere_to_ray.Dot(sphere_to_ray) - 1
	discriminant := b*b - 4*a*c

	res := NewIntersections()
	if discriminant >= 0 {
		res.Add((-b-math.Sqrt(discriminant))/(2*a), s)
		res.Add((-b+math.Sqrt(discriminant))/(2*a), s)
	}
	return res
}
