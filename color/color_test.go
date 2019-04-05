package color_test

import (
	"github.com/kieron-pivotal/rays/color"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Color", func() {

	It("can create a color", func() {
		c := color.New(-0.5, 0.4, 1.7)
		Expect(c.Red()).To(BeNumerically("~", -0.5))
		Expect(c.Green()).To(BeNumerically("~", 0.4))
		Expect(c.Blue()).To(BeNumerically("~", 1.7))
	})

	It("can add colors", func() {
		c1 := color.New(0.9, 0.6, 0.75)
		c2 := color.New(0.7, 0.1, 0.25)

		Expect(c1.Add(c2)).To(color.Equal(color.New(1.6, 0.7, 1.0)))
	})

	It("can subtract colors", func() {
		c1 := color.New(0.9, 0.6, 0.75)
		c2 := color.New(0.7, 0.1, 0.25)

		Expect(c1.Subtract(c2)).To(color.Equal(color.New(0.2, 0.5, 0.5)))
	})

	It("can multiply a color by a scalar", func() {
		c := color.New(0.2, 0.3, 0.4)
		Expect(c.Multiply(2)).To(color.Equal(color.New(0.4, 0.6, 0.8)))
	})

	It("can multiply a color by a color", func() {
		c1 := color.New(1, 0.2, 0.4)
		c2 := color.New(0.9, 1, 0.1)
		Expect(c1.ColorMultiply(c2)).To(color.Equal(color.New(0.9, 0.2, 0.04)))
	})

})
