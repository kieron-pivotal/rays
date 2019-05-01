package world

import (
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/light"
	"github.com/kieron-pivotal/rays/material"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"
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

func (w *World) ShadeHit(comps shape.Computations) color.Color {
	inShadow := w.InShadow(comps.OverPoint)
	return comps.Object.Material().Lighting(*w.LightSource, comps.Object.GetTransform(), comps.Point, comps.EyeV, comps.NormalV, inShadow)
}

func (w *World) ColorAt(r ray.Ray) color.Color {
	ix := w.Intersections(r)
	hit := ix.Hit()
	if hit != nil {
		comps := hit.PrepareComputations(r)
		return w.ShadeHit(comps)
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
