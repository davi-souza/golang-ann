package learn

import (
	network "../network"
	"math"
)

func Cost(output, target []float64) float64 {
	var sum float64
	sum = 0.0
	for index, _ := range target {
		sum += math.Pow( (output[index] - target[index]), 2)
	}
	return sum/float64(len(target))
}

func SigmoidDerivative(x float64) float64 {
	return x * (1.0 - x)
}
func ReLUDerivative(x float64) float64 {
	if x < 0.0 {
		return 0.0
	} else {
		return 1.0
	}
}

func Learn(n *network.Network, input, target []float64, step float64) {
	n.SetInput(input)
	n.NetInOutCalc()

	deltas := make([][]float64, len(n.HiddenLayers) + 1)

	last := len(n.HiddenLayers)

	l := n.OutLayer

	deltas[last] = make([]float64, len(l.Neurons))

	for i, neuron := range l.Neurons {
		deltas[last][i] = SigmoidDerivative(neuron.Activation) * (target[i] - neuron.Activation)
	}

	for i := last - 1; i >= 0; i-- {
		l := n.HiddenLayers[i]
		deltas[i] = make([]float64, len(l.Neurons))
		for j, neuron := range l.Neurons {

			var sum float64 = 0
			for k, s := range neuron.OutSynapses {
				sum += s.Weight * deltas[i+1][k]
			}

			deltas[i][j] = SigmoidDerivative(neuron.Activation) * sum
		}
	}

	// last = len(n.HiddenLayers)

	for j, neuron := range n.OutLayer.Neurons {
		for _, s := range neuron.InSynapses {
			s.Weight += step * deltas[last][j] * s.In
		}
	}

	for i, l := range n.HiddenLayers {
		for j, neuron := range l.Neurons {
			for _, s := range neuron.InSynapses {
				s.Weight += step * deltas[i][j] * s.In
			}
		}
	}

}