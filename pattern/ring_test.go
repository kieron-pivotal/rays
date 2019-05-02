package pattern_test

import (
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/pattern"
	"github.com/kieron-pivotal/rays/tuple"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ring", func() {
	It("should extend in both x and z", func() {
		p := pattern.Ring{
			A: white,
			B: black,
		}
		Expect(p.PatternAt(tuple.Point(0, 0, 0))).To(color.Equal(white))
		Expect(p.PatternAt(tuple.Point(1, 0, 0))).To(color.Equal(black))
		Expect(p.PatternAt(tuple.Point(0, 0, 1))).To(color.Equal(black))
		Expect(p.PatternAt(tuple.Point(0.708, 0, 0.708))).To(color.Equal(black))
	})
})
