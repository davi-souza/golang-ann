package network

import (
	"log"
	"math/rand"
)

type Network struct {
	InLayer *Layer
	OutLayer *Layer
	HiddenLayers []*Layer
}

// create new network
func CreateNetwork(numberOfNeuronInInput, numberOfHiddenLayer, numberOfNeuronPerHiddenLayer, numberOfNeuronInOutput int) *Network {
	createdNetwork := &Network{}

	createdNetwork.InLayer = CreateLayer(numberOfNeuronInInput)
	createdNetwork.OutLayer = CreateLayer(numberOfNeuronInOutput)

	for ; numberOfHiddenLayer > 0 ; numberOfHiddenLayer-- {
		createdLayer := CreateLayer(numberOfNeuronPerHiddenLayer)
		createdNetwork.HiddenLayers = append(createdNetwork.HiddenLayers, createdLayer)
	}

	createdNetwork.CreateAllSynapses()
	return createdNetwork
}

// create all the synapses between layers
func (network *Network) CreateAllSynapses() {
	network.CreateInputLayerSynapses()
	network.CreateHiddenLayersSynapses()
	network.CreateOutputLayerSynapses()
}
func (network *Network) CreateInputLayerSynapses() {
	network.InLayer.ConnectToLayer(network.HiddenLayers[0])
}
func (network *Network) CreateHiddenLayersSynapses() {
	for index := 0 ; index < len(network.HiddenLayers)-1 ; index++ {
		network.HiddenLayers[index].ConnectToLayer(network.HiddenLayers[index+1])
	}
}
func (network *Network) CreateOutputLayerSynapses() {
	network.HiddenLayers[len(network.HiddenLayers)-1].ConnectToLayer(network.OutLayer)
}

// randomize synapses weights
func (network *Network) RandomizeSynapsesWeights() {
	network.RandomizeInputLayerSynapses()
	network.RandomizeHiddenLayersSynapses()
	network.RandomizeOutputLayerSynapses()
}
func (network *Network) RandomizeInputLayerSynapses() {
	for _, neuron := range network.InLayer.Neurons {
		for _, synapse := range neuron.OutSynapses {
			synapse.Weight = rand.Float64()
		}
	}
}
func (network *Network) RandomizeHiddenLayersSynapses() {
	for _, layer := range network.HiddenLayers {
		for _, neuron := range layer.Neurons {
			for _, synapse := range neuron.InSynapses {
				synapse.Weight = rand.Float64()
			}
			for _, synapse := range neuron.OutSynapses {
				synapse.Weight = rand.Float64()
			}
		}
	}
}
func (network *Network) RandomizeOutputLayerSynapses() {
	for _, neuron := range network.OutLayer.Neurons {
		for _, synapse := range neuron.InSynapses {
			synapse.Weight = rand.Float64()
		}
	}
}

// setting the input
func (network *Network) SetInput(input []float64) {
	if len(input) != len(network.InLayer.Neurons) {
		log.Fatal("The input length must be equal to number of neurons in the Input Layer")
	}
	for index, number := range input {
		network.InLayer.Neurons[index].Activation = number
	}
}

// do all the calculus of the output
func (network *Network) NetInOutCalc() {
	network.InLayer.LayerOutCalc()
	for _, layer := range network.HiddenLayers {
		layer.LayerInOutCalc()
	}
	network.OutLayer.LayerInCalc()
}

// return the output slice
func (network *Network) Output() ([]float64) {
	var output []float64
	for _, neuron := range network.OutLayer.Neurons {
		output = append(output,neuron.Activation)
	}
	return output
}