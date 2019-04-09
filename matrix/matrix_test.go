package matrix_test

import (
	"github.com/kieron-pivotal/rays/matrix"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Matrix", func() {

	It("can create a 4x4 matrix with contents", func() {
		m := matrix.New(4, 4, 1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5)

		Expect(m.Rows()).To(Equal(4))
		Expect(m.Cols()).To(Equal(4))

		Expect(m.Val(0, 0)).To(BeNumerically("~", 1))
		Expect(m.Val(0, 3)).To(BeNumerically("~", 4))
		Expect(m.Val(1, 0)).To(BeNumerically("~", 5.5))
		Expect(m.Val(1, 2)).To(BeNumerically("~", 7.5))
		Expect(m.Val(2, 2)).To(BeNumerically("~", 11))
		Expect(m.Val(3, 0)).To(BeNumerically("~", 13.5))
		Expect(m.Val(3, 2)).To(BeNumerically("~", 15.5))
	})

	It("panics when value accessed outside of range", func() {
		m := matrix.New(4, 4, 1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5)
		Expect(func() { m.Val(4, 4) }).To(Panic())
	})

	It("treats unassigned vals as zeroes", func() {
		m := matrix.New(4, 4)
		Expect(m.Val(1, 1)).To(BeNumerically("~", 0))
	})

	It("has it's own copy of the matrix values", func() {
		vals := []float64{1.0, 2.0}
		m := matrix.New(4, 4, vals...)
		Expect(m.Val(0, 1)).To(BeNumerically("~", 2))
		vals[1] = 3.0
		Expect(m.Val(0, 1)).To(BeNumerically("~", 2))
	})

	It("can produce different sized matrices", func() {
		m := matrix.New(2, 2, -3, 5, 1, -2)
		Expect(m.Val(0, 0)).To(BeNumerically("~", -3))
		Expect(m.Val(0, 1)).To(BeNumerically("~", 5))
		Expect(m.Val(1, 0)).To(BeNumerically("~", 1))
		Expect(m.Val(1, 1)).To(BeNumerically("~", -2))
	})

	It("can do a 3x3 too", func() {

	})
})
