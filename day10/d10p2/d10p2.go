package d10p2

import (
	"bufio"
	"log"
	"os"
	"sort"
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
	var points []int
	var opened []rune
	for scanner.Scan() {
		line = scanner.Text()
		if !isIllegal(line, &opened) {
			points = append(points, autocomplete(&opened))
		}
		opened = nil
	}
	log.Printf("sum of illegal characters is: %d", getScore(points))
}

func getScore(points []int) int {
	sort.Ints(points)
	return points[(len(points)-1)/2]
}

func autocomplete(opened *[]rune) int {
	var points int
	for i := len(*opened) - 1; i >= 0; i-- {
		points *= 5
		points += getPoints((*opened)[i])
	}
	return points
}

func isIllegal(line string, opened *[]rune) bool {
	for _, c := range line {
		switch c {
		case '[', '(', '{', '<':
			*opened = append(*opened, c)
		case ']', ')', '}', '>':
			if (*opened)[len(*opened)-1] == openingPair(c) {
				*opened = (*opened)[:len(*opened)-1]
			} else {
				return true
			}
		default:
			log.Fatalf("unexpected: %c", c)
		}
	}
	return false
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
	case '[':
		return 2
	case '(':
		return 1
	case '{':
		return 3
	case '<':
		return 4
	default:
		return 0
	}
}
