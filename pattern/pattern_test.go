package pattern_test

import (
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/pattern"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	black = color.New(0, 0, 0)
	white = color.New(1, 1, 1)
)

var _ = Describe("Pattern", func() {

	var (
		p pattern.Stripe
	)

	BeforeEach(func() {
		p = pattern.NewStripe(black, white)
	})

	Context("stripes", func() {
		It("can create a stripe", func() {
			Expect(p.A).To(Equal(black))
			Expect(p.B).To(Equal(white))
		})

		It("is constant in y", func() {
			Expect(p.StripeAt(tuple.Point(0, 0, 0))).To(Equal(black))
			Expect(p.StripeAt(tuple.Point(0, 1, 0))).To(Equal(black))
			Expect(p.StripeAt(tuple.Point(0, 2, 0))).To(Equal(black))
		})

		It("is constant in z", func() {
			Expect(p.StripeAt(tuple.Point(0, 0, 0))).To(Equal(black))
			Expect(p.StripeAt(tuple.Point(0, 0, 1))).To(Equal(black))
			Expect(p.StripeAt(tuple.Point(0, 0, 2))).To(Equal(black))
		})

		It("alternates in x", func() {
			Expect(p.StripeAt(tuple.Point(0, 0, 0))).To(Equal(black))
			Expect(p.StripeAt(tuple.Point(0.9, 0, 0))).To(Equal(black))
			Expect(p.StripeAt(tuple.Point(1, 0, 0))).To(Equal(white))
			Expect(p.StripeAt(tuple.Point(-0.1, 0, 0))).To(Equal(white))
			Expect(p.StripeAt(tuple.Point(-1, 0, 0))).To(Equal(white))
			Expect(p.StripeAt(tuple.Point(-1.1, 0, 0))).To(Equal(black))
		})
	})

	Context("stripes with an object transformation", func() {
		It("moves with the object", func() {
			o := shape.NewSphere()
			o.SetTransform(matrix.Scaling(2, 2, 2))
			p := pattern.NewStripe(white, black)
			c := p.StripeAtObject(o.GetTransform(), tuple.Point(1.5, 0, 0))
			Expect(c).To(color.Equal(white))
		})
	})

	Context("stripes with a pattern transformation", func() {
		It("adjusts the pattern", func() {
			o := shape.NewSphere()
			p := pattern.NewStripe(white, black)
			p.SetTransform(matrix.Scaling(2, 2, 2))
			c := p.StripeAtObject(o.GetTransform(), tuple.Point(1.5, 0, 0))
			Expect(c).To(color.Equal(white))
		})
	})

	Context("both object and pattern transformation", func() {
		It("adjusts point and pattern", func() {
			o := shape.NewSphere()
			o.SetTransform(matrix.Scaling(2, 2, 2))
			p := pattern.NewStripe(white, black)
			p.SetTransform(matrix.Translation(0.5, 0, 0))
			c := p.StripeAtObject(o.GetTransform(), tuple.Point(2.5, 0, 0))
			Expect(c).To(color.Equal(white))
		})
	})
})
