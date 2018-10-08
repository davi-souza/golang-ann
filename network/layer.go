package network


type Layer struct {
	Neurons []*Neuron
}

// create new layer with the size == numberOfNeurons
func CreateLayer(size int) *Layer {
	// new layer
	createdLayer := &Layer{}
	// adding "size" Neurons to the new layer
	for ; size > 0 ; size-- {
		createdLayer.Neurons = append(createdLayer.Neurons, CreateNeuron())
	}
	return createdLayer
}

// connecting one layer to another
func (layer *Layer) ConnectToLayer(targetLayer *Layer) {
	// for each neuron of the layer
	for _, neuron := range layer.Neurons {
		// for each neuron of the target layer
		for _, targetNeuron := range targetLayer.Neurons {
			// create synapse
			neuron.CreateSynapseTo(targetNeuron,0.2)
		}
	}
}

// calculation funcs
func (layer *Layer) LayerInCalc() {
	for _, neuron := range layer.Neurons {
		neuron.InCalc()
	}
}
func (layer *Layer) LayerOutCalc() {
	for _, neuron := range layer.Neurons {
		neuron.OutCalc()
	}
}
func (layer *Layer) LayerInOutCalc() {
	for _, neuron := range layer.Neurons {
		neuron.InOutCalc()
	}
}