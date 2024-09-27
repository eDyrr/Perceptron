package main

type MultiLayerNetwork struct {
	L_rate       float64
	NeuralLayers []NeuralLayer
	T_func       transferFunction
	T_func_d     transferFunction
}

func PrepareMLPNet(l []int, lr float64, tf transferFunction, trd transferFunction) (mlp MultiLayerNetwork) {
	mlp.L_rate = lr
	mlp.T_func = tf
	mlp.T_func_d = trd

	mlp.NeuralLayers = make([]NeuralLayer, len(l))
	for il, ql := range l {
		if il != 0 {
			mlp.NeuralLayers[il] = PrepareLayer(ql, l[il-1])
		} else {
			mlp.NeuralLayers[il] = PrepareLayer(ql, 0)
		}
	}

	log.WithFields(log.Fields{
		"level":        "info",
		"msg":          "multilayer pereceptron init completed",
		"layers":       len(mlp.NeuralLayers),
		"learningRate": mlp.L_rate,
	}).Info("Complete Multilayer Perceptron init.")

	return
}
