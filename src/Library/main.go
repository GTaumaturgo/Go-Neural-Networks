package main

import (
	"fmt"
)

func main() {
	m := newMatrix(2, 2)
	m.values[0][1] = 1
	m.values[1][0] = 1
	m2 := multMatrices(m, m)
	m2 = multMatrices(m2, m)

	fmt.Println(m2)
}
