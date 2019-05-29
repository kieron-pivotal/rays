package shape_test

import (
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cube", func() {

	Context("ray intersection", func() {

		DescribeTable("t values when ray intersects cube", func(rayOrigin, rayDirection tuple.Tuple, t1, t2 float64) {
			c := shape.Cube{}
			r := ray.New(rayOrigin, rayDirection)
			xs := c.LocalIntersect(r)
			Expect(len(xs)).To(Equal(2))
			Expect(xs[0]).To(Equal(t1))
			Expect(xs[1]).To(Equal(t2))
		},

			Entry("+x", tuple.Point(5, 0.5, 0), tuple.Vector(-1, 0, 0), 4.0, 6.0),
			Entry("-x", tuple.Point(-5, 0.5, 0), tuple.Vector(1, 0, 0), 4.0, 6.0),
			Entry("+y", tuple.Point(0.5, 5, 0), tuple.Vector(0, -1, 0), 4.0, 6.0),
			Entry("-y", tuple.Point(0.5, -5, 0), tuple.Vector(0, 1, 0), 4.0, 6.0),
			Entry("+z", tuple.Point(0.5, 0, 5), tuple.Vector(0, 0, -1), 4.0, 6.0),
			Entry("-z", tuple.Point(0.5, 0, -5), tuple.Vector(0, 0, 1), 4.0, 6.0),
			Entry("inside", tuple.Point(0, 0.5, 0), tuple.Vector(0, 0, 1), -1.0, 1.0),
		)

		DescribeTable("ray misses cube", func(rayOrigin, rayDirection tuple.Tuple) {
			c := shape.Cube{}
			r := ray.New(rayOrigin, rayDirection)
			xs := c.LocalIntersect(r)
			Expect(len(xs)).To(Equal(0))
		},

			Entry("1", tuple.Point(-2, 0, 0), tuple.Vector(0.2673, 0.5345, 0.8018)),
			Entry("2", tuple.Point(0, -2, 0), tuple.Vector(0.8018, 0.2673, 0.5345)),
			Entry("3", tuple.Point(0, 0, -2), tuple.Vector(0.5345, 0.8018, 0.2673)),
			Entry("4", tuple.Point(2, 0, 2), tuple.Vector(0, 0, -1)),
			Entry("5", tuple.Point(0, 2, 2), tuple.Vector(0, -1, 0)),
			Entry("6", tuple.Point(2, 2, 0), tuple.Vector(-1, 0, 0)),
		)

	})

	Context("normals", func() {
		DescribeTable("normals on cube surfaces", func(p, normal tuple.Tuple) {
			c := shape.Cube{}
			Expect(c.LocalNormalAt(p)).To(tuple.Equal(normal))
		},

			Entry("1", tuple.Point(1, 0.5, -0.8), tuple.Vector(1, 0, 0)),
			Entry("2", tuple.Point(-1, -0.2, 0.9), tuple.Vector(-1, 0, 0)),
			Entry("3", tuple.Point(-0.4, 1, -0.1), tuple.Vector(0, 1, 0)),
			Entry("4", tuple.Point(0.3, -1, -0.7), tuple.Vector(0, -1, 0)),
			Entry("5", tuple.Point(-0.6, 0.3, 1), tuple.Vector(0, 0, 1)),
			Entry("6", tuple.Point(0.4, 0.4, -1), tuple.Vector(0, 0, -1)),
			Entry("7", tuple.Point(1, 1, 1), tuple.Vector(1, 0, 0)),
			Entry("8", tuple.Point(-1, -1, -1), tuple.Vector(-1, 0, 0)),
		)
	})

})
