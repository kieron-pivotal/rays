package matrix

import (
	"math"

	"github.com/kieron-pivotal/rays/tuple"
)

func Translation(x, y, z float64) Matrix {
	m := Identity(4, 4)
	m.Set(0, 3, x)
	m.Set(1, 3, y)
	m.Set(2, 3, z)
	return m
}

func (m Matrix) Translate(x, y, z float64) Matrix {
	t := Translation(x, y, z)
	return t.Multiply(m)
}

func Scaling(x, y, z float64) Matrix {
	m := Identity(4, 4)
	m.Set(0, 0, x)
	m.Set(1, 1, y)
	m.Set(2, 2, z)
	return m
}

func (m Matrix) Scale(x, y, z float64) Matrix {
	t := Scaling(x, y, z)
	return t.Multiply(m)
}

func RotationX(a float64) Matrix {
	m := Identity(4, 4)
	m.Set(1, 1, math.Cos(a))
	m.Set(1, 2, -math.Sin(a))
	m.Set(2, 1, math.Sin(a))
	m.Set(2, 2, math.Cos(a))
	return m
}

func (m Matrix) RotateX(a float64) Matrix {
	t := RotationX(a)
	return t.Multiply(m)
}

func RotationY(a float64) Matrix {
	m := Identity(4, 4)
	m.Set(0, 0, math.Cos(a))
	m.Set(0, 2, math.Sin(a))
	m.Set(2, 0, -math.Sin(a))
	m.Set(2, 2, math.Cos(a))
	return m
}

func (m Matrix) RotateY(a float64) Matrix {
	t := RotationY(a)
	return t.Multiply(m)
}

func RotationZ(a float64) Matrix {
	m := Identity(4, 4)
	m.Set(0, 0, math.Cos(a))
	m.Set(0, 1, -math.Sin(a))
	m.Set(1, 0, math.Sin(a))
	m.Set(1, 1, math.Cos(a))
	return m
}

func (m Matrix) RotateZ(a float64) Matrix {
	t := RotationZ(a)
	return t.Multiply(m)
}

func Shearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	m := Identity(4, 4)
	m.Set(0, 1, xy)
	m.Set(0, 2, xz)
	m.Set(1, 0, yx)
	m.Set(1, 2, yz)
	m.Set(2, 0, zx)
	m.Set(2, 1, zy)
	return m
}

func (m Matrix) Shear(xy, xz, yx, yz, zx, zy float64) Matrix {
	t := Shearing(xy, xz, yx, yz, zx, zy)
	return t.Multiply(m)
}

func ViewTransformation(from, to, up tuple.Tuple) Matrix {
	forwardNormal := to.Subtract(from).Normalize()
	upNormal := up.Normalize()
	left := forwardNormal.Cross(upNormal)
	trueUp := left.Cross(forwardNormal)

	orientation := New(4, 4,
		left.X, left.Y, left.Z, 0,
		trueUp.X, trueUp.Y, trueUp.Z, 0,
		-forwardNormal.X, -forwardNormal.Y, -forwardNormal.Z, 0,
		0, 0, 0, 1,
	)
	return orientation.Multiply(Translation(-from.X, -from.Y, -from.Z))
}
