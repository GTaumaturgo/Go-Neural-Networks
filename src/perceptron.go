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
