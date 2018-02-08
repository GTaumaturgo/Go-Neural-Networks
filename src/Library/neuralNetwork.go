package main

type neuralNetwork struct {
	inputNodes  int
	hiddenNodes int
	outputNodes int
	weightsIH   Matrix
	weightsHO   Matrix
	biasH       Matrix
	biasO       Matrix
}

func newNeuralNetwork(numI, numH, numO int) neuralNetwork {
	nn := neuralNetwork{
		inputNodes:  numI,
		hiddenNodes: numH,
		outputNodes: numO,
		weightsIH:   newMatrix(numH, numI),
		weightsHO:   newMatrix(numO, numH),
		biasH:       newMatrix(numH, 1),
		biasO:       newMatrix(numO, 1),
	}
	nn.weightsIH.randomize()
	nn.weightsHO.randomize()
	nn.biasH.randomize()
	nn.biasO.randomize()

	return nn
}

func (nn neuralNetwork) feedForward(input []float64) []float64 {
	hid := multMatrices(nn.weightsIH, matrixFromArray(input))
	hid = addMatrices(hid, nn.biasH)
	hid.mapf(sigmoid)

	out := multMatrices(nn.weightsHO, hid)
	out = addMatrices(out, nn.biasO)
	out.mapf(sigmoid)
	return arrayFromMatrix(out)
}

func (nn neuralNetwork) train(input, target []float64) {
	// out := matrixFromArray(nn.feedForward(input))
	// outErr := subtractMatrices(matrixFromArray(target), out)
	// wHOT := transpose(nn.weightsHO)

	// hidErr := multMatrices(wHOT, outErr)

}
