package camera

import (
	"math"

	"github.com/kieron-pivotal/rays/canvas"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/tuple"
	"github.com/kieron-pivotal/rays/world"
)

type Camera struct {
	HSize       int
	VSize       int
	FieldOfView float64
	Transform   matrix.Matrix
	HalfWidth   float64
	HalfHeight  float64
	PixelSize   float64
}

func New(hsize, vsize int, fieldOfView float64) Camera {
	c := Camera{
		HSize:       hsize,
		VSize:       vsize,
		FieldOfView: fieldOfView,
		Transform:   matrix.Identity(4, 4),
	}
	c.calcSizes()
	return c
}

func (c *Camera) calcSizes() {
	halfView := math.Tan(c.FieldOfView / 2)
	aspect := float64(c.HSize) / float64(c.VSize)

	if aspect >= 1 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspect
	} else {
		c.HalfWidth = halfView * aspect
		c.HalfHeight = halfView
	}
	c.PixelSize = (c.HalfWidth * 2.0) / float64(c.HSize)
}

func (c Camera) RayForPixel(px, py int) ray.Ray {
	xoffset := (float64(px) + 0.5) * c.PixelSize
	yoffset := (float64(py) + 0.5) * c.PixelSize
	worldX := c.HalfWidth - xoffset
	worldY := c.HalfHeight - yoffset
	pixel := c.Transform.Inverse().TupleMultiply(tuple.Point(worldX, worldY, -1))
	origin := c.Transform.Inverse().TupleMultiply(tuple.Point(0, 0, 0))
	direction := pixel.Subtract(origin).Normalize()
	return ray.New(origin, direction)
}

func (c Camera) Render(w *world.World) *canvas.Canvas {
	image := canvas.New(c.HSize, c.VSize)

	for py := 0; py < c.VSize; py++ {
		for px := 0; px < c.HSize; px++ {
			ray := c.RayForPixel(px, py)
			color := w.ColorAt(ray)
			image.SetPixel(px, py, color)
		}
	}
	return image
}
