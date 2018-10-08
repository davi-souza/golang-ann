package main

import (
	"fmt"
	network "./network"
	dataset "./dataset"
	learn "./learn"
)

func main()  {
	fmt.Println("\n----- Begin -----\n")
	dataset, target := dataset.Dataset()


	fmt.Println("\n----- Creating Network -----\n")
	Net := network.CreateNetwork(len(dataset[0]),2,9,len(target[0]))
	fmt.Println("\n----- Randomize Network Weights -----\n")
	Net.RandomizeSynapsesWeights()

	fmt.Println("\n----- Start Network Learn -----\n")
	for i := 5 ; i > 0 ; i-- {
		fmt.Printf("\n----- Learn step i: %d -----\n\n",i)
		for index, _ := range dataset {
			learn.Learn(Net, dataset[index], target[index], 0.5)
		}
	}
	fmt.Println("\n----- Finish Network Learn -----\n")

	for i:= 0 ; i < 10 ; i++ {
		Net.SetInput(dataset[i])
		Net.NetInOutCalc()
		output := Net.Output()
		fmt.Println("target",target[i])
		fmt.Println("output",output,"\n")
	}
}