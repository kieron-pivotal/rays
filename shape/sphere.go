package shape

import (
	"math"

	"github.com/kieron-pivotal/rays/material"
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/tuple"
)

type Sphere struct{}

func NewSphere() *Object {
	return New(Sphere{})
}

func NewGlassSphere() *Object {
	s := New(Sphere{})
	m := material.New()
	m.Transparency = 1.0
	m.RefractiveIndex = 1.5
	s.SetMaterial(m)
	return s
}

func (s Sphere) Name() string {
	return "Unit sphere"
}

func (s Sphere) LocalIntersect(ray ray.Ray) []float64 {
	sphereToRay := ray.Origin.Subtract(tuple.Point(0, 0, 0))

	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c

	res := []float64{}
	if discriminant >= 0 {
		res = append(res, (-b-math.Sqrt(discriminant))/(2*a))
		res = append(res, (-b+math.Sqrt(discriminant))/(2*a))
	}
	return res
}

func (s Sphere) LocalNormalAt(p tuple.Tuple) tuple.Tuple {
	return p.Subtract(tuple.Point(0, 0, 0))
}
