package geometry_test

import (
	"math"

	"github.com/kieron-pivotal/rays/geometry"
	"github.com/kieron-pivotal/rays/geometry/tuple_matcher"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tuple", func() {

	Context("point", func() {
		It("is a point when w is 1.0", func() {
			t := geometry.NewTuple(4.3, -4.2, 3.1, 1.0)
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
			t := geometry.NewTuple(4.3, -4.2, 3.1, 0.0)
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

		DescribeTable("checking component differences", func(t1, t2 geometry.Tuple) {
			Expect(t1.Equals(t2)).To(BeFalse())
		},

			Entry("x", geometry.Vector(1, 2, 3), geometry.Vector(0, 2, 3)),
			Entry("y", geometry.Vector(1, 2, 3), geometry.Vector(1, 1, 3)),
			Entry("z", geometry.Vector(1, 2, 3), geometry.Vector(1, 2, 4)),
			Entry("w", geometry.Point(1, 2, 3), geometry.Vector(1, 2, 3)),
		)
	})

	Context("arithmetic", func() {
		It("can add points and vectors", func() {
			p := geometry.Point(3, -2, 5)
			v := geometry.Vector(-2, 3, 1)
			Expect(p.Add(v)).To(tuple_matcher.Equal(geometry.Point(1, 1, 6)))
		})

		It("can add two vectors", func() {
			v := geometry.Vector(3, -2, 1)
			w := geometry.Vector(1, -5, -4)
			Expect(v.Add(w)).To(tuple_matcher.Equal(geometry.Vector(4, -7, -3)))
		})

		It("produces nonsense when adding two points", func() {
			p := geometry.Point(1, 2, 3)
			q := geometry.Point(4, 5, 6)
			r := p.Add(q)
			Expect(r.IsPoint()).To(BeFalse())
			Expect(r.IsVector()).To(BeFalse())
		})

		It("can subtract two points to give a vector", func() {
			p := geometry.Point(1, 2, 3)
			q := geometry.Point(4, 5, 6)
			v := p.Subtract(q)
			Expect(v).To(tuple_matcher.Equal(geometry.Vector(-3, -3, -3)))
		})

		It("can subtract two vectors to give another vector", func() {
			v := geometry.Vector(3, 1, 2)
			w := geometry.Vector(1, 2, 1)
			u := v.Subtract(w)
			Expect(u).To(tuple_matcher.Equal(geometry.Vector(2, -1, 1)))
		})

		It("can subtract a vector from a point to give a point", func() {
			p := geometry.Point(3, 2, 1)
			v := geometry.Vector(5, 6, 7)
			Expect(p.Subtract(v)).To(tuple_matcher.Equal(geometry.Point(-2, -4, -6)))
		})

		It("can subtract a vector from a vector to give a vector", func() {
			v := geometry.Vector(3, 2, 1)
			w := geometry.Vector(5, 6, 7)
			Expect(v.Subtract(w)).To(tuple_matcher.Equal(geometry.Vector(-2, -4, -6)))
		})
	})

	Context("negating", func() {
		It("can negate an arbitrary tuple", func() {
			t := geometry.NewTuple(1, -2, 3, -4)
			Expect(t.Negate()).To(tuple_matcher.Equal(geometry.NewTuple(-1, 2, -3, 4)))
		})

		It("can negate a vector", func() {
			v := geometry.Vector(1, -2, 3)
			Expect(v.Negate()).To(tuple_matcher.Equal(geometry.Vector(-1, 2, -3)))
		})
	})

	Context("scalar multiplication", func() {
		It("can multiple a tuple by a scalar", func() {
			v := geometry.NewTuple(1, -2, 3, -4)
			Expect(v.Multiply(3.5)).To(tuple_matcher.Equal(geometry.NewTuple(3.5, -7, 10.5, -14)))
		})

		It("can multiple a tuple by a fractional scalar", func() {
			v := geometry.NewTuple(1, -2, 3, -4)
			Expect(v.Multiply(0.5)).To(tuple_matcher.Equal(geometry.NewTuple(0.5, -1, 1.5, -2)))
		})
	})

	Context("scalar division", func() {
		It("can divide a tuple by a scalar", func() {
			t := geometry.NewTuple(1, -2, 3, -4)
			Expect(t.Divide(2)).To(tuple_matcher.Equal(geometry.NewTuple(0.5, -1, 1.5, -2)))
		})
	})

	Context("magnitude", func() {
		It("can give the magnitude of a vector", func() {
			x := geometry.Vector(1, 0, 0)
			Expect(x.Magnitude()).To(BeNumerically("~", 1))
			y := geometry.Vector(0, 1, 0)
			Expect(y.Magnitude()).To(BeNumerically("~", 1))
			z := geometry.Vector(0, 0, 1)
			Expect(z.Magnitude()).To(BeNumerically("~", 1))

			v := geometry.Vector(1, 2, 3)
			Expect(v.Magnitude()).To(BeNumerically("~", math.Sqrt(14)))

			w := geometry.Vector(-1, -2, -3)
			Expect(w.Magnitude()).To(BeNumerically("~", math.Sqrt(14)))
		})
	})

	Context("normalization", func() {
		It("normalizes (4, 0, 0) to (1, 0, 0)", func() {
			v := geometry.Vector(4, 0, 0)
			Expect(v.Normalize()).To(Equal(geometry.Vector(1, 0, 0)))
		})

		It("normalizes (1, 2, 3) correctly", func() {
			v := geometry.Vector(1, 2, 3)
			r14 := math.Sqrt(14)
			n := v.Normalize()
			Expect(n).To(Equal(geometry.Vector(1/r14, 2/r14, 3/r14)))
			Expect(n.Magnitude()).To(BeNumerically("~", 1))
		})
	})

	Context("dot product", func() {
		It("can calculate it correctly", func() {
			v := geometry.Vector(1, 2, 3)
			w := geometry.Vector(2, 3, 4)
			Expect(v.Dot(w)).To(BeNumerically("~", 20))
		})
	})

	Context("cross product", func() {
		It("can calculate it correctly", func() {
			v := geometry.Vector(1, 2, 3)
			w := geometry.Vector(2, 3, 4)
			Expect(v.Cross(w)).To(tuple_matcher.Equal(geometry.Vector(-1, 2, -1)))
			Expect(w.Cross(v)).To(tuple_matcher.Equal(geometry.Vector(1, -2, 1)))
		})
	})
})
