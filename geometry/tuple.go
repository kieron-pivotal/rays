package geometry

const EPSILON = 0.00001

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func Point(x, y, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 1.0}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: 0.0}
}

func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

func (t Tuple) Equals(s Tuple) bool {
	return floatEquals(t.X, s.X) &&
		floatEquals(t.X, s.X) &&
		floatEquals(t.Z, s.Z) &&
		floatEquals(t.W, s.W)
}

func floatEquals(f, g float64) bool {
	diff := f - g
	if diff < 0 {
		diff = -diff
	}
	return diff < EPSILON
}
