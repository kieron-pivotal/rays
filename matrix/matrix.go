package matrix

import "log"

type Matrix struct {
	rows   int
	cols   int
	values []float64
}

func New(rows, cols int, vals ...float64) *Matrix {
	valsCopy := make([]float64, len(vals))
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
	idx := r*m.rows + c
	if idx > len(m.values)-1 {
		return 0.0
	}
	return m.values[r*m.rows+c]
}
