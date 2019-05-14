package world

import (
	"math"

	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/light"
	"github.com/kieron-pivotal/rays/material"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"
)

const (
	REFLECT_MAX_RECURSION = 5
)

type World struct {
	Objects     []*shape.Object
	LightSource *light.Point
}

func New() *World {
	w := World{}
	return &w
}

func Default() *World {
	w := New()
	lightSource := light.NewPoint(tuple.Point(-10, 10, -10), color.New(1, 1, 1))
	w.LightSource = &lightSource
	obj1 := shape.NewSphere()
	obj2 := shape.NewSphere()
	mat1 := material.New()
	mat1.Color = color.New(0.8, 1, 0.6)
	mat1.Diffuse = 0.7
	mat1.Specular = 0.2
	obj1.SetMaterial(mat1)
	obj2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5))
	w.AddObject(obj1)
	w.AddObject(obj2)
	return w
}

func (w *World) AddObject(obj *shape.Object) {
	w.Objects = append(w.Objects, obj)
}

func (w *World) Intersections(r ray.Ray) *shape.Intersections {
	ix := shape.NewIntersections()
	for _, o := range w.Objects {
		oix := o.Intersect(r)
		for i := 0; i < oix.Count(); i++ {
			intersection := oix.Get(i)
			ix.Add(intersection.T, intersection.Object)
		}
	}
	return ix
}

func (w *World) ShadeHit(comps shape.Computations, optRemaining ...int) color.Color {
	remaining := REFLECT_MAX_RECURSION
	if len(optRemaining) == 1 {
		remaining = optRemaining[0]
	}
	inShadow := w.InShadow(comps.OverPoint)
	surface := comps.Object.Material().Lighting(
		*w.LightSource, comps.Object.GetTransform(), comps.Point, comps.EyeV, comps.NormalV, inShadow)
	reflected := w.ReflectedColor(comps, remaining)
	refracted := w.RefractedColor(comps, remaining)
	return surface.Add(reflected).Add(refracted)
}

func (w *World) ColorAt(r ray.Ray, optRemaining ...int) color.Color {
	remaining := REFLECT_MAX_RECURSION
	if len(optRemaining) == 1 {
		remaining = optRemaining[0]
	}
	ix := w.Intersections(r)
	hit := ix.Hit()
	if hit != nil {
		comps := hit.PrepareComputations(r, ix)
		return w.ShadeHit(comps, remaining)
	}
	return color.Color{}
}

func (w *World) InShadow(p tuple.Tuple) bool {
	pointToLight := w.LightSource.Position.Subtract(p)
	distance := pointToLight.Magnitude()
	ray := ray.New(p, pointToLight.Normalize())
	ix := w.Intersections(ray)
	hit := ix.Hit()
	return hit != nil && hit.T < distance
}

func (w *World) ReflectedColor(comps shape.Computations, remaining int) color.Color {
	m := comps.Object.Material()
	if m.Reflective < tuple.EPSILON || remaining == 0 {
		return color.Color{}
	}

	reflectRay := ray.New(comps.OverPoint, comps.ReflectV)
	color := w.ColorAt(reflectRay, remaining-1)
	return color.Multiply(m.Reflective)
}

func (w *World) RefractedColor(comps shape.Computations, remaining int) color.Color {
	m := comps.Object.Material()
	if m.Transparency == 0.0 || remaining == 0 {
		return color.New(0, 0, 0)
	}

	nRatio := comps.N1 / comps.N2
	cosI := comps.EyeV.Dot(comps.NormalV)
	sin2T := nRatio * nRatio * (1 - cosI*cosI)
	if sin2T > 1 {
		return color.New(0, 0, 0)
	}

	cosT := math.Sqrt(1 - sin2T)
	direction := comps.NormalV.Multiply(nRatio*cosI - cosT).Subtract(comps.EyeV.Multiply(nRatio))
	refractRay := ray.New(comps.UnderPoint, direction)
	return w.ColorAt(refractRay, remaining-1).Multiply(comps.Object.Material().Transparency)
}
