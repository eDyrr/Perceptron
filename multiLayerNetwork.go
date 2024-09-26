package main

type MultiLayerNetwork struct {
	L_rate       float64
	NeuralLayers []NeuralLayer
	T_func       transferFunction
	T_func_d     transferFunction
}
