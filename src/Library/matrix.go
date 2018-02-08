package main

import (
	"math/rand"
	"time"
)

/*Matrix is a type that implements a mathematical matrix*/
type Matrix struct {
	rows   int
	cols   int
	values [][]float64
}

func newMatrix(rows, cols int) Matrix {
	r := Matrix{
		rows: rows,
		cols: cols,
	}
	r.values = make([][]float64, rows)
	for i := 0; i < len(r.values); i++ {
		r.values[i] = make([]float64, cols)
	}
	return r
}

func (m Matrix) multEscalar(e int) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.values[i][j] *= float64(e)
		}
	}
}

func (m Matrix) addEscalar(e int) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.values[i][j] += float64(e)
		}
	}
}

// check if matrices have the same dimensions?
func addMatrices(m1, m2 Matrix) Matrix {
	r := newMatrix(m1.rows, m1.cols)
	for i := 0; i < r.rows; i++ {
		for j := 0; j < r.cols; j++ {
			r.values[i][j] = m1.values[i][j] + m2.values[i][j]
		}
	}
	return r
}

func subtractMatrices(m1, m2 Matrix) Matrix {
	r := newMatrix(m1.rows, m1.cols)
	for i := 0; i < r.rows; i++ {
		for j := 0; j < r.cols; j++ {
			r.values[i][j] = m1.values[i][j] - m2.values[i][j]
		}
	}
	return r
}

func multMatrices(m1, m2 Matrix) Matrix {
	result := newMatrix(m1.rows, m2.cols)

	for i := 0; i < m1.rows; i++ {
		for j := 0; j < m2.cols; j++ {
			sum := 0.0
			for k := 0; k < m1.cols; k++ {
				sum += m1.values[i][k] * m2.values[k][j]
			}
			result.values[i][j] = sum
		}
	}
	return result
}

func transpose(m Matrix) Matrix {
	r := newMatrix(m.cols, m.rows)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			r.values[j][i] = m.values[i][j]
		}
	}
	return r
}

func matrixFromArray(inp []float64) Matrix {
	m := newMatrix(len(inp), 1)
	for i := 0; i < m.rows; i++ {
		m.values[i][0] = inp[i]
	}
	return m
}
func arrayFromMatrix(m Matrix) []float64 {
	a := make([]float64, m.rows)
	for i := 0; i < m.rows; i++ {
		a[i] = m.values[i][0]
	}
	return a
}

func (m Matrix) randomize() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.values[i][j] = (r.Float64() * 2) - 1
		}
	}
}

func (m Matrix) mapf(fun func(float64) float64) {

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.values[i][j] = fun(m.values[i][j])
		}
	}
}
