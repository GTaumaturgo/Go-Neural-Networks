package main

type neuralNetwork struct {
	inputNodes  int
	hiddenNodes int
	outputNodes int
	weightsIH   Matrix
	weightsHO   Matrix
	biasH       Matrix
	biasO       Matrix
	lr          float64
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
		lr:          0.08,
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

	outputs := multMatrices(nn.weightsHO, hid)
	outputs = addMatrices(outputs, nn.biasO)
	outputs.mapf(sigmoid)
	return arrayFromMatrix(outputs)
}

func (nn *neuralNetwork) train(inputs, targets []float64) {

	hid := multMatrices(nn.weightsIH, matrixFromArray(inputs))
	hid = addMatrices(hid, nn.biasH)
	hid.mapf(sigmoid)

	outputs := multMatrices(nn.weightsHO, hid)
	outputs = addMatrices(outputs, nn.biasO)
	outputs.mapf(sigmoid)
	outputErrors := subtractMatrices(matrixFromArray(targets), outputs)

	outputGradient := subtractMatrices(outputs, outputs) // zeroed matrix
	outputGradient = addMatrices(outputGradient, outputs)
	outputGradient.mapf(dsigmoid)
	outputGradient = multMatrices(outputGradient, outputErrors)
	outputGradient.multEscalar(nn.lr)

	hiddenT := transpose(hid)
	whoDeltas := multMatrices(outputGradient, hiddenT)

	nn.weightsHO = addMatrices(nn.weightsHO, whoDeltas)
	nn.biasO = addMatrices(nn.biasO, outputGradient)

	whoT := transpose(nn.weightsHO)
	hiddenErrors := multMatrices(whoT, outputErrors)

	hiddenGradient := subtractMatrices(hid, hid)
	hiddenGradient = addMatrices(hiddenGradient, hid)

	hiddenGradient.mapf(dsigmoid)
	hiddenGradient = multMatrices(hiddenGradient, hiddenErrors)
	hiddenGradient.multEscalar(nn.lr)

	inputsT := transpose(matrixFromArray(inputs))
	wihDeltas := multMatrices(hiddenGradient, inputsT)

	nn.weightsIH = addMatrices(nn.weightsIH, wihDeltas)
	nn.biasH = addMatrices(nn.biasH, hiddenGradient)

}
