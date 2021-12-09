package d2p2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	// measurements or test
	file = "measurements"
)

func Run() {
	measurements, err := os.Open("day2/d2p2/" + file + ".txt")
	if err != nil {
		log.Fatal("open: ", err)
	}
	defer measurements.Close()

	var depth, horizon, aim int
	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")
		if len(cmd) != 2 {
			log.Fatalf("split failed, scanned: %s", cmd)
		}

		moveAmount, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatal("atoi: ", err)
		}

		switch cmd[0] {
		case "up":
			aim -= moveAmount
		case "down":
			aim += moveAmount
		case "forward":
			horizon += moveAmount
			depth += moveAmount * aim
		default:
			log.Fatalf("unknown cmd: %s", cmd)
		}
	}
	log.Printf("depth: %d * horizon: %d = %d", depth, horizon, depth*horizon)
}
