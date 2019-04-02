package geometry

import "math"

const EPSILON = 0.00001

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: w}
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

func (t Tuple) Add(s Tuple) Tuple {
	return Tuple{
		X: t.X + s.X,
		Y: t.Y + s.Y,
		Z: t.Z + s.Z,
		W: t.W + s.W,
	}
}

func (t Tuple) Subtract(s Tuple) Tuple {
	return Tuple{
		X: t.X - s.X,
		Y: t.Y - s.Y,
		Z: t.Z - s.Z,
		W: t.W - s.W,
	}
}

func (t Tuple) Negate() Tuple {
	return Tuple{
		X: -t.X,
		Y: -t.Y,
		Z: -t.Z,
		W: -t.W,
	}
}

func (t Tuple) Multiply(c float64) Tuple {
	return Tuple{
		X: t.X * c,
		Y: t.Y * c,
		Z: t.Z * c,
		W: t.W * c,
	}
}

func (t Tuple) Divide(c float64) Tuple {
	return Tuple{
		X: t.X / c,
		Y: t.Y / c,
		Z: t.Z / c,
		W: t.W / c,
	}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
}

func floatEquals(f, g float64) bool {
	diff := f - g
	if diff < 0 {
		diff = -diff
	}
	return diff < EPSILON
}
