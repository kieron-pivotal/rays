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
	id                 int64
	transform          matrix.Matrix
	inverseTransform   matrix.Matrix
	transposeTransform matrix.Matrix
	material           material.Material
	localObject        LocalObject
}

func New(obj LocalObject) *Object {
	o := Object{
		id:                 GetNextCounter(),
		transform:          matrix.Identity(4, 4),
		inverseTransform:   matrix.Identity(4, 4),
		transposeTransform: matrix.Identity(4, 4),
		material:           material.New(),
		localObject:        obj,
	}
	return &o
}

func (o *Object) Intersect(ray ray.Ray) *Intersections {
	ray2 := ray.Transform(o.inverseTransform)
	res := NewIntersections()
	for _, t := range o.localObject.LocalIntersect(ray2) {
		res.Add(t, o)
	}
	return res
}

func (o *Object) NormalAt(p tuple.Tuple) tuple.Tuple {
	objPoint := o.inverseTransform.TupleMultiply(p)
	objNormal := o.localObject.LocalNormalAt(objPoint)
	worldNormal := o.transposeTransform.TupleMultiply(objNormal)
	worldNormal.W = 0
	return worldNormal.Normalize()
}

func (o *Object) SetTransform(t matrix.Matrix) {
	o.transform = t
	o.inverseTransform = t.Inverse()
	o.transposeTransform = o.inverseTransform.Transpose()
}

func (o *Object) GetTransform() matrix.Matrix {
	return o.transform
}

func (o *Object) GetInverseTransform() matrix.Matrix {
	return o.inverseTransform
}

func (o *Object) Material() material.Material {
	return o.material
}

func (o *Object) SetMaterial(m material.Material) {
	o.material = m
}
