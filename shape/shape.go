package shape

import (
	"sync/atomic"

	"github.com/kieron-pivotal/rays/material"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/tuple"
)

var counter int64

func GetNextCounter() int64 {
	return atomic.AddInt64(&counter, 1)
}

//go:generate counterfeiter -o fakes/fake_local_object.go . LocalObject

type LocalObject interface {
	Name() string
	LocalIntersect(ray.Ray) []float64
	LocalNormalAt(tuple.Tuple) tuple.Tuple
}

type Object struct {
	id        int64
	transform matrix.Matrix
	material  material.Material
	LocalObject
}

func New(obj LocalObject) *Object {
	o := Object{
		id:          GetNextCounter(),
		transform:   matrix.Identity(4, 4),
		material:    material.New(),
		LocalObject: obj,
	}
	return &o
}

func (o *Object) Intersect(ray ray.Ray) *Intersections {
	ray2 := ray.Transform(o.transform.Inverse())
	res := NewIntersections()
	for _, t := range o.LocalIntersect(ray2) {
		res.Add(t, o)
	}
	return res
}

func (o *Object) NormalAt(p tuple.Tuple) tuple.Tuple {
	objPoint := o.transform.Inverse().TupleMultiply(p)
	objNormal := o.LocalNormalAt(objPoint)
	worldNormal := o.transform.Inverse().Transpose().TupleMultiply(objNormal)
	worldNormal.W = 0
	return worldNormal.Normalize()
}

func (o *Object) SetTransform(t matrix.Matrix) {
	o.transform = t
}

func (o *Object) GetTransform() matrix.Matrix {
	return o.transform
}

func (o *Object) Material() material.Material {
	return o.material
}

func (o *Object) SetMaterial(m material.Material) {
	o.material = m
}
