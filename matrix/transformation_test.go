package matrix_test

import (
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/tuple"

	. "github.com/onsi/ginkgo"
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
