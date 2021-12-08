package d1p1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Run() {
	measurements, err := os.Open("d1p1/measurements.txt")
	if err != nil {
		log.Fatal("open: ", err)
	}
	defer measurements.Close()
	var previousMeasurement, currentMeasurement, depthIncreased int

	scanner := bufio.NewScanner(measurements)
	scanner.Scan()
	previousMeasurement = scanMeasurement(scanner)
	for scanner.Scan() {
		currentMeasurement = scanMeasurement(scanner)
		if previousMeasurement < currentMeasurement {
			depthIncreased++
		}
		previousMeasurement = currentMeasurement
	}
	log.Printf("Depth increased %d times\n", depthIncreased)
}

func scanMeasurement(scanner *bufio.Scanner) int {
	currentMeasurement, err := strconv.Atoi(string(scanner.Text()))
	if err != nil {
		log.Fatal("atoi: ", err)
	}
	return currentMeasurement
}
