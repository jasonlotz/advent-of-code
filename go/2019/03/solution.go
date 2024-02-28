package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var INPUT_FILE = "../../../input-files/2019/03/input.txt"

type Instruction struct {
	direction string
	distance  int
}

var directions = map[string][2]int{
	"U": {0, 1},
	"D": {0, -1},
	"L": {-1, 0},
	"R": {1, 0},
}

func main() {
	part1()
}

func part1() {
	input := utils.ProcessStringLinesFile(INPUT_FILE)

	instructions1 := processInstructions(strings.Split(input[0], ","))
	instructions2 := processInstructions(strings.Split(input[1], ","))

	fmt.Println(instructions1)
	fmt.Println(instructions2)

	// TODO: Implement solution
	// process instructions 1 tracking each location in a map/set
	// process instructions 2 tracking each location in a map/set
	// find the intersection of the two maps
}

func processInstructions(stringInstructions []string) []Instruction {
	instructions := make([]Instruction, len(stringInstructions))

	for i, rawInstruction := range stringInstructions {
		direction := string(rawInstruction[0])
		distance, err := strconv.Atoi(string(rawInstruction[1:]))
		if err != nil {
			panic(err)
		}

		instructions[i] = Instruction{direction, distance}
	}

	return instructions
}
