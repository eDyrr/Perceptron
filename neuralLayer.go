package main

type NeuralLayer struct {
	NeuronUnits []NeuronUnit
	Length      int
}

func PrepareLayer(n int, p int) (l NeuralLayer) {
	l = NeuralLayer{NeuronUnits: make([]NeuronUnit, n), Length: n}

	for i := 0; i < n; i++ {
		RandomNeuronInit(&l.NeuronUnits[i], p)
	}

	log.WithFields(log.Fields{
		"level":               "info",
		"msg":                 "multilayer perceptron init completed",
		"neurons":             len(l.NeuronUnits),
		"lengthPreviousLayer": l.Length,
	}).Info("Complete NeuralLayer init.")

	return
}
