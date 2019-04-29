package shape_test

import (
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Plane", func() {

	var (
		p shape.Plane
	)

	BeforeEach(func() {
		p = shape.Plane{}
	})

	It("has a constant normal everywhere", func() {
		Expect(p.LocalNormalAt(tuple.Point(0, 0, 0))).To(tuple.Equal(tuple.Vector(0, 1, 0)))
		Expect(p.LocalNormalAt(tuple.Point(10, 0, -10))).To(tuple.Equal(tuple.Vector(0, 1, 0)))
		Expect(p.LocalNormalAt(tuple.Point(5, 0, -150))).To(tuple.Equal(tuple.Vector(0, 1, 0)))
	})

	Context("intersection", func() {
		When("the ray is parallel to the plane", func() {
			It("doesn't intersect", func() {
				r := ray.New(tuple.Point(0, 10, 0), tuple.Vector(0, 0, 1))
				ix := p.LocalIntersect(r)
				Expect(ix).To(HaveLen(0))
			})
		})

		When("the ray is coplanar with the plane", func() {
			It("doesn't intersect", func() {
				r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
				ix := p.LocalIntersect(r)
				Expect(ix).To(HaveLen(0))
			})
		})

		When("the ray intersects the plane", func() {
			It("intersects from above", func() {
				r := ray.New(tuple.Point(0, 1, 0), tuple.Vector(0, -1, 0))
				ix := p.LocalIntersect(r)
				Expect(ix).To(HaveLen(1))
				Expect(ix[0]).To(BeNumerically("~", 1))
			})

			It("intersects from below", func() {
				r := ray.New(tuple.Point(0, -1, 0), tuple.Vector(0, 1, 0))
				ix := p.LocalIntersect(r)
				Expect(ix).To(HaveLen(1))
				Expect(ix[0]).To(BeNumerically("~", 1))
			})
		})
	})
})
