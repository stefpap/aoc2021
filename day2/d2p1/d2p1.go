package d2p1

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
	measurements, err := os.Open("day2/d2p1/"+file+".txt")
	if err != nil {
		log.Fatal("open: ", err)
	}
	defer measurements.Close()

	var depth, horizon int
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
			depth -= moveAmount
		case "down":
			depth += moveAmount
		case "forward":
			horizon += moveAmount
		default:
			log.Fatalf("unknown cmd: %s", cmd)
		}
	}
	log.Printf("depth: %d * horizon: %d = %d", depth, horizon, depth* horizon)
}
