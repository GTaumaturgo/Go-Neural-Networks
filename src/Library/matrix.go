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

}
