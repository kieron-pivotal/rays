package pattern_test

import (
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/pattern"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gradient", func() {
	It("interpolates between two colors", func() {
		p := pattern.Gradient{
			A: white,
			B: black,
		}

		Expect(p.PatternAt(tuple.Point(0, 0, 0))).To(color.Equal(white))
		Expect(p.PatternAt(tuple.Point(0.25, 0, 0))).To(color.Equal(color.New(0.75, 0.75, 0.75)))
		Expect(p.PatternAt(tuple.Point(0.5, 0, 0))).To(color.Equal(color.New(0.5, 0.5, 0.5)))
		Expect(p.PatternAt(tuple.Point(0.75, 0, 0))).To(color.Equal(color.New(0.25, 0.25, 0.25)))
	})
})
