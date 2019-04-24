package material_test

import (
	"math"

	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/light"
	"github.com/kieron-pivotal/rays/material"
	"github.com/kieron-pivotal/rays/tuple"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Material", func() {

	It("has default vals", func() {
		m := material.New()
		Expect(m.Color).To(color.Equal(color.New(1, 1, 1)))
		Expect(m.Ambient).To(BeNumerically("~", 0.1))
		Expect(m.Diffuse).To(BeNumerically("~", 0.9))
		Expect(m.Specular).To(BeNumerically("~", 0.9))
		Expect(m.Shininess).To(BeNumerically("~", 200))
	})

	Context("shading", func() {
		var (
			p  tuple.Tuple
			m  material.Material
			r2 = math.Sqrt(2)
		)

		BeforeEach(func() {
			p = tuple.Point(0, 0, 0)
			m = material.New()
		})

		DescribeTable("lighting",
			func(eye, normal tuple.Tuple, l light.Point, inShadow bool, expected color.Color) {
				Expect(m.Lighting(l, p, eye, normal, inShadow)).To(color.Equal(expected))
			},

			Entry("eye in front of light",
				tuple.Vector(0, 0, -1),
				tuple.Vector(0, 0, -1),
				light.NewPoint(tuple.Point(0, 0, -10), color.New(1, 1, 1)),
				false,
				color.New(1.9, 1.9, 1.9),
			),

			Entry("eye offset 45 degs",
				tuple.Vector(0, r2/2, -r2/2),
				tuple.Vector(0, 0, -1),
				light.NewPoint(tuple.Point(0, 0, -10), color.New(1, 1, 1)),
				false,
				color.New(1.0, 1.0, 1.0),
			),

			Entry("light offset 45 degs",
				tuple.Vector(0, 0, -1),
				tuple.Vector(0, 0, -1),
				light.NewPoint(tuple.Point(0, 10, -10), color.New(1, 1, 1)),
				false,
				color.New(0.7364, 0.7364, 0.7364),
			),

			Entry("light and eye offset opposite 45 degs",
				tuple.Vector(0, -r2/2, -r2/2),
				tuple.Vector(0, 0, -1),
				light.NewPoint(tuple.Point(0, 10, -10), color.New(1, 1, 1)),
				false,
				color.New(1.6364, 1.6364, 1.6364),
			),

			Entry("light behind the surface",
				tuple.Vector(0, 0, -1),
				tuple.Vector(0, 0, -1),
				light.NewPoint(tuple.Point(0, 0, 10), color.New(1, 1, 1)),
				false,
				color.New(0.1, 0.1, 0.1),
			),

			Entry("in shadow",
				tuple.Vector(0, 0, -1),
				tuple.Vector(0, 0, -1),
				light.NewPoint(tuple.Point(0, 0, -10), color.New(1, 1, 1)),
				true,
				color.New(0.1, 0.1, 0.1),
			),
		)

	})

})
