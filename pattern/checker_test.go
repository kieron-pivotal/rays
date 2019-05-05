package pattern_test

import (
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/pattern"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Checker", func() {

	var (
		p pattern.Checker
	)

	BeforeEach(func() {
		p = pattern.Checker{
			A: white,
			B: black,
		}
	})

	It("repeats in x", func() {
		Expect(p.PatternAt(tuple.Point(0, 0, 0))).To(color.Equal(white))
		Expect(p.PatternAt(tuple.Point(0.99, 0, 0))).To(color.Equal(white))
		Expect(p.PatternAt(tuple.Point(1.01, 0, 0))).To(color.Equal(black))
	})

	It("repeats in y", func() {
		Expect(p.PatternAt(tuple.Point(0, 0, 0))).To(color.Equal(white))
		Expect(p.PatternAt(tuple.Point(0, 0.99, 0))).To(color.Equal(white))
		Expect(p.PatternAt(tuple.Point(0, 1.01, 0))).To(color.Equal(black))
	})

	It("repeats in z", func() {
		Expect(p.PatternAt(tuple.Point(0, 0, 0))).To(color.Equal(white))
		Expect(p.PatternAt(tuple.Point(0, 0, 0.99))).To(color.Equal(white))
		Expect(p.PatternAt(tuple.Point(0, 0, 1.01))).To(color.Equal(black))
	})

})
