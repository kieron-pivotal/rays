package shape_test

import (
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sphere", func() {

	It("intersects with a ray at two points", func() {
		ray := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))

		s := shape.NewSphere()

		xs := s.Intersect(ray)
		Expect(xs.Count()).To(Equal(2))
		Expect(xs.Get(0).T).To(Equal(4.0))
		Expect(xs.Get(1).T).To(Equal(6.0))
	})

	It("intersects at a tangent", func() {
		ray := ray.New(tuple.Point(0, 1, -5), tuple.Vector(0, 0, 1))

		s := shape.NewSphere()

		xs := s.Intersect(ray)
		Expect(xs.Count()).To(Equal(2))
		Expect(xs.Get(0).T).To(Equal(5.0))
		Expect(xs.Get(1).T).To(Equal(5.0))
	})

	It("misses", func() {
		ray := ray.New(tuple.Point(0, 2, -5), tuple.Vector(0, 0, 1))

		s := shape.NewSphere()

		xs := s.Intersect(ray)
		Expect(xs.Count()).To(Equal(0))
	})

	It("intersects with a ray originating inside the sphere", func() {
		ray := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))

		s := shape.NewSphere()

		xs := s.Intersect(ray)
		Expect(xs.Count()).To(Equal(2))
		Expect(xs.Get(0).T).To(Equal(-1.0))
		Expect(xs.Get(1).T).To(Equal(1.0))
	})

	It("intersects with a ray in front of the sphere", func() {
		ray := ray.New(tuple.Point(0, 0, 5), tuple.Vector(0, 0, 1))

		s := shape.NewSphere()

		xs := s.Intersect(ray)
		Expect(xs.Count()).To(Equal(2))
		Expect(xs.Get(0).T).To(Equal(-6.0))
		Expect(xs.Get(1).T).To(Equal(-4.0))
	})

})
