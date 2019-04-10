package matrix

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
