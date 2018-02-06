package main

import (
	"fmt"
)

func main() {
	p := newPerceptron()
	po := []point{}
	for index := 0; index < 30000; index++ {
		po = append(po, newPoint())
	}
	for iter := 0; iter < 50; iter++ {
		c := 0
		for index := 0; index < 30000; index++ {
			if p.guess([]float64{po[index].x, po[index].y}) == po[index].label {
				c++
			}
		}

		for index := 0; index < 30000; index++ {
			inp := []float64{po[index].x, po[index].y, 1}
			p.train(inp, po[index].label)
		}
		fmt.Println(c)
	}
}
