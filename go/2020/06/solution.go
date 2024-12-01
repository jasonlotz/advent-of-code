package main

import (
	"fmt"
	"strings"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2020/06/input.txt"
var sampleInputFile = "../../../input-files/2020/06/input-sample.txt"
var isSampleMode = false

func getInput() [][]string {
	file := inputFile

	if isSampleMode {
		file = sampleInputFile
	}

	input := strings.TrimSpace(utils.ProcessFullFile(file))

	splitInput := strings.Split(input, "\n\n")

	var result [][]string

	for _, line := range splitInput {
		result = append(result, strings.Split(line, "\n"))
	}

	return result
}

func main() {
	part1()
	part2()
}

func part1() {
	input := getInput()
	result := 0

	for _, group := range input {
		var answers = make(map[rune]bool)

		for _, person := range group {
			for _, answer := range person {
				answers[answer] = true
			}
		}

		for range answers {
			result++
		}
	}

	fmt.Println("Part 1: ", result)
}

func part2() {
	input := getInput()

	results := 0

	for _, group := range input {
		var answers = make(map[rune]int)

		for _, person := range group {
			for _, answer := range person {
				answers[answer]++
			}
		}

		for _, count := range answers {
			if count == len(group) {
				results++
			}
		}
	}

	fmt.Println("Part 2: ", results)
}
