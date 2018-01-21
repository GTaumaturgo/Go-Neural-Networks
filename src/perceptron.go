package main

import (
	"math/rand"
	"time"
)

const inputs = 2
const lr = 0.1

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

func (p perceptron) train(inputs []float64, target int) {
	g := p.guess(inputs)
	er := target - g
	for index := 0; index < len(p.weights); index++ {
		p.weights[index] += float64(er) * inputs[index]
	}
}

// The activation function
func sign(f float64) int {
	if f >= 0 {
		return 1
	}
	return -1
}
