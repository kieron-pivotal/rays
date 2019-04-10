package matrix_test

import (
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Matrix", func() {

	Context("construction", func() {

		It("can create a 4x4 matrix with contents", func() {
			m := matrix.New(4, 4, 1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5)

			Expect(m.Rows()).To(Equal(4))
			Expect(m.Cols()).To(Equal(4))

			Expect(m.Get(0, 0)).To(BeNumerically("~", 1))
			Expect(m.Get(0, 3)).To(BeNumerically("~", 4))
			Expect(m.Get(1, 0)).To(BeNumerically("~", 5.5))
			Expect(m.Get(1, 2)).To(BeNumerically("~", 7.5))
			Expect(m.Get(2, 2)).To(BeNumerically("~", 11))
			Expect(m.Get(3, 0)).To(BeNumerically("~", 13.5))
			Expect(m.Get(3, 2)).To(BeNumerically("~", 15.5))
		})

		It("panics when value accessed outside of range", func() {
			m := matrix.New(4, 4, 1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5)
			Expect(func() { m.Get(4, 4) }).To(Panic())
		})

		It("treats unassigned vals as zeroes", func() {
			m := matrix.New(4, 4)
			Expect(m.Get(1, 1)).To(BeNumerically("~", 0))
		})

		It("has it's own copy of the matrix values", func() {
			vals := []float64{1.0, 2.0}
			m := matrix.New(4, 4, vals...)
			Expect(m.Get(0, 1)).To(BeNumerically("~", 2))
			vals[1] = 3.0
			Expect(m.Get(0, 1)).To(BeNumerically("~", 2))
		})

		It("can produce different sized matrices", func() {
			m := matrix.New(2, 2, -3, 5, 1, -2)
			Expect(m.Get(0, 0)).To(BeNumerically("~", -3))
			Expect(m.Get(0, 1)).To(BeNumerically("~", 5))
			Expect(m.Get(1, 0)).To(BeNumerically("~", 1))
			Expect(m.Get(1, 1)).To(BeNumerically("~", -2))
		})

		It("can do a 3x3 too", func() {
			m := matrix.New(3, 3, -3, 5, 0, 1, -2, 7, 0, 1, 1)
			Expect(m.Get(0, 0)).To(BeNumerically("~", -3))
			Expect(m.Get(1, 1)).To(BeNumerically("~", -2))
			Expect(m.Get(2, 2)).To(BeNumerically("~", 1))
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
			Expect(m.TupleMultiply(t)).To(tuple.Equal(tuple.New(18, 24, 33, 1)))
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

	Context("determinants", func() {
		It("can calculate a 2x2 determinant", func() {
			m := matrix.New(2, 2, 1, 5, -3, 2)
			Expect(m.Determinant()).To(BeNumerically("~", 17))
		})

		It("can calc the determinant of a 3x3 matrix", func() {
			m := matrix.New(3, 3,
				1, 2, 6,
				-5, 8, -4,
				2, 6, 4,
			)
			Expect(m.Cofactor(0, 0)).To(BeNumerically("~", 56))
			Expect(m.Cofactor(0, 1)).To(BeNumerically("~", 12))
			Expect(m.Cofactor(0, 2)).To(BeNumerically("~", -46))
			Expect(m.Determinant()).To(BeNumerically("~", -196))
		})

		It("can calc the determinant of a 4x4 matrix", func() {
			m := matrix.New(4, 4,
				-2, -8, 3, 5,
				-3, 1, 7, 3,
				1, 2, -9, 6,
				-6, 7, 7, -9,
			)
			Expect(m.Cofactor(0, 0)).To(BeNumerically("~", 690))
			Expect(m.Cofactor(0, 1)).To(BeNumerically("~", 447))
			Expect(m.Cofactor(0, 2)).To(BeNumerically("~", 210))
			Expect(m.Cofactor(0, 3)).To(BeNumerically("~", 51))
			Expect(m.Determinant()).To(BeNumerically("~", -4071))
		})
	})

	Context("submatrices", func() {
		It("gives a 2x2 submatrix of a 3x3 matrix", func() {
			m3 := matrix.New(3, 3, 1, 5, 0, -3, 2, 7, 0, 6, -3)
			m2 := matrix.New(2, 2, -3, 2, 0, 6)
			Expect(m3.Submatrix(0, 2)).To(matrix.Equal(m2))
		})

		It("gives a 3x3 submatrix of a 4x4 matrix", func() {
			m4 := matrix.New(4, 4,
				-6, 1, 1, 6,
				-8, 5, 8, 6,
				-1, 0, 8, 2,
				-7, 1, -1, 1,
			)
			m3 := matrix.New(3, 3,
				-6, 1, 6,
				-8, 8, 6,
				-7, -1, 1,
			)
			Expect(m4.Submatrix(2, 1)).To(matrix.Equal(m3))
		})
	})

	Context("minors", func() {
		It("calcs the minor of a 3x3 matrix", func() {
			m3 := matrix.New(3, 3,
				3, 5, 0,
				2, -1, -7,
				6, -1, 5,
			)
			Expect(m3.Minor(1, 0)).To(BeNumerically("~", 25))
		})
	})

	Context("cofactors", func() {
		It("calcs the cofactor of a 3x3 matrix", func() {
			m := matrix.New(3, 3,
				3, 5, 0,
				2, -1, -7,
				6, -1, 5,
			)

			Expect(m.Minor(0, 0)).To(BeNumerically("~", -12))
			Expect(m.Cofactor(0, 0)).To(BeNumerically("~", -12))
			Expect(m.Minor(1, 0)).To(BeNumerically("~", 25))
			Expect(m.Cofactor(1, 0)).To(BeNumerically("~", -25))
		})
	})

	Context("invertibility", func() {
		It("says an invertible matrix is invertible", func() {
			m := matrix.New(4, 4,
				6, 4, 4, 4,
				5, 5, 7, 6,
				4, -9, 3, -7,
				9, 1, 7, -6,
			)
			Expect(m.Determinant()).To(BeNumerically("~", -2120))
			Expect(m.IsInvertible()).To(BeTrue())
		})

		It("says a non-invertible matrix is not invertible", func() {
			m := matrix.New(4, 4,
				-4, 2, -2, -3,
				9, 6, 2, 6,
				0, -5, 1, -5,
				0, 0, 0, 0,
			)
			Expect(m.Determinant()).To(BeNumerically("~", 0))
			Expect(m.IsInvertible()).To(BeFalse())
		})
	})

	Context("inverses", func() {
		It("can invert a 4x4 matrix", func() {
			m := matrix.New(4, 4,
				-5, 2, 6, -8,
				1, -5, 1, 8,
				7, 7, -6, -7,
				1, -3, 7, 4,
			)
			n := m.Inverse()
			Expect(m.Determinant()).To(BeNumerically("~", 532))
			Expect(m.Cofactor(2, 3)).To(BeNumerically("~", -160))
			Expect(n.Get(3, 2)).To(BeNumerically("~", -160.0/532.0))
			Expect(m.Cofactor(3, 2)).To(BeNumerically("~", 105))
			Expect(n.Get(2, 3)).To(BeNumerically("~", 105.0/532.0))

			expectedInverse := matrix.New(4, 4,
				0.21805, 0.45113, 0.24060, -0.04511,
				-0.80827, -1.45677, -0.44361, 0.52068,
				-0.07895, -0.22368, -0.05263, 0.19737,
				-0.52256, -0.81391, -0.30075, 0.30639,
			)
			Expect(n).To(matrix.Equal(expectedInverse))
		})

		It("can reverse a multiplication with a inverse", func() {
			a := matrix.New(4, 4,
				3, -9, 7, 3,
				3, -8, 2, -9,
				-4, 4, 4, 1,
				-6, 5, -1, 1,
			)
			b := matrix.New(4, 4,
				8, 2, 2, 2,
				3, -1, 7, 0,
				7, 0, 5, 4,
				6, -2, 0, 5,
			)
			c := a.Multiply(b)

			Expect(c.Multiply(b.Inverse())).To(matrix.Equal(a))
		})
	})

	DescribeTable("more inverses", func(m, inv *matrix.Matrix) {
		Expect(m.Inverse()).To(matrix.Equal(inv))
	},
		Entry("1",
			matrix.New(4, 4,
				8, -5, 9, 2,
				7, 5, 6, 1,
				-6, 0, 9, 6,
				-3, 0, -9, -4,
			),
			matrix.New(4, 4,
				-0.15385, -0.15385, -0.28205, -0.53846,
				-0.07692, 0.12308, 0.02564, 0.03077,
				0.35897, 0.35897, 0.43590, 0.92308,
				-0.69231, -0.69231, -0.76923, -1.92308,
			),
		),
		Entry("2",
			matrix.New(4, 4,
				9, 3, 0, 9,
				-5, -2, -6, -3,
				-4, 9, 6, 4,
				-7, 6, 6, 2,
			),
			matrix.New(4, 4,
				-0.04074, -0.07778, 0.14444, -0.22222,
				-0.07778, 0.03333, 0.36667, -0.33333,
				-0.02901, -0.14630, -0.10926, 0.12963,
				0.17778, 0.06667, -0.26667, 0.33333,
			),
		),
	)

})

var _ = Describe("play", func() {
	var (
		m = matrix.New(4, 4,
			8, -5, 9, 2,
			7, 5, 6, 1,
			-6, 0, 9, 6,
			-3, 0, -9, -4,
		)
	)

	It("inverse of identity is identity", func() {
		id := matrix.Identity(4, 4)
		Expect(id.Inverse()).To(matrix.Equal(id))
	})

	It("a matrix times its inverse is identity", func() {
		Expect(m.Multiply(m.Inverse())).To(matrix.Equal(matrix.Identity(4, 4)))
	})

	It("order makes no difference doing inverse and transpose together", func() {
		Expect(m.Transpose().Inverse()).To(matrix.Equal(m.Inverse().Transpose()))
	})

	It("can multiply a single tuple element by changing a 1 in an identity matrix", func() {
		id := matrix.Identity(4, 4)
		id.Set(2, 2, 10)
		t := tuple.New(1, 2, 3, 4)
		Expect(id.TupleMultiply(t)).To(tuple.Equal(tuple.New(1, 2, 30, 4)))
	})
})
