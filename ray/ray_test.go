package ray_test

import (
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ray", func() {
	It("can create and query a ray", func() {
		origin := tuple.Point(1, 2, 3)
		direction := tuple.Vector(4, 5, 6)

		r := ray.New(origin, direction)
		Expect(r.Origin).To(tuple.Equal(origin))
		Expect(r.Direction).To(tuple.Equal(direction))
	})

	It("can give position at time/distance t", func() {
		p := tuple.Point(2, 3, 4)
		d := tuple.Vector(1, 0, 0)
		r := ray.New(p, d)

		Expect(r.Position(0)).To(tuple.Equal(tuple.Point(2, 3, 4)))
		Expect(r.Position(1)).To(tuple.Equal(tuple.Point(3, 3, 4)))
		Expect(r.Position(-1)).To(tuple.Equal(tuple.Point(1, 3, 4)))
		Expect(r.Position(2.5)).To(tuple.Equal(tuple.Point(4.5, 3, 4)))
	})

})

var _ = Describe("tranformations", func() {
	It("can be translated", func() {
		r := ray.New(tuple.Point(1, 2, 3), tuple.Vector(0, 1, 0))
		m := matrix.Translation(3, 4, 5)
		r2 := r.Transform(m)
		Expect(r2.Origin).To(tuple.Equal(tuple.Point(4, 6, 8)))
		Expect(r2.Direction).To(tuple.Equal(tuple.Vector(0, 1, 0)))
	})

	It("can be scaled", func() {
		r := ray.New(tuple.Point(1, 2, 3), tuple.Vector(0, 1, 0))
		m := matrix.Scaling(2, 3, 4)
		r2 := r.Transform(m)
		Expect(r2.Origin).To(tuple.Equal(tuple.Point(2, 6, 12)))
		Expect(r2.Direction).To(tuple.Equal(tuple.Vector(0, 3, 0)))
	})
})
