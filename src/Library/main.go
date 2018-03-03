package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nn := newNeuralNetwork(2, 2, 1)

	samples := make([]Sample, 4)
	samples[0] = Sample{[]float64{1, 1}, []float64{0}}
	samples[1] = Sample{[]float64{0, 0}, []float64{0}}
	samples[2] = Sample{[]float64{1, 0}, []float64{1}}
	samples[3] = Sample{[]float64{0, 1}, []float64{1}}

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := 0; i < 50000; i++ {
		k := r.Intn(4)

		nn.train(samples[k].inputs, samples[k].targets)
	}

	fmt.Println(nn.feedForward(samples[0].inputs))
	fmt.Println(nn.feedForward(samples[1].inputs))
	fmt.Println(nn.feedForward(samples[2].inputs))
	fmt.Println(nn.feedForward(samples[3].inputs))
}
