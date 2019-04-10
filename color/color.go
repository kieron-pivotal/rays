package color

import "github.com/kieron-pivotal/rays/tuple"

type Color struct {
	tuple tuple.Tuple
}

func New(r, g, b float64) Color {
	return Color{
		tuple: tuple.Vector(r, g, b),
	}
}

func (c Color) Red() float64 {
	return c.tuple.X
}

func (c Color) Green() float64 {
	return c.tuple.Y
}

func (c Color) Blue() float64 {
	return c.tuple.Z
}

func (c Color) Add(d Color) Color {
	return Color{
		tuple: c.tuple.Add(d.tuple),
	}
}

func (c Color) Subtract(d Color) Color {
	return Color{
		tuple: c.tuple.Subtract(d.tuple),
	}
}

func (c Color) Multiply(f float64) Color {
	return Color{
		tuple: c.tuple.Multiply(f),
	}
}

func (c Color) ColorMultiply(d Color) Color {
	return Color{
		tuple: c.tuple.BitwiseMultiply(d.tuple),
	}
}

func (c Color) Equals(d Color) bool {
	return c.tuple.Equals(d.tuple)
}
