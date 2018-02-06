package main

import (
	"fmt"
)

func main() {
	p := newPerceptron(3)
	po := []point{}
	for index := 0; index < 10; index++ {
		po = append(po, newPoint())
		fmt.Printf("%.1f %.1f %d\n", po[index].x, po[index].y, po[index].label)

	}
	for iter := 0; iter < 5; iter++ {
		c := 0
		for index := 0; index < 10; index++ {
			if p.guess([]float64{po[index].x, po[index].y}) == po[index].label {
				c++
			}
		}

		for index := 0; index < 10; index++ {
			inp := []float64{po[index].x, po[index].y, po[index].bias}
			p.train(inp, po[index].label)
		}
		fmt.Println(c)
	}
}
