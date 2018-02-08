package main

import "math"

func sigmoid(x float64) float64 {
	return math.Exp(x) / (math.Exp(x) + 1)
}
