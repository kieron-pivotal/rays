package pattern_test

import (
	"github.com/kieron-pivotal/rays/pattern"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stripe", func() {
	var (
		p pattern.Stripe
	)

	BeforeEach(func() {
		p = pattern.Stripe{
			A: black,
			B: white,
		}
	})

	Context("stripes", func() {
		It("is constant in y", func() {
			Expect(p.PatternAt(tuple.Point(0, 0, 0))).To(Equal(black))
			Expect(p.PatternAt(tuple.Point(0, 1, 0))).To(Equal(black))
			Expect(p.PatternAt(tuple.Point(0, 2, 0))).To(Equal(black))
		})

		It("is constant in z", func() {
			Expect(p.PatternAt(tuple.Point(0, 0, 0))).To(Equal(black))
			Expect(p.PatternAt(tuple.Point(0, 0, 1))).To(Equal(black))
			Expect(p.PatternAt(tuple.Point(0, 0, 2))).To(Equal(black))
		})

		It("alternates in x", func() {
			Expect(p.PatternAt(tuple.Point(0, 0, 0))).To(Equal(black))
			Expect(p.PatternAt(tuple.Point(0.9, 0, 0))).To(Equal(black))
			Expect(p.PatternAt(tuple.Point(1, 0, 0))).To(Equal(white))
			Expect(p.PatternAt(tuple.Point(-0.1, 0, 0))).To(Equal(white))
			Expect(p.PatternAt(tuple.Point(-1, 0, 0))).To(Equal(white))
			Expect(p.PatternAt(tuple.Point(-1.1, 0, 0))).To(Equal(black))
		})
	})
})
