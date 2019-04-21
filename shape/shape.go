package shape

import (
	"sync/atomic"

	"github.com/kieron-pivotal/rays/material"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/tuple"
)

type Object interface {
	Intersect(ray ray.Ray) *Intersections
	SetTransform(t matrix.Matrix)
	GetTransform() matrix.Matrix
	Material() material.Material
	SetMaterial(material.Material)
	NormalAt(tuple.Tuple) tuple.Tuple
}

var counter int64

func GetNextCounter() int64 {
	return atomic.AddInt64(&counter, 1)
}
