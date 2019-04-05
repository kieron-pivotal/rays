package canvas_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kieron-pivotal/rays/canvas"
	"github.com/kieron-pivotal/rays/color"
)

var _ = Describe("Canvas", func() {
	It("can create blank canvas with correct size", func() {
		canvas := canvas.New(10, 20)
		Expect(canvas.Width).To(Equal(10))
		Expect(canvas.Height).To(Equal(20))

		for x := 0; x < 10; x++ {
			for y := 0; y < 20; y++ {
				Expect(canvas.Pixel(x, y)).To(color.Equal(color.New(0, 0, 0)))
			}
		}
	})

	It("can set the colour of a pixel", func() {
		canvas := canvas.New(10, 20)
		red := color.New(1, 0, 0)
		canvas.SetPixel(2, 3, red)
		Expect(canvas.Pixel(2, 3)).To(color.Equal(red))
	})

	Describe("rendering to PPM", func() {

		It("can render the header", func() {
			canvas := canvas.New(5, 3)
			ppm := canvas.ToPPM()
			Expect(ppm).To(HavePrefix("P3\n5 3\n255\n"))
		})

		It("can render the pixels", func() {
			canvas := canvas.New(5, 3)
			c1 := color.New(1.5, 0, 0)
			c2 := color.New(0, 0.5, 0)
			c3 := color.New(-0.5, 0, 1)
			canvas.SetPixel(0, 0, c1)
			canvas.SetPixel(2, 1, c2)
			canvas.SetPixel(4, 2, c3)

			ppm := canvas.ToPPM()

			expectedData := `
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 127 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`
			Expect(ppm).To(HaveSuffix(expectedData))
		})

		It("doesn't produce lines longer than 70", func() {
			canvas := canvas.New(10, 2)
			c := color.New(1, 0.8, 0.6)
			for i := 0; i < 10; i++ {
				for j := 0; j < 2; j++ {
					canvas.SetPixel(i, j, c)
				}
			}
			ppm := canvas.ToPPM()

			expectedData := `
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
`
			Expect(ppm).To(HaveSuffix(expectedData))
		})
	})
})
