package main

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
func (m Matrix) addMatrix(m1 Matrix) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.values[i][j] += m1.values[i][j]
		}
	}

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

func (m *Matrix) transpose() {
	r := newMatrix(m.cols, m.rows)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			r.values[j][i] = m.values[i][j]
		}
	}
	*m = r
}
