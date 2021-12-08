package d4p1

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

type cell struct {
	visited bool
	number int
}

func Run() {
	measurements, err := os.Open("day4/d4p1/"+file+".txt")
	if err != nil {
		log.Fatal("open: ", err)
	}
	defer measurements.Close()
	scanner := bufio.NewScanner(measurements)
	numbers := readNumbers(scanner)
	var tables [][5][5]cell
	tables = readTables(scanner)

	for _, number := range numbers {
		log.Println("reading number: ", number)
		for tableNumber, _ := range tables {
			markVisited(&tables[tableNumber], number)
			if isWinner(&tables[tableNumber]) {
				log.Println(sumUnvisited(&tables[tableNumber]) , number, sumUnvisited(&tables[tableNumber]) * number)
				log.Printf("Winner is table: %d, output: %d\n", tableNumber, sumUnvisited(&tables[tableNumber]) * number)
				goto Exit
			}
		}
	}
	Exit:
}

func markVisited(table *[5][5]cell, number int) {
	for i, row := range table {
		for j, cell := range row {
			if cell.number == number {
				table[i][j].visited = true
				return
			}
		}
	}
	return
}

func sumUnvisited(table *[5][5]cell) int {
	sum := 0
	for _, row := range table {
		for _, cell := range row {
			if cell.visited == false {
				sum += cell.number
			}
		}
	}
	return sum
}

func isWinner(table *[5][5]cell) bool {
	//var int check
	for check := 0; check < 5; check++ {
		if table[check][0].visited &&
			table[check][1].visited &&
			table[check][2].visited &&
			table[check][3].visited &&
			table[check][4].visited {
			return true
		}
		if table[0][check].visited &&
			table[1][check].visited &&
			table[2][check].visited &&
			table[3][check].visited &&
			table[4][check].visited {
			return true
		}
	}

	return false
}

func readTables(scanner *bufio.Scanner) [][5][5]cell {
	var table [5][5]cell
	tables := make([][5][5]cell, 0)
	rowNumber := -1
	for scanner.Scan() {
		if rowNumber == -1 {
			rowNumber++
			continue
		}

		row := strings.Split(strings.Join(strings.Fields(scanner.Text())," "), " ")
		table[rowNumber] = [5]cell{
			cell{number: parseInt(row[0])},
			cell{number: parseInt(row[1])},
			cell{number: parseInt(row[2])},
			cell{number: parseInt(row[3])},
			cell{number: parseInt(row[4])},
		}
		if rowNumber == 4 {
			tables = append(tables, table)
			rowNumber = -1
			continue
		}
		rowNumber++
	}
	return tables
}

func parseInt(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("atoi: ", err)
	}
	return number
}

func readNumbers(scanner *bufio.Scanner) []int {
	scanner.Scan()
	cmd := strings.Split(scanner.Text(), ",")
	var numbers []int
	for _, n := range cmd {

		numbers = append(numbers, parseInt(n))
	}
	return numbers
}

func printTable (table [5][5]cell) {
	for _, row := range table {
		log.Println(row[0], row[1], row[2], row[3], row[4])
	}
	log.Println()
}
