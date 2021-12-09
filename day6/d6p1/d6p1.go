package d6p1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	// measurements or test
	file = "measurements"
)

// wrong 362639
func Run() {
	measurements, err := os.Open("day6/" + file + ".txt")
	if err != nil {
		log.Fatal("open: ", err)
	}
	defer measurements.Close()
	scanner := bufio.NewScanner(measurements)
	lanternfish := readLanternfish(scanner)
	days := 256
	start := time.Now()
	for day := 0; day < days; day++ {
		dayPassed(&lanternfish)
		//log.Printf("After %d day: %v\n", day, lanternfish)
	}
	sum := 0
	for _, v := range lanternfish {
		sum += v
	}
	elapsed := time.Since(start)
	log.Printf("After %d days the total number of lanternfish is %d, in %v", days, sum, elapsed)
}

func dayPassed(lanternfish *[9]int) {
	var toHatch int
	for daysToHatch, _ := range lanternfish {
		if lanternfish[daysToHatch] == 0 {
			continue
		}
		switch daysToHatch {
		case 0:
			toHatch += lanternfish[0]
			lanternfish[0] = 0
		default:
			lanternfish[daysToHatch-1] = lanternfish[daysToHatch]
			lanternfish[daysToHatch] = 0
		}
	}
	lanternfish[6] += toHatch
	lanternfish[8] += toHatch
}

func parseInt(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("atoi: ", err)
	}
	return number
}

func readLanternfish(scanner *bufio.Scanner) [9]int {
	scanner.Scan()
	cmd := strings.Split(scanner.Text(), ",")
	var numbers [9]int
	for _, n := range cmd {
		numbers[parseInt(n)] += 1
	}
	return numbers
}
