package pattern

import (
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/tuple"
)

type Pattern struct {
	actualPattern ActualPattern
	transform     matrix.Matrix
	inverseTransform matrix.Matrix
}

//go:generate counterfeiter -o fakes/fake_actual_pattern.go . ActualPattern

type ActualPattern interface {
	PatternAt(p tuple.Tuple) color.Color
}

//go:generate counterfeiter -o fakes/fake_inv_transform_getter.go . InvTransformGetter

type InvTransformGetter interface {
	GetInverseTransform() matrix.Matrix
}

func New(actualPattern ActualPattern) Pattern {
	return Pattern{
		actualPattern: actualPattern,
		transform:     matrix.Identity(4, 4),
		inverseTransform : matrix.Identity(4, 4),
	}
}

func (p Pattern) GetTransform() matrix.Matrix {
	return p.transform
}

func (p *Pattern) SetTransform(t matrix.Matrix) {
	p.transform = t
	p.inverseTransform = t.Inverse()
}

func (p Pattern) PatternAtShape(obj InvTransformGetter, wp tuple.Tuple) color.Color {
	op := obj.GetInverseTransform().TupleMultiply(wp)
	pp := p.inverseTransform.TupleMultiply(op)
	return p.actualPattern.PatternAt(pp)
}
