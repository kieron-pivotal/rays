package matrix

import "math"

func Translation(x, y, z float64) Matrix {
	m := Identity(4, 4)
	m.Set(0, 3, x)
	m.Set(1, 3, y)
	m.Set(2, 3, z)
	return m
}

func (m Matrix) Translation(x, y, z float64) Matrix {
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

func (m Matrix) Scaling(x, y, z float64) Matrix {
	t := Scaling(x, y, z)
	return m.Multiply(t)
}

func RotationX(a float64) Matrix {
	m := Identity(4, 4)
	m.Set(1, 1, math.Cos(a))
	m.Set(1, 2, -math.Sin(a))
	m.Set(2, 1, math.Sin(a))
	m.Set(2, 2, math.Cos(a))
	return m
}

func (m Matrix) RotationX(a float64) Matrix {
	t := RotationX(a)
	return m.Multiply(t)
}

func RotationY(a float64) Matrix {
	m := Identity(4, 4)
	m.Set(0, 0, math.Cos(a))
	m.Set(0, 2, math.Sin(a))
	m.Set(2, 0, -math.Sin(a))
	m.Set(2, 2, math.Cos(a))
	return m
}

func (m Matrix) RotationY(a float64) Matrix {
	t := RotationY(a)
	return m.Multiply(t)
}

func RotationZ(a float64) Matrix {
	m := Identity(4, 4)
	m.Set(0, 0, math.Cos(a))
	m.Set(0, 1, -math.Sin(a))
	m.Set(1, 0, math.Sin(a))
	m.Set(1, 1, math.Cos(a))
	return m
}

func (m Matrix) RotationZ(a float64) Matrix {
	t := RotationZ(a)
	return m.Multiply(t)
}

func Shear(xy, xz, yx, yz, zx, zy float64) Matrix {
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
	t := Shear(xy, xz, yx, yz, zx, zy)
	return m.Multiply(t)
}
