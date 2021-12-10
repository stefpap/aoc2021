package d10p1

import (
	"bufio"
	"log"
	"os"
)

const (
	// measurements or test
	file = "measurements"
)

func Run() {
	measurements, err := os.Open("day10/" + file + ".txt")
	if err != nil {
		log.Fatal("open: ", err)
	}
	defer measurements.Close()
	scanner := bufio.NewScanner(measurements)
	var line string
	var sum int
	var opened []rune
	for scanner.Scan() {
		line = scanner.Text()
		sum += findIllegalPoints(line, &opened)
	}
	log.Printf("sum of illegal characters is: %d", sum)
}

func findIllegalPoints(line string, opened *[]rune) int {
	for _, c := range line {
		switch c {
		case '[', '(', '{', '<':
			*opened = append(*opened, c)
		case ']', ')', '}', '>':
			if (*opened)[len(*opened)-1] == openingPair(c) {
				*opened = (*opened)[:len(*opened)-1]
			} else {
				return getPoints(c)
			}
		default:
			log.Fatalf("unexpected: %c", c)
		}
	}
	return 0
}

func openingPair(c int32) rune {
	switch c {
	case ']':
		return '['
	case ')':
		return '('
	case '}':
		return '{'
	case '>':
		return '<'
	default:
		log.Fatalf("unexpected: %c", c)
	}
	return '*'
}

func getPoints(c int32) int {
	switch c {
	case ']':
		return 57
	case ')':
		return 3
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		return 0
	}
}
