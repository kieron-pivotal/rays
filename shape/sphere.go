package shape

import (
	"math"

	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/tuple"
)

type Sphere struct {
	id        int64
	transform matrix.Matrix
}

func NewSphere() *Sphere {
	s := Sphere{
		id:        GetNextCounter(),
		transform: matrix.Identity(4, 4),
	}
	return &s
}

func (s *Sphere) Intersect(ray ray.Ray) *Intersections {
	ray2 := ray.Transform(s.transform.Inverse())
	sphere_to_ray := ray2.Origin.Subtract(tuple.Point(0, 0, 0))

	a := ray2.Direction.Dot(ray2.Direction)
	b := 2 * ray2.Direction.Dot(sphere_to_ray)
	c := sphere_to_ray.Dot(sphere_to_ray) - 1
	discriminant := b*b - 4*a*c

	res := NewIntersections()
	if discriminant >= 0 {
		res.Add((-b-math.Sqrt(discriminant))/(2*a), s)
		res.Add((-b+math.Sqrt(discriminant))/(2*a), s)
	}
	return res
}

func (s *Sphere) SetTransform(t matrix.Matrix) {
	s.transform = t
}

func (s *Sphere) GetTransform() matrix.Matrix {
	return s.transform
}

func (s *Sphere) NormalAt(p tuple.Tuple) tuple.Tuple {
	objPoint := s.transform.Inverse().TupleMultiply(p)
	objNormal := objPoint.Subtract(tuple.Point(0, 0, 0))
	worldNormal := s.transform.Inverse().Transpose().TupleMultiply(objNormal)
	worldNormal.W = 0
	return worldNormal.Normalize()
}
