package network

type Synapse struct {
	In float64
	Out float64
	Weight float64
}

// Calculate the signal of the synapse
func (s *Synapse) Signal(input float64) {
	s.In = input
	s.Out = s.In * s.Weight
}

// Create new synapse
func CreateNewSynapse(weight float64) *Synapse {
	return &Synapse{Weight: weight}
}

// Create new synapse and make its connections
func CreateSynapseFromTo(from,to *Neuron, weight float64) *Synapse {
	// Creating new synapse with weight
	newSynapse := CreateNewSynapse(weight)
	// Adding created synapse to the "from" Neuron
	from.OutSynapses = append(from.OutSynapses, newSynapse)
	// Adding created synapses to the "to" Neuron
	to.InSynapses = append(to.InSynapses, newSynapse)
	return newSynapse
}