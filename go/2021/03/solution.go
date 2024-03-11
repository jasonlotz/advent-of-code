package main

import (
	"fmt"
	"strconv"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2021/03/input.txt"
var testInputFile = "../../../input-files/2021/03/input-sample.txt"
var isTestMode = false

func getInput() []string {
	file := inputFile

	if isTestMode {
		file = testInputFile
	}

	return utils.ProcessStringLinesFile(file)
}

func main() {
	part1()
}

func part1() {
	input := getInput()

	countOnes := make([]int, len(input[0]))
	countZeroes := make([]int, len(input[0]))

	for _, line := range input {
		for i, char := range line {
			if char == '1' {
				countOnes[i]++
			} else {
				countZeroes[i]++
			}
		}
	}

	fmt.Println("Zeroes:", countZeroes)
	fmt.Println("Ones:", countOnes)

	gamma, epsilon := calcGammaAndEpsilon(countZeroes, countOnes)

	fmt.Println("Gamma", gamma)
	fmt.Println("Epsilon:", epsilon)

	fmt.Println("Part 1:", gamma*epsilon)
}

func calcGammaAndEpsilon(countZeroes []int, countOnes []int) (gamma int64, epsilon int64) {
	gammaStr := ""
	epsilonStr := ""

	for i, c := range countZeroes {
		if c > countOnes[i] {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}

	gamma, _ = strconv.ParseInt(gammaStr, 2, 64)
	epsilon, _ = strconv.ParseInt(epsilonStr, 2, 64)

	return gamma, epsilon
}
