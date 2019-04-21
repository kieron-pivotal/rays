package shape_test

import (
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Intersection", func() {

	var (
		s  *shape.Sphere
		ix *shape.Intersections
	)

	BeforeEach(func() {
		s = shape.NewSphere()
		ix = shape.NewIntersections()
	})

	It("gets lower of 1 and 2", func() {
		ix.Add(1, s)
		ix.Add(2, s)
		i := ix.Hit()
		Expect(i.T).To(BeNumerically("~", 1))
	})

	It("gets higher of 1 and -1", func() {
		ix.Add(-1, s)
		ix.Add(1, s)
		i := ix.Hit()
		Expect(i.T).To(BeNumerically("~", 1))
	})

	It("returns nil for no hit", func() {
		ix.Add(-1, s)
		ix.Add(-2, s)
		i := ix.Hit()
		Expect(i).To(BeNil())
	})

	It("returns nil for no hit", func() {
		ix.Add(5, s)
		ix.Add(7, s)
		ix.Add(-3, s)
		ix.Add(2, s)
		i := ix.Hit()
		Expect(i.T).To(BeNumerically("~", 2))
	})

	Context("PrepareComputations", func() {
		It("can prepare common details", func() {
			o := shape.NewSphere()
			i := shape.Intersection{
				T:      4,
				Object: o,
			}
			r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
			comps := i.PrepareComputations(r)
			Expect(comps.T).To(BeNumerically("~", 4))
			Expect(comps.Object).To(Equal(o))
			Expect(comps.Point).To(tuple.Equal(tuple.Point(0, 0, -1)))
			Expect(comps.EyeV).To(tuple.Equal(tuple.Vector(0, 0, -1)))
			Expect(comps.NormalV).To(tuple.Equal(tuple.Vector(0, 0, -1)))
		})

		It("can identify an outside hit", func() {
			o := shape.NewSphere()
			i := shape.Intersection{
				T:      4,
				Object: o,
			}
			r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
			comps := i.PrepareComputations(r)
			Expect(comps.Inside).To(BeFalse())
		})

		It("can identify an inside hit", func() {
			o := shape.NewSphere()
			i := shape.Intersection{
				T:      1,
				Object: o,
			}
			r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
			comps := i.PrepareComputations(r)
			Expect(comps.Point).To(tuple.Equal(tuple.Point(0, 0, 1)))
			Expect(comps.EyeV).To(tuple.Equal(tuple.Vector(0, 0, -1)))
			Expect(comps.Inside).To(BeTrue())
			Expect(comps.NormalV).To(tuple.Equal(tuple.Vector(0, 0, -1)))
		})
	})
})
