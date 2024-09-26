package main

type NeuralLayer struct {
	NeuronUnits []NeuronUnit
	Length      int
}

func PrepareLayer(n int, p int) (l NeuralLayer) {
	l = NeuralLayer{NeuronUnits: make([]NeuronUnit, n), Length: n}
}
