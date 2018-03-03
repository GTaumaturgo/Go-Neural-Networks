package main

import "math"

func sigmoid(x float64) float64 {
	return math.Exp(x) / (math.Exp(x) + 1)
}

func dsigmoid(y float64) float64 {
	return y * (1 - y)
}
