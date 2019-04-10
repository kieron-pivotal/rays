package tuple_test

import (
	"math"

	"github.com/kieron-pivotal/rays/tuple"
	"github.com/kieron-pivotal/rays/tuple/tuple_matcher"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tuple", func() {

	Context("point", func() {
		It("is a point when w is 1.0", func() {
			t := tuple.New(4.3, -4.2, 3.1, 1.0)
			Expect(t.X).To(Equal(4.3))
			Expect(t.Y).To(Equal(-4.2))
			Expect(t.Z).To(Equal(3.1))
			Expect(t.W).To(Equal(1.0))
			Expect(t.IsPoint()).To(BeTrue())
			Expect(t.IsVector()).To(BeFalse())
		})

		It("can be created with Point constructor", func() {
			p := tuple.Point(4, -4, 3)
			Expect(p.IsPoint()).To(BeTrue())
		})
	})

	Context("vector", func() {
		It("is a vector when w is 0.0", func() {
			t := tuple.New(4.3, -4.2, 3.1, 0.0)
			Expect(t.X).To(Equal(4.3))
			Expect(t.Y).To(Equal(-4.2))
			Expect(t.Z).To(Equal(3.1))
			Expect(t.W).To(Equal(0.0))
			Expect(t.IsPoint()).To(BeFalse())
			Expect(t.IsVector()).To(BeTrue())
		})

		It("can be created with Vector constructor", func() {
			p := tuple.Vector(4, -4, 3)
			Expect(p.IsVector()).To(BeTrue())
		})
	})

	Context("equality", func() {
		It("says a tuple equals itself", func() {
			p := tuple.Point(1, 2, 3)
			Expect(p.Equals(p)).To(BeTrue())
		})

		It("a point doesn't equal a vector", func() {
			p := tuple.Point(1, 2, 3)
			v := tuple.Vector(1, 2, 3)
			Expect(p.Equals(v)).To(BeFalse())
		})

		It("allows differences of < 0.00001", func() {
			p := tuple.Point(1, 2, 3)
			q := tuple.Point(1.000004, 1.999991, 2.999991)
			Expect(p.Equals(q)).To(BeTrue())
			Expect(q.Equals(p)).To(BeTrue())
		})

		It("can use the custom tuple matcher", func() {
			p := tuple.Point(1, 2, 3)
			q := tuple.Point(1.000004, 1.999991, 2.999991)
			Expect(p).To(tuple_matcher.Equal(q))
			r := tuple.Vector(1, 2, 3)
			Expect(p).ToNot(tuple_matcher.Equal(r))
		})

		DescribeTable("checking component differences", func(t1, t2 tuple.Tuple) {
			Expect(t1.Equals(t2)).To(BeFalse())
		},

			Entry("x", tuple.Vector(1, 2, 3), tuple.Vector(0, 2, 3)),
			Entry("y", tuple.Vector(1, 2, 3), tuple.Vector(1, 1, 3)),
			Entry("z", tuple.Vector(1, 2, 3), tuple.Vector(1, 2, 4)),
			Entry("w", tuple.Point(1, 2, 3), tuple.Vector(1, 2, 3)),
		)
	})

	Context("arithmetic", func() {
		It("can add points and vectors", func() {
			p := tuple.Point(3, -2, 5)
			v := tuple.Vector(-2, 3, 1)
			Expect(p.Add(v)).To(tuple_matcher.Equal(tuple.Point(1, 1, 6)))
		})

		It("can add two vectors", func() {
			v := tuple.Vector(3, -2, 1)
			w := tuple.Vector(1, -5, -4)
			Expect(v.Add(w)).To(tuple_matcher.Equal(tuple.Vector(4, -7, -3)))
		})

		It("produces nonsense when adding two points", func() {
			p := tuple.Point(1, 2, 3)
			q := tuple.Point(4, 5, 6)
			r := p.Add(q)
			Expect(r.IsPoint()).To(BeFalse())
			Expect(r.IsVector()).To(BeFalse())
		})

		It("can subtract two points to give a vector", func() {
			p := tuple.Point(1, 2, 3)
			q := tuple.Point(4, 5, 6)
			v := p.Subtract(q)
			Expect(v).To(tuple_matcher.Equal(tuple.Vector(-3, -3, -3)))
		})

		It("can subtract two vectors to give another vector", func() {
			v := tuple.Vector(3, 1, 2)
			w := tuple.Vector(1, 2, 1)
			u := v.Subtract(w)
			Expect(u).To(tuple_matcher.Equal(tuple.Vector(2, -1, 1)))
		})

		It("can subtract a vector from a point to give a point", func() {
			p := tuple.Point(3, 2, 1)
			v := tuple.Vector(5, 6, 7)
			Expect(p.Subtract(v)).To(tuple_matcher.Equal(tuple.Point(-2, -4, -6)))
		})

		It("can subtract a vector from a vector to give a vector", func() {
			v := tuple.Vector(3, 2, 1)
			w := tuple.Vector(5, 6, 7)
			Expect(v.Subtract(w)).To(tuple_matcher.Equal(tuple.Vector(-2, -4, -6)))
		})
	})

	Context("negating", func() {
		It("can negate an arbitrary tuple", func() {
			t := tuple.New(1, -2, 3, -4)
			Expect(t.Negate()).To(tuple_matcher.Equal(tuple.New(-1, 2, -3, 4)))
		})

		It("can negate a vector", func() {
			v := tuple.Vector(1, -2, 3)
			Expect(v.Negate()).To(tuple_matcher.Equal(tuple.Vector(-1, 2, -3)))
		})
	})

	Context("scalar multiplication", func() {
		It("can multiple a tuple by a scalar", func() {
			v := tuple.New(1, -2, 3, -4)
			Expect(v.Multiply(3.5)).To(tuple_matcher.Equal(tuple.New(3.5, -7, 10.5, -14)))
		})

		It("can multiple a tuple by a fractional scalar", func() {
			v := tuple.New(1, -2, 3, -4)
			Expect(v.Multiply(0.5)).To(tuple_matcher.Equal(tuple.New(0.5, -1, 1.5, -2)))
		})
	})

	Context("scalar division", func() {
		It("can divide a tuple by a scalar", func() {
			t := tuple.New(1, -2, 3, -4)
			Expect(t.Divide(2)).To(tuple_matcher.Equal(tuple.New(0.5, -1, 1.5, -2)))
		})
	})

	Context("magnitude", func() {
		It("can give the magnitude of a vector", func() {
			x := tuple.Vector(1, 0, 0)
			Expect(x.Magnitude()).To(BeNumerically("~", 1))
			y := tuple.Vector(0, 1, 0)
			Expect(y.Magnitude()).To(BeNumerically("~", 1))
			z := tuple.Vector(0, 0, 1)
			Expect(z.Magnitude()).To(BeNumerically("~", 1))

			v := tuple.Vector(1, 2, 3)
			Expect(v.Magnitude()).To(BeNumerically("~", math.Sqrt(14)))

			w := tuple.Vector(-1, -2, -3)
			Expect(w.Magnitude()).To(BeNumerically("~", math.Sqrt(14)))
		})
	})

	Context("normalization", func() {
		It("normalizes (4, 0, 0) to (1, 0, 0)", func() {
			v := tuple.Vector(4, 0, 0)
			Expect(v.Normalize()).To(Equal(tuple.Vector(1, 0, 0)))
		})

		It("normalizes (1, 2, 3) correctly", func() {
			v := tuple.Vector(1, 2, 3)
			r14 := math.Sqrt(14)
			n := v.Normalize()
			Expect(n).To(Equal(tuple.Vector(1/r14, 2/r14, 3/r14)))
			Expect(n.Magnitude()).To(BeNumerically("~", 1))
		})
	})

	Context("dot product", func() {
		It("can calculate it correctly", func() {
			v := tuple.Vector(1, 2, 3)
			w := tuple.Vector(2, 3, 4)
			Expect(v.Dot(w)).To(BeNumerically("~", 20))
		})
	})

	Context("cross product", func() {
		It("can calculate it correctly", func() {
			v := tuple.Vector(1, 2, 3)
			w := tuple.Vector(2, 3, 4)
			Expect(v.Cross(w)).To(tuple_matcher.Equal(tuple.Vector(-1, 2, -1)))
			Expect(w.Cross(v)).To(tuple_matcher.Equal(tuple.Vector(1, -2, 1)))
		})
	})
})
