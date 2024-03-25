package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2022/06/input.txt"
var isTestMode = false

func getInput() string {
	file := inputFile

	return utils.ProcessSingleStringLineFile(file)
}

func main() {
	part1()
	part2()
}

func part1() {
	input := getInput()
	markerIdx := findMarker(input, 4)
	fmt.Println("Part 1:", markerIdx)
}

func part2() {
	input := getInput()
	markerIdx := findMarker(input, 14)
	fmt.Println("Part 2:", markerIdx)
}

func findMarker(datastream string, distinct int) int {
	var markerIdx int

	for markerIdx = distinct; markerIdx < len(datastream); markerIdx++ {
		potentialMarker := datastream[markerIdx-distinct : markerIdx]

		if !hasDuplicates(potentialMarker) {
			return markerIdx
		}
	}

	return markerIdx
}

func hasDuplicates(s string) bool {
	charMap := make(map[rune]bool)
	for _, char := range s {
		if _, exists := charMap[char]; exists {
			return true
		}
		charMap[char] = true
	}
	return false
}
