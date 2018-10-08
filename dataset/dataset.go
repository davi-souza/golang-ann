package dataset

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

const isFirstId = 0 //0 if it's not and 1 if it is
const numberOfPossibleOutput = 7

func Dataset() ([][]float64, [][]float64) {
	
	var dataset [][]float64
	var target [][]float64
	
	csvFile, _ := os.Open("dataset/raw-dataset.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		rawDatasetLength := len(line)
		// inputLength := rawDatasetLength - 1 - isFirstId
		outputIndex := rawDatasetLength - 1
		
		var newDatasetLine []float64

		for i := isFirstId ; i < outputIndex ; i++ {

			floatNumber, _ := strconv.ParseFloat(line[i], 64)
			
			newDatasetLine = append(newDatasetLine, floatNumber)

		}

		dataset = append(dataset, newDatasetLine)

		intNumber, _ := strconv.Atoi(line[outputIndex])
		var arrayToApend []float64
		for i:=0 ; i < numberOfPossibleOutput ; i++ {
			if i == intNumber-1 {
				arrayToApend = append(arrayToApend, 1.0)
			} else {
				arrayToApend = append(arrayToApend, 0.0)
			}
		}
		target = append(target,arrayToApend)
	}

	return dataset, target
}