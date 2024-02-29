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
	input := utils.ProcessStringLinesFile(INPUT_FILE)

	instructions1 := processInstructions(strings.Split(input[0], ","))
	instructions2 := processInstructions(strings.Split(input[1], ","))

	locations1 := getLocations(instructions1)
	locations2 := getLocations(instructions2)

	minDistance := findMinDistance(locations1, locations2)
	fmt.Println("Part 1 (minDistance):", minDistance)

	minSteps := findMinSteps(locations1, locations2)
	fmt.Println("Part 2 (minSteps):", minSteps)
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

func getLocations(instructions []Instruction) map[[2]int]int {
	locations := make(map[[2]int]int)
	currentLocation := [2]int{0, 0}
	steps := 0

	for _, instruction := range instructions {
		for i := 0; i < instruction.distance; i++ {
			steps++
			currentLocation[0] += directions[instruction.direction][0]
			currentLocation[1] += directions[instruction.direction][1]

			if _, ok := locations[currentLocation]; !ok {
				locations[currentLocation] = steps
			}
		}
	}

	return locations
}

func findIntersections(locations1 map[[2]int]int, locations2 map[[2]int]int) map[[2]int]bool {
	intersections := make(map[[2]int]bool)
	for loc := range locations1 {
		if _, ok := locations2[loc]; ok {
			intersections[loc] = true
		}
	}
	return intersections
}

func findMinDistance(locations1 map[[2]int]int, locations2 map[[2]int]int) int {
	minDistance := -1

	intersections := findIntersections(locations1, locations2)

	for loc := range intersections {
		distance := utils.Abs(loc[0]) + utils.Abs(loc[1])
		if minDistance == -1 || distance < minDistance {
			minDistance = distance
		}
	}
	return minDistance
}

func findMinSteps(locations1 map[[2]int]int, locations2 map[[2]int]int) int {
	minSteps := -1

	intersections := findIntersections(locations1, locations2)

	for loc := range intersections {
		steps := locations1[loc] + locations2[loc]
		if minSteps == -1 || steps < minSteps {
			minSteps = steps
		}
	}
	return minSteps
}
