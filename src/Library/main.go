package main

import "fmt"

func main() {
	nn := newNeuralNetwork(2, 2, 2)
	fmt.Println(nn)
	input := []float64{1, 0}
	target := []float64{1, 0}
	nn.train(input, target)

}
