package main

type neuralNetwork struct {
	inputNodes  int
	hiddenNodes int
	outputNodes int
}

func newNeuralNetwork(numI, NumH, numO int) neuralNetwork {
	return neuralNetwork{
		inputNodes:  numI,
		hiddenNodes: NumH,
		outputNodes: numO,
	}
}
