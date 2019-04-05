package geometry

import (
	"log"
	"math"
)

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
		floatEquals(t.Y, s.Y) &&
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

func (t Tuple) BitwiseMultiply(s Tuple) Tuple {
	return Tuple{
		X: t.X * s.X,
		Y: t.Y * s.Y,
		Z: t.Z * s.Z,
		W: t.W * s.W,
	}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
}

func (t Tuple) Normalize() Tuple {
	mag := t.Magnitude()
	return Tuple{
		X: t.X / mag,
		Y: t.Y / mag,
		Z: t.Z / mag,
		W: t.W / mag,
	}
}

func (t Tuple) Dot(s Tuple) float64 {
	return t.X*s.X + t.Y*s.Y + t.Z*s.Z + t.W*s.W
}

func (t Tuple) Cross(s Tuple) Tuple {
	if !(t.IsVector() && s.IsVector()) {
		log.Fatal("Both operands must be vectors for a cross product")
	}

	return Tuple{
		X: t.Y*s.Z - t.Z*s.Y,
		Y: t.Z*s.X - t.X*s.Z,
		Z: t.X*s.Y - t.Y*s.X,
		W: 0.0,
	}
}

func floatEquals(f, g float64) bool {
	diff := f - g
	if diff < 0 {
		diff = -diff
	}
	return diff < EPSILON
}
