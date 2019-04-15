package shape_test

import (
	"github.com/kieron-pivotal/rays/shape"
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
})
