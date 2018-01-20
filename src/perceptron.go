package main

import (
	"math/rand"
	"time"
)

const inputs = 2

type perceptron struct {
	weights []float64
}

func newPerceptron() perceptron {
	res := perceptron{}
	res.weights = make([]float64, inputs)
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for index := 0; index < inputs; index++ {
		res.weights[index] = (r.Float64() * 2) - 1
	}
	return res
}

func (p perceptron) guess(inputs []float64) int {
	sum := 0.0
	for index := 0; index < len(inputs); index++ {
		sum += inputs[index] * p.weights[index]
	}
	return sign(sum)
}

// The activation function
func sign(f float64) int {
	if f >= 0 {
		return 1
	}
	return -1
}
