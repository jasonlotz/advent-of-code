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
	part2()
}

func part1() {
	input := getInput()

	zeroes, ones := countByPosition(input)

	fmt.Println("Zeroes:", zeroes)
	fmt.Println("Ones:", ones)

	gamma, epsilon := calcGammaAndEpsilon(zeroes, ones)

	fmt.Println("Gamma", gamma)
	fmt.Println("Epsilon:", epsilon)

	fmt.Println("Part 1:", gamma*epsilon)
}

func part2() {
	input := getInput()

	oxygen, co2 := calcOxygenAndCo2(input)

	fmt.Println("Oxygen:", oxygen)
	fmt.Println("CO2:", co2)
	fmt.Println("Part 2:", oxygen*co2)
}

func countByPosition(input []string) (zeroes []int, ones []int) {
	zeroes = make([]int, len(input[0]))
	ones = make([]int, len(input[0]))

	for _, line := range input {
		for i, char := range line {
			if char == '1' {
				ones[i]++
			} else {
				zeroes[i]++
			}
		}
	}
	return zeroes, ones
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

func calcOxygenAndCo2(input []string) (oxygen int64, co2 int64) {
	length := len(input[0])

	oxygenInput := make([]string, len(input))
	copy(oxygenInput, input)

	for i := range length {
		zeroes, ones := countByPosition(oxygenInput)
		if zeroes[i] > ones[i] {
			oxygenInput = filterByPosition(oxygenInput, i, '0')
		} else if zeroes[i] < ones[i] {
			oxygenInput = filterByPosition(oxygenInput, i, '1')
		} else {
			oxygenInput = filterByPosition(oxygenInput, i, '1')
		}

		if len(oxygenInput) == 1 {
			oxygen, _ = strconv.ParseInt(oxygenInput[0], 2, 64)
			break
		}
	}

	co2Input := make([]string, len(input))
	copy(co2Input, input)

	for i := range length {
		zeroes, ones := countByPosition(co2Input)
		if zeroes[i] > ones[i] {
			co2Input = filterByPosition(co2Input, i, '1')
		} else if zeroes[i] < ones[i] {
			co2Input = filterByPosition(co2Input, i, '0')
		} else {
			co2Input = filterByPosition(co2Input, i, '0')
		}

		if len(co2Input) == 1 {
			co2, _ = strconv.ParseInt(co2Input[0], 2, 64)
			break
		}
	}

	return oxygen, co2
}

func filterByPosition(input []string, position int, char byte) []string {
	if len(input) == 1 {
		return input
	}

	filtered := make([]string, 0)

	for _, line := range input {
		if line[position] == char {
			filtered = append(filtered, line)
		}
	}

	return filtered
}
