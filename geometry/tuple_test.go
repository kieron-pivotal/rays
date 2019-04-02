package geometry_test

import (
	"github.com/kieron-pivotal/rays/geometry"
	"github.com/kieron-pivotal/rays/geometry/tuple_matcher"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tuple", func() {

	Context("point", func() {
		It("is a point when w is 1.0", func() {
			t := geometry.Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1.0}
			Expect(t.X).To(Equal(4.3))
			Expect(t.Y).To(Equal(-4.2))
			Expect(t.Z).To(Equal(3.1))
			Expect(t.W).To(Equal(1.0))
			Expect(t.IsPoint()).To(BeTrue())
			Expect(t.IsVector()).To(BeFalse())
		})

		It("can be created with Point constructor", func() {
			p := geometry.Point(4, -4, 3)
			Expect(p.IsPoint()).To(BeTrue())
		})
	})

	Context("vector", func() {
		It("is a vector when w is 0.0", func() {
			t := geometry.Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0.0}
			Expect(t.X).To(Equal(4.3))
			Expect(t.Y).To(Equal(-4.2))
			Expect(t.Z).To(Equal(3.1))
			Expect(t.W).To(Equal(0.0))
			Expect(t.IsPoint()).To(BeFalse())
			Expect(t.IsVector()).To(BeTrue())
		})

		It("can be created with Vector constructor", func() {
			p := geometry.Vector(4, -4, 3)
			Expect(p.IsVector()).To(BeTrue())
		})
	})

	Context("equality", func() {
		It("says a tuple equals itself", func() {
			p := geometry.Point(1, 2, 3)
			Expect(p.Equals(p)).To(BeTrue())
		})

		It("a point doesn't equal a vector", func() {
			p := geometry.Point(1, 2, 3)
			v := geometry.Vector(1, 2, 3)
			Expect(p.Equals(v)).To(BeFalse())
		})

		It("allows differences of < 0.00001", func() {
			p := geometry.Point(1, 2, 3)
			q := geometry.Point(1.000004, 1.999991, 2.999991)
			Expect(p.Equals(q)).To(BeTrue())
			Expect(q.Equals(p)).To(BeTrue())
		})

		It("can use the custom tuple matcher", func() {
			p := geometry.Point(1, 2, 3)
			q := geometry.Point(1.000004, 1.999991, 2.999991)
			Expect(p).To(tuple_matcher.Equal(q))
			r := geometry.Vector(1, 2, 3)
			Expect(p).ToNot(tuple_matcher.Equal(r))
		})
	})

})
