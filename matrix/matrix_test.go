package matrix_test

import (
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/tuple"
	"github.com/kieron-pivotal/rays/tuple/tuple_matcher"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Matrix", func() {

	Context("construction", func() {

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
			m := matrix.New(3, 3, -3, 5, 0, 1, -2, 7, 0, 1, 1)
			Expect(m.Val(0, 0)).To(BeNumerically("~", -3))
			Expect(m.Val(1, 1)).To(BeNumerically("~", -2))
			Expect(m.Val(2, 2)).To(BeNumerically("~", 1))
		})
	})

	Context("equality", func() {
		It("can see two identical matrices are the same", func() {
			m1 := matrix.New(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
			m2 := matrix.New(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
			Expect(m1.Equals(m2)).To(BeTrue())
		})
		It("can see two different matrices are not the same", func() {
			m1 := matrix.New(4, 4, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1)
			m2 := matrix.New(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
			Expect(m1.Equals(m2)).To(BeFalse())
		})
	})

	Context("matrix multiplication", func() {
		It("can multiply two 4x4 matrices", func() {
			m1 := matrix.New(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
			m2 := matrix.New(4, 4, -2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8)
			prod := matrix.New(4, 4, 20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42)
			Expect(m1.Multiply(m2)).To(matrix.Equal(prod))
		})

		It("is original matrix when multiplied by identity matrix", func() {
			m1 := matrix.New(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
			id := matrix.Identity(4, 4)
			Expect(m1.Multiply(id)).To(matrix.Equal(m1))

		})
	})

	Context("tuple multiplication", func() {
		It("can multiply a matrix by a tuple", func() {
			m := matrix.New(4, 4, 1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1)
			t := tuple.New(1, 2, 3, 1)
			Expect(m.TupleMultiply(t)).To(tuple_matcher.Equal(tuple.New(18, 24, 33, 1)))
		})
	})

	Context("transposition", func() {
		It("can transpose a matrix", func() {
			m := matrix.New(4, 4, 0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8)
			t := matrix.New(4, 4, 0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8)
			Expect(m.Transpose()).To(matrix.Equal(t))
		})

		It("transpose identity to identity", func() {
			id := matrix.Identity(4, 4)
			Expect(id.Transpose()).To(matrix.Equal(id))
		})
	})
})
