package camera_test

import (
	"math"

	"github.com/kieron-pivotal/rays/camera"
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/tuple"
	"github.com/kieron-pivotal/rays/world"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Camera", func() {
	It("can be constructed", func() {
		hsize := 160
		vsize := 120
		fieldOfView := math.Pi / 2
		c := camera.New(hsize, vsize, fieldOfView)
		Expect(c.HSize).To(BeNumerically("~", hsize))
		Expect(c.VSize).To(BeNumerically("~", vsize))
		Expect(c.FieldOfView).To(BeNumerically("~", fieldOfView))
		Expect(c.Transform).To(matrix.Equal(matrix.Identity(4, 4)))
	})

	It("knows the pixel size - landscape", func() {
		c := camera.New(200, 125, math.Pi/2)
		Expect(c.PixelSize).To(BeNumerically("~", 0.01))
	})

	It("knows the pixel size - portrait", func() {
		c := camera.New(125, 200, math.Pi/2)
		Expect(c.PixelSize).To(BeNumerically("~", 0.01))
	})

	Context("constructing a ray from a pixel", func() {
		It("works for ray through centre of canvax", func() {
			c := camera.New(201, 101, math.Pi/2)
			r := c.RayForPixel(100, 50)
			Expect(r.Origin).To(tuple.Equal(tuple.Point(0, 0, 0)))
			Expect(r.Direction).To(tuple.Equal(tuple.Vector(0, 0, -1)))
		})

		It("works for ray through corner of canvax", func() {
			c := camera.New(201, 101, math.Pi/2)
			r := c.RayForPixel(0, 0)
			Expect(r.Origin).To(tuple.Equal(tuple.Point(0, 0, 0)))
			Expect(r.Direction).To(tuple.Equal(tuple.Vector(0.66519, 0.33259, -0.66851)))
		})

		It("works when the camera is transformed", func() {
			c := camera.New(201, 101, math.Pi/2)
			c.Transform = matrix.Identity(4, 4).Translate(0, -2, 5).RotateY(math.Pi / 4)
			r := c.RayForPixel(100, 50)
			Expect(r.Origin).To(tuple.Equal(tuple.Point(0, 2, -5)))
			Expect(r.Direction).To(tuple.Equal(tuple.Vector(math.Sqrt(2)/2, 0, -math.Sqrt(2)/2)))
		})
	})

	Context("rendering a world", func() {
		It("can color a pixel", func() {
			w := world.Default()
			c := camera.New(11, 11, math.Pi/2)
			from := tuple.Point(0, 0, -5)
			to := tuple.Point(0, 0, 0)
			up := tuple.Vector(0, 1, 0)
			c.Transform = matrix.ViewTransformation(from, to, up)
			image := c.Render(w)
			Expect(image.Pixel(5, 5)).To(color.Equal(color.New(0.38066, 0.47583, 0.2855)))
		})
	})
})
