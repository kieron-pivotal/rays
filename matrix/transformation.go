package matrix

import "math"

func Translation(x, y, z float64) Matrix {
	m := Identity(4, 4)
	m.Set(0, 3, x)
	m.Set(1, 3, y)
	m.Set(2, 3, z)
	return m
}

func Scaling(x, y, z float64) Matrix {
	m := Identity(4, 4)
	m.Set(0, 0, x)
	m.Set(1, 1, y)
	m.Set(2, 2, z)
	return m
}

func RotationX(a float64) Matrix {
	m := Identity(4, 4)
	m.Set(1, 1, math.Cos(a))
	m.Set(1, 2, -math.Sin(a))
	m.Set(2, 1, math.Sin(a))
	m.Set(2, 2, math.Cos(a))
	return m
}

func RotationY(a float64) Matrix {
	m := Identity(4, 4)
	m.Set(0, 0, math.Cos(a))
	m.Set(0, 2, math.Sin(a))
	m.Set(2, 0, -math.Sin(a))
	m.Set(2, 2, math.Cos(a))
	return m
}

func RotationZ(a float64) Matrix {
	m := Identity(4, 4)
	m.Set(0, 0, math.Cos(a))
	m.Set(0, 1, -math.Sin(a))
	m.Set(1, 0, math.Sin(a))
	m.Set(1, 1, math.Cos(a))
	return m
}
