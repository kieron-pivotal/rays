package shape_test

import (
	"math"

	"github.com/kieron-pivotal/rays/material"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/shape/shapefakes"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shape", func() {
	var (
		s           *shape.Object
		localObject *shapefakes.FakeLocalObject
	)

	BeforeEach(func() {
		localObject = new(shapefakes.FakeLocalObject)
		s = shape.New(localObject)
	})

	Context("transformations", func() {
		It("has identity as the default transformation", func() {
			Expect(s.GetTransform()).To(matrix.Equal(matrix.Identity(4, 4)))
		})

		It("can be assigned a transformation", func() {
			t := matrix.Translation(1, 2, 3)
			s.SetTransform(t)
			Expect(s.GetTransform()).To(matrix.Equal(matrix.Translation(1, 2, 3)))
		})
	})

	Context("material", func() {
		It("has a default material", func() {
			m := s.Material()
			Expect(m).To(Equal(material.New()))
		})

		It("can be assigned a material", func() {
			m := material.New()
			m.Ambient = 1
			s.SetMaterial(m)
			Expect(s.Material().Ambient).To(BeNumerically("~", 1))
		})
	})

	Context("intersections", func() {
		It("calls the local intersect method with an inversely scaled ray", func() {
			r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
			r2 := ray.New(tuple.Point(0, 0, -2.5), tuple.Vector(0, 0, 0.5))
			s.SetTransform(matrix.Scaling(2, 2, 2))
			localObject.LocalIntersectReturns([]float64{3, 7})

			xs := s.Intersect(r)
			Expect(localObject.LocalIntersectCallCount()).To(Equal(1))
			Expect(localObject.LocalIntersectArgsForCall(0)).To(Equal(r2))

			Expect(xs.Count()).To(Equal(2))
			Expect(xs.Get(0).T).To(BeNumerically("~", 3))
			Expect(xs.Get(1).T).To(BeNumerically("~", 7))
		})

		It("calls the local intersect method with an inversely translated ray", func() {
			r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
			r2 := ray.New(tuple.Point(-5, 0, -5), tuple.Vector(0, 0, 1))
			s.SetTransform(matrix.Translation(5, 0, 0))
			localObject.LocalIntersectReturns([]float64{})

			xs := s.Intersect(r)
			Expect(localObject.LocalIntersectCallCount()).To(Equal(1))
			Expect(localObject.LocalIntersectArgsForCall(0)).To(Equal(r2))
			Expect(xs.Count()).To(Equal(0))
		})
	})

	Context("normals", func() {
		It("calls the local normal function on the local object", func() {
			s.SetTransform(matrix.Translation(0, 1, 0))
			s.NormalAt(tuple.Point(0, 1.70711, -0.70711))
			Expect(localObject.LocalNormalAtCallCount()).To(Equal(1))
			Expect(localObject.LocalNormalAtArgsForCall(0)).To(tuple.Equal(tuple.Point(0, 0.70711, -0.70711)))
		})

		r2 := math.Sqrt(2)

		DescribeTable("calculating the normal on a transformed unit sphere",
			func(point tuple.Tuple, transformation matrix.Matrix, normal tuple.Tuple) {
				s.SetTransform(transformation)
				localObject.LocalNormalAtStub = func(p tuple.Tuple) tuple.Tuple {
					return p
				}
				Expect(s.NormalAt(point)).To(tuple.Equal(normal))
				Expect(normal).To(tuple.Equal(normal.Normalize()))
			},

			Entry("translation",
				tuple.Point(0, 1.70711, -0.70711),
				matrix.Translation(0, 1, 0),
				tuple.Vector(0, 0.70711, -0.70711)),

			Entry("scale and rotation",
				tuple.Point(0, r2/2, -r2/2),
				matrix.Identity(4, 4).RotateZ(math.Pi/5).Scale(1, 0.5, 1),
				tuple.Vector(0, 0.97014, -0.24254)),
		)
	})
})
