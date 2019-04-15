package shape_test

import (
	"github.com/kieron-pivotal/rays/matrix"
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
		t := shape.NewSphere()

		xs := s.Intersect(ray)
		Expect(xs.Count()).To(Equal(2))
		Expect(xs.Get(0).T).To(Equal(-6.0))
		Expect(xs.Get(1).T).To(Equal(-4.0))
		Expect(xs.Get(0).Shape).To(Equal(s))
		Expect(xs.Get(0).Shape).NotTo(Equal(t))
	})

	Context("transformations", func() {
		var (
			s *shape.Sphere
		)

		BeforeEach(func() {
			s = shape.NewSphere()
		})

		It("has the identity as the default transformation", func() {
			Expect(s.GetTransform()).To(matrix.Equal(matrix.Identity(4, 4)))
		})

		It("can be set to another transformation", func() {
			t := matrix.Translation(2, 3, 4)
			s.SetTransform(t)
			Expect(s.GetTransform()).To(matrix.Equal(matrix.Translation(2, 3, 4)))
		})

		It("intersects a scaled sphere with a ray", func() {
			r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
			s.SetTransform(matrix.Scaling(2, 2, 2))
			xs := s.Intersect(r)

			Expect(xs.Count()).To(Equal(2))
			Expect(xs.Get(0).T).To(BeNumerically("~", 3))
			Expect(xs.Get(1).T).To(BeNumerically("~", 7))
		})

		It("intersects a translated sphere with a ray", func() {
			r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
			s.SetTransform(matrix.Translation(5, 0, 0))
			xs := s.Intersect(r)

			Expect(xs.Count()).To(Equal(0))
		})
	})

})
