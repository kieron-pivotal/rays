package play_test

import (
	"math"
	"os"

	"github.com/kieron-pivotal/rays/canvas"
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Clock", func() {

	It("can draw a clockface", func() {

		c := canvas.New(600, 600)

		for i := 0; i < 12; i++ {
			p := tuple.Point(1, 0, 0)
			t := matrix.Identity(4, 4).
				RotationZ(float64(i)*math.Pi/6).
				Scaling(250, 250, 250).
				Translation(300, 300, 0)

			p = t.TupleMultiply(p)
			c.SetPixel(int(math.Round(p.X)), int(math.Round(p.Y)), color.New(255, 255, 255))
		}

		file, err := os.Create("clock.ppm")
		Expect(err).NotTo(HaveOccurred())

		file.WriteString(c.ToPPM())
		file.Close()
	})

})
