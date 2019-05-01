package play_test

import (
	"os"

	"github.com/kieron-pivotal/rays/canvas"
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/light"
	"github.com/kieron-pivotal/rays/material"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FirstRender", func() {

	It("can draw a circle from a sphere", func() {
		origin := tuple.Point(-5, 0, 0)
		wallX := float64(10)
		wallSize := float64(7)
		canvasPixels := 100
		pixelSize := wallSize / float64(canvasPixels)

		s := shape.NewSphere()
		s.SetTransform(matrix.Identity(4, 4).Scale(1, 0.5, 1).Shear(0, 0, 1, 0, 0, 1))
		canv := canvas.New(canvasPixels, canvasPixels)
		col := color.New(255, 0, 0)

		for r := 0; r < canvasPixels; r++ {
			for c := 0; c < canvasPixels; c++ {
				coord := tuple.Point(wallX, float64(r)*pixelSize-wallSize/2, float64(c)*pixelSize-wallSize/2)
				ray := ray.New(origin, coord.Subtract(origin))

				ix := s.Intersect(ray)
				if ix.Count() > 0 {
					canv.SetPixel(r, c, col)
				}
			}
		}
		file, err := os.Create("first_render.ppm")
		Expect(err).NotTo(HaveOccurred())

		file.WriteString(canv.ToPPM())
		file.Close()
	})

	It("can draw a lighted circle from a sphere", func() {
		origin := tuple.Point(-5, 0, 0)
		wallX := float64(10)
		wallSize := float64(7)
		canvasPixels := 600
		m := material.New()
		m.Color = color.New(1, 0.2, 1)

		pixelSize := wallSize / float64(canvasPixels)
		s := shape.NewSphere()
		s.SetMaterial(m)

		l := light.NewPoint(tuple.Point(-10, 10, -10), color.New(1, 1, 1))
		// s.SetTransform(matrix.Identity(4, 4).Scale(1, 0.7, 1).RotateX(math.Pi / 4))
		canv := canvas.New(canvasPixels, canvasPixels)

		for r := 0; r < canvasPixels; r++ {
			for c := 0; c < canvasPixels; c++ {
				coord := tuple.Point(wallX, wallSize/2-float64(r)*pixelSize, float64(c)*pixelSize-wallSize/2)
				ray := ray.New(origin, coord.Subtract(origin).Normalize())

				ix := s.Intersect(ray)
				if ix.Count() > 0 {
					hit := ix.Get(0)
					hitPoint := ray.Position(hit.T)
					normal := s.NormalAt(hitPoint)
					eye := ray.Direction.Multiply(-1)
					col := hit.Object.Material().Lighting(l, hit.Object.GetTransform(), hitPoint, eye, normal, false)
					canv.SetPixel(r, c, col)
				}
			}
		}
		file, err := os.Create("second_render.ppm")
		Expect(err).NotTo(HaveOccurred())

		file.WriteString(canv.ToPPM())
		file.Close()
	})

})
