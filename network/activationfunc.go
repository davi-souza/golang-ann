package network

import "math"

// One of the most used activation function
func Sigmoid(input float64) float64 {
	return (1/(1+math.Exp(-input)))
}

func ReLU(input float64) float64 {
	if input > 0 {
		return input
	} else {
		return 0.0
	}
}