package d1p2

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const (
	// measurements or test
	file = "measurements"
	numberOfMeasurements = 3
)

type measurements []int

type d1p2 struct {
	scanner *bufio.Scanner
	three measurements
}

func Run() {
	measurementsFile, err := os.Open(file + ".txt")
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	defer measurementsFile.Close()

	c := &d1p2{
		scanner: bufio.NewScanner(measurementsFile),
		three: make(measurements, 3),
	}

	var times, scanned int
	var ok bool
	for i := 0; i < numberOfMeasurements; i++ {
		ok = c.canScan()
		if !ok {
			log.Printf("less than %d measurements: %d\n", numberOfMeasurements, times)
			return
		}
		c.three[i] = c.getScanned()
	}
	previousMeasurements := c.three.Some()

	for c.canScan() {
		scanned = c.getScanned()
		c.Push(scanned)
		currentMeasurements := c.three.Some()
		if currentMeasurements > previousMeasurements {
			times++
		}
		previousMeasurements = currentMeasurements
	}
	log.Printf("times: %d", times)
}

func (c *d1p2) canScan() bool {
	return c.scanner.Scan()
}

func (c *d1p2) getScanned() int {
	var currentMeasurement int
	var err error
	currentMeasurement, err = strconv.Atoi(c.scanner.Text())
	if err != nil {
		log.Fatalf("atoi: %v",err)
	}
	return currentMeasurement
}

// Some sums the 3 current measurements. wink wink
func (m measurements) Some() int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}

func (c *d1p2) Push(measurement int) {
	c.three = append(c.three[1:], measurement)
}
