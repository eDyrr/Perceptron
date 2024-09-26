package main

type NeuronUnit struct {
	Weights []float64
	Bias    float64
	Lrate   float64
	Value   float64
	Delta   float64
}
