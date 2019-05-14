package shape_test

import (
	"math"

	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Unit Sphere", func() {
	var (
		s shape.Sphere
	)

	BeforeEach(func() {
		s = shape.Sphere{}
	})

	It("intersects with a ray at two points", func() {
		ray := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))

		xs := s.LocalIntersect(ray)
		Expect(xs).To(HaveLen(2))
		Expect(xs[0]).To(Equal(4.0))
		Expect(xs[1]).To(Equal(6.0))
	})

	It("intersects at a tangent", func() {
		ray := ray.New(tuple.Point(0, 1, -5), tuple.Vector(0, 0, 1))

		xs := s.LocalIntersect(ray)
		Expect(xs).To(HaveLen(2))
		Expect(xs[0]).To(Equal(5.0))
		Expect(xs[1]).To(Equal(5.0))
	})

	It("misses", func() {
		ray := ray.New(tuple.Point(0, 2, -5), tuple.Vector(0, 0, 1))

		xs := s.LocalIntersect(ray)
		Expect(xs).To(BeEmpty())
	})

	It("intersects with a ray originating inside the sphere", func() {
		ray := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))

		xs := s.LocalIntersect(ray)
		Expect(xs).To(HaveLen(2))
		Expect(xs[0]).To(Equal(-1.0))
		Expect(xs[1]).To(Equal(1.0))
	})

	It("intersects with a ray in front of the sphere", func() {
		ray := ray.New(tuple.Point(0, 0, 5), tuple.Vector(0, 0, 1))

		xs := s.LocalIntersect(ray)
		Expect(xs).To(HaveLen(2))
		Expect(xs[0]).To(Equal(-6.0))
		Expect(xs[1]).To(Equal(-4.0))
	})

	Context("normal", func() {
		r3 := math.Sqrt(3)

		DescribeTable("calculating the normal on the unit sphere centered on origin",
			func(point, normal tuple.Tuple) {
				Expect(s.LocalNormalAt(point)).To(tuple.Equal(normal))
				Expect(normal).To(tuple.Equal(normal.Normalize()))
			},

			Entry("1, 0, 0", tuple.Point(1, 0, 0), tuple.Vector(1, 0, 0)),
			Entry("0, 1, 0", tuple.Point(0, 1, 0), tuple.Vector(0, 1, 0)),
			Entry("0, 0, 1", tuple.Point(0, 0, 1), tuple.Vector(0, 0, 1)),
			Entry("r3/3, ...", tuple.Point(r3/3, r3/3, r3/3), tuple.Vector(r3/3, r3/3, r3/3)),
		)
	})

	It("is possible to easily create a glass sphere", func() {
		s := shape.NewGlassSphere()
		m := s.Material()
		Expect(m.Transparency).To(BeNumerically("~", 1.0))
		Expect(m.RefractiveIndex).To(BeNumerically("~", 1.5))
	})

})
