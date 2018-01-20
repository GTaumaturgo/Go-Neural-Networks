package main

import (
	"fmt"
)

func main() {
	p := newPerceptron()
	po := newPoint()
	inp := []float64{po.x, po.y}
	guess := p.guess(inp)
	fmt.Println(po)
	fmt.Println(guess)
}
