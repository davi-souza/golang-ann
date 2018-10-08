package network

type Neuron struct {
	Bias float64
	Activation float64
	InSynapses []*Synapse
	OutSynapses []*Synapse
}

// create new neuron
func CreateNeuron() *Neuron {
	return &Neuron{}
}

// create synapse from a neuron to another
func (neuron *Neuron) CreateSynapseTo(targetNeuron *Neuron, weight float64) {
	CreateSynapseFromTo(neuron, targetNeuron, weight)
}

// calculate the input value
func (neuron *Neuron) InCalc() {
	var sum = 0.0
	for _, inSynapse := range neuron.InSynapses {
		sum += inSynapse.Out
	}
	neuron.Activation = Sigmoid(sum)
}
// calculate the output value
func (neuron *Neuron) OutCalc() {
	for _, outSynapse := range neuron.OutSynapses {
		outSynapse.Signal(neuron.Activation)
	}
}
// do all the calculation
func (neuron *Neuron) InOutCalc() {
	neuron.InCalc()
	neuron.OutCalc()
}
