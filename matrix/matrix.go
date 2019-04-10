package matrix

import (
	"fmt"

	"github.com/kieron-pivotal/rays/tuple"
)

type Matrix struct {
	rows   int
	cols   int
	values []float64
}

func New(rows, cols int, vals ...float64) Matrix {
	valsCopy := make([]float64, rows*cols)
	copy(valsCopy, vals)
	m := Matrix{
		rows:   rows,
		cols:   cols,
		values: valsCopy,
	}
	return m
}

func (m Matrix) Rows() int {
	return m.rows
}

func (m Matrix) Cols() int {
	return m.cols
}

func (m Matrix) Get(r, c int) float64 {
	if r > m.rows-1 || c > m.cols-1 {
		panic(fmt.Sprintf("row %d, col %d not contained in a %dx%d matrix", r, c, m.rows, m.cols))
	}
	idx := r*m.cols + c
	if idx > len(m.values)-1 {
		return 0.0
	}
	return m.values[idx]
}

func (m Matrix) Set(r, c int, v float64) {
	m.values[r*m.cols+c] = v
}

func (m Matrix) Equals(n Matrix) bool {
	if m.rows != n.rows || m.cols != n.cols {
		return false
	}
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			if !floatEquals(m.Get(r, c), n.Get(r, c)) {
				return false
			}
		}
	}
	return true
}

func floatEquals(a, b float64) bool {
	const EPSILON = 0.00001
	diff := a - b
	if diff < 0 {
		diff *= -1
	}
	return diff < EPSILON
}

func (m Matrix) Multiply(n Matrix) Matrix {
	out := New(m.rows, n.cols)

	for r := 0; r < m.rows; r++ {
		for c := 0; c < n.cols; c++ {
			var v float64
			for i := 0; i < m.cols; i++ {
				v += out.Get(r, c) + m.Get(r, i)*n.Get(i, c)
			}
			out.Set(r, c, v)
		}
	}
	return out
}

func (m Matrix) TupleMultiply(t tuple.Tuple) tuple.Tuple {
	tm := New(4, 1, t.X, t.Y, t.Z, t.W)
	p := m.Multiply(tm)
	return tuple.Tuple{
		X: p.values[0],
		Y: p.values[1],
		Z: p.values[2],
		W: p.values[3],
	}
}

func Identity(r, c int) Matrix {
	m := New(r, c)

	for i := 0; i < r; i++ {
		m.Set(i, i, 1)
	}
	return m
}

func (m Matrix) Transpose() Matrix {
	t := New(m.cols, m.rows)
	for c := 0; c < m.rows; c++ {
		for r := 0; r < m.cols; r++ {
			t.Set(r, c, m.Get(c, r))
		}
	}
	return t
}

func (m Matrix) Determinant() float64 {
	if m.rows == 2 && m.cols == 2 {
		return m.Get(0, 0)*m.Get(1, 1) - m.Get(0, 1)*m.Get(1, 0)
	}
	det := float64(0)
	for i := 0; i < m.cols; i++ {
		det += m.Get(0, i) * m.Cofactor(0, i)
	}
	return det
}

func (m Matrix) Submatrix(r, c int) Matrix {
	o := New(m.rows-1, m.cols-1)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			if i == r || j == c {
				continue
			}
			row := i
			if row > r {
				row--
			}
			col := j
			if col > c {
				col--
			}
			o.Set(row, col, m.Get(i, j))
		}
	}
	return o
}

func (m Matrix) Minor(r, c int) float64 {
	return m.Submatrix(r, c).Determinant()
}

func (m Matrix) Cofactor(r, c int) float64 {
	min := m.Minor(r, c)
	if (r+c)%2 == 1 {
		min *= -1
	}
	return min
}

func (m Matrix) IsInvertible() bool {
	return m.Determinant() != 0.0
}

func (m Matrix) Inverse() Matrix {
	if !m.IsInvertible() {
		panic("matrix is not invertible")
	}
	det := m.Determinant()
	n := New(m.rows, m.cols)
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			n.Set(c, r, m.Cofactor(r, c)/det)
		}
	}
	return n
}
