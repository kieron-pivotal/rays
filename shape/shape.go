package shape

import (
	"sync/atomic"

	"github.com/kieron-pivotal/rays/material"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/ray"
)

type Object interface {
	Intersect(ray ray.Ray) *Intersections
	SetTransform(t matrix.Matrix)
	Material() material.Material
}

var counter int64

func GetNextCounter() int64 {
	return atomic.AddInt64(&counter, 1)
}
