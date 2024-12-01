package main

import (
	"fmt"
	"sort"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2020/05/input.txt"
var testSampleFile = "../../../input-files/2020/05/input-sample.txt"
var isSampleMode = false

func getInput() []string {
	file := inputFile

	if isSampleMode {
		file = testSampleFile
	}

	return utils.ProcessStringLinesFile(file)
}

func main() {
	input := getInput()
	highestID := 0
	foundSeats := []int{}

	for _, seat := range input {
		_, _, id := decodeSeat(seat)
		foundSeats = append(foundSeats, id)
		if id > highestID {
			highestID = id
		}
	}
	fmt.Println("Part 1:", highestID)
	fmt.Println("Part 2:", findMissingSeat(foundSeats))
}

// Use binary space partitioning algorithm to decode the seat
func decodeSeat(seat string) (row int, col int, id int) {
	row = 0
	col = 0
	for _, c := range seat {
		switch c {
		case 'F':
			row = row << 1
		case 'B':
			row = row<<1 + 1
		case 'L':
			col = col << 1
		case 'R':
			col = col<<1 + 1
		}
	}
	return row, col, row*8 + col
}

func findMissingSeat(seats []int) int {
	sort.Ints(seats)

	for i := 1; i < len(seats); i++ {
		// If the seat before is exactly 2 less, then the missing seat is in between
		if seats[i] == seats[i-1]+2 {
			return seats[i] - 1
		}
	}

	return 0
}
