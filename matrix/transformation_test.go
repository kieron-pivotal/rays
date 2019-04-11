package matrix_test

import (
	"math"

	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/tuple"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Translation", func() {
	var (
		p           = tuple.Point(-3, 4, 5)
		translation = matrix.Translation(5, -3, 2)
	)

	It("translates a point", func() {
		Expect(translation.TupleMultiply(p)).To(tuple.Equal(tuple.Point(2, 1, 7)))
	})

	It("sends it back with the inverse translation", func() {
		Expect(translation.Inverse().TupleMultiply(p)).To(tuple.Equal(tuple.Point(-8, 7, 3)))
	})

	It("does not affect vectors", func() {
		v := tuple.Vector(-3, 4, 5)
		Expect(translation.TupleMultiply(v)).To(tuple.Equal(v))
	})

})

var _ = Describe("scaling", func() {
	var (
		scaling = matrix.Scaling(2, 3, 4)
		point   = tuple.Point(-4, 6, 8)
		vector  = tuple.Vector(-4, 6, 8)
	)

	It("can scale a point", func() {
		Expect(scaling.TupleMultiply(point)).To(tuple.Equal(tuple.Point(-8, 18, 32)))
	})

	It("can scale a vector", func() {
		Expect(scaling.TupleMultiply(vector)).To(tuple.Equal(tuple.Vector(-8, 18, 32)))
	})

	It("scales the other way with an inverse", func() {
		Expect(scaling.Inverse().TupleMultiply(vector)).To(tuple.Equal(tuple.Vector(-2, 2, 2)))
	})
})

var _ = Describe("reflection", func() {
	It("can reflect by scaling by a negative value", func() {
		reflect := matrix.Scaling(-1, 1, 1)
		point := tuple.Point(2, 3, 4)
		Expect(reflect.TupleMultiply(point)).To(tuple.Equal(tuple.Point(-2, 3, 4)))
	})
})

var _ = Describe("rotation", func() {
	It("can rotate about x-axis", func() {
		eighthTurn := matrix.RotationX(math.Pi / 4)
		quarterTurn := matrix.RotationX(math.Pi / 2)
		point := tuple.Point(0, 1, 0)
		Expect(eighthTurn.TupleMultiply(point)).To(tuple.Equal(tuple.Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2)))
		Expect(quarterTurn.TupleMultiply(point)).To(tuple.Equal(tuple.Point(0, 0, 1)))
	})

	It("rotates the other way with an inverse", func() {
		eighthTurn := matrix.RotationX(math.Pi / 4)
		point := tuple.Point(0, 1, 0)
		Expect(eighthTurn.Inverse().TupleMultiply(point)).To(tuple.Equal(tuple.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)))
	})

	It("can rotate about the y-axis", func() {
		eighthTurn := matrix.RotationY(math.Pi / 4)
		quarterTurn := matrix.RotationY(math.Pi / 2)
		point := tuple.Point(0, 0, 1)
		Expect(eighthTurn.TupleMultiply(point)).To(tuple.Equal(tuple.Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)))
		Expect(quarterTurn.TupleMultiply(point)).To(tuple.Equal(tuple.Point(1, 0, 0)))
	})

	It("can rotate about the z-axis", func() {
		eighthTurn := matrix.RotationZ(math.Pi / 4)
		quarterTurn := matrix.RotationZ(math.Pi / 2)
		point := tuple.Point(0, 1, 0)
		Expect(eighthTurn.TupleMultiply(point)).To(tuple.Equal(tuple.Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)))
		Expect(quarterTurn.TupleMultiply(point)).To(tuple.Equal(tuple.Point(-1, 0, 0)))
	})
})

var _ = DescribeTable("shearing", func(shear matrix.Matrix, point, result tuple.Tuple) {
	Expect(shear.TupleMultiply(point)).To(tuple.Equal(result))
},
	Entry("x wrt y", matrix.Shear(1, 0, 0, 0, 0, 0), tuple.Point(2, 3, 4), tuple.Point(5, 3, 4)),
	Entry("x wrt z", matrix.Shear(0, 1, 0, 0, 0, 0), tuple.Point(2, 3, 4), tuple.Point(6, 3, 4)),
	Entry("y wrt x", matrix.Shear(0, 0, 1, 0, 0, 0), tuple.Point(2, 3, 4), tuple.Point(2, 5, 4)),
	Entry("y wrt z", matrix.Shear(0, 0, 0, 1, 0, 0), tuple.Point(2, 3, 4), tuple.Point(2, 7, 4)),
	Entry("z wrt x", matrix.Shear(0, 0, 0, 0, 1, 0), tuple.Point(2, 3, 4), tuple.Point(2, 3, 6)),
	Entry("z wrt y", matrix.Shear(0, 0, 0, 0, 0, 1), tuple.Point(2, 3, 4), tuple.Point(2, 3, 7)),
)
