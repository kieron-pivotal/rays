package matrix

import (
	"log"

	"github.com/kieron-pivotal/rays/tuple"
)

type Matrix struct {
	rows   int
	cols   int
	values []float64
}

func New(rows, cols int, vals ...float64) *Matrix {
	valsCopy := make([]float64, rows*cols)
	copy(valsCopy, vals)
	m := Matrix{
		rows:   rows,
		cols:   cols,
		values: valsCopy,
	}
	return &m
}

func (m *Matrix) Rows() int {
	return m.rows
}

func (m *Matrix) Cols() int {
	return m.cols
}

func (m *Matrix) Val(r, c int) float64 {
	if r > m.rows-1 || c > m.cols-1 {
		log.Panicf("row %d, col %d not contained in a %dx%d matrix", r, c, m.rows, m.cols)
	}
	idx := r*m.cols + c
	if idx > len(m.values)-1 {
		return 0.0
	}
	return m.values[idx]
}

func (m *Matrix) Set(r, c int, v float64) {
	m.values[r*m.cols+c] = v
}

func (m *Matrix) Equals(n *Matrix) bool {
	if m.rows != n.rows || m.cols != n.cols {
		return false
	}
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			if !floatEquals(m.Val(r, c), n.Val(r, c)) {
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

func (m *Matrix) Multiply(n *Matrix) *Matrix {
	out := New(m.rows, n.cols)

	for r := 0; r < m.rows; r++ {
		for c := 0; c < n.cols; c++ {
			var v float64
			for i := 0; i < m.cols; i++ {
				v += out.Val(r, c) + m.Val(r, i)*n.Val(i, c)
			}
			out.Set(r, c, v)
		}
	}
	return out
}

func (m *Matrix) TupleMultiply(t tuple.Tuple) tuple.Tuple {
	tm := New(4, 1, t.X, t.Y, t.Z, t.W)
	p := m.Multiply(tm)
	return tuple.Tuple{
		X: p.values[0],
		Y: p.values[1],
		Z: p.values[2],
		W: p.values[3],
	}
}

func Identity(r, c int) *Matrix {
	m := New(r, c)

	for i := 0; i < r; i++ {
		m.Set(i, i, 1)
	}
	return m
}

func (m *Matrix) Transpose() *Matrix {
	t := New(m.cols, m.rows)
	for c := 0; c < m.rows; c++ {
		for r := 0; r < m.cols; r++ {
			t.Set(r, c, m.Val(c, r))
		}
	}
	return t
}
