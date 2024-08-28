package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2015/03/input.txt"

func getInput() string {
	return utils.ProcessSingleStringLineFile(inputFile)
}

type Position struct {
	X int
	Y int
}

func (p *Position) Move(direction rune) {
	switch direction {
	case '^':
		p.Y++
	case 'v':
		p.Y--
	case '>':
		p.X++
	case '<':
		p.X--
	}
}

func main() {
	visitHouses()
	visitHousesWithRobot()
}

func visitHouses() {
	currentPosition := Position{}
	visitedHouses := map[Position]bool{currentPosition: true}
	visitedHouses[currentPosition] = true

	input := getInput()

	for _, direction := range input {
		currentPosition.Move(direction)
		visitedHouses[currentPosition] = true
	}

	fmt.Println("Part 1:", len(visitedHouses))
}

func visitHousesWithRobot() {
	currentPositionSanta := Position{}
	currentPositionRobot := Position{}
	visitedHouses := map[Position]bool{currentPositionSanta: true}

	input := getInput()

	for i, direction := range input {
		if i%2 == 0 {
			currentPositionSanta.Move(direction)
			visitedHouses[currentPositionSanta] = true
		} else {
			currentPositionRobot.Move(direction)
			visitedHouses[currentPositionRobot] = true
		}
	}

	fmt.Println("Part 2:", len(visitedHouses))
}
