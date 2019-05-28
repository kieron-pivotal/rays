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

		DescribeTable("t values", func(rayOrigin, rayDirection tuple.Tuple, t1, t2 float64) {
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

	})

})
