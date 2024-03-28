package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2021/06/input.txt"
var testInputFile = "../../../input-files/2021/06/input-sample.txt"
var isTestMode = false

func getInput() []int {
	file := inputFile

	if isTestMode {
		file = testInputFile
	}

	input := strings.TrimSpace(utils.ProcessFullFile(file))
	fishStr := strings.Split(input, ",")

	fish := make([]int, len(fishStr))
	for i, f := range fishStr {
		var err error
		fish[i], err = strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
	}

	return fish
}

func main() {
	part1()
	part2()
}

// Leaving this here for reference for the solution used to solve part 1
// The part 2 solution works for part 1 as well and is more efficient
func part1() {
	input := getInput()

	turns := 80

	fish := input
	for i := 0; i < turns; i++ {
		fmt.Println("Round", i+1)
		fish = playRound1(fish)
	}

	// fmt.Println("Final state (Part 1):", fish)
	fmt.Println("Part 1:", len(fish))
}

// This solution works for both part 1 and part 2
// Instead of tracking the fish timers in an array, track the count of each fish timer
func part2() {
	input := getInput()

	fishCount := make([]int, 9)

	for _, f := range input {
		fishCount[f]++
	}

	turns := 256

	for i := 0; i < turns; i++ {
		fishCount = playRound2(fishCount)
	}

	fmt.Println("Final state (Part 2):", fishCount)

	var total int64
	for i := 0; i < len(fishCount); i++ {
		total += int64(fishCount[i])
	}

	fmt.Println("Part 2:", total)
}

func playRound1(fish []int) []int {
	out := make([]int, len(fish))

	for i := 0; i < len(fish); i++ {
		if fish[i] == 0 {
			out[i] = 6
			out = append(out, 8)
		} else {
			out[i] = fish[i] - 1
		}
	}

	return out
}

func playRound2(fishCount []int) []int {
	out := make([]int, len(fishCount))

	for i := 0; i <= 8; i++ {
		if i == 0 {
			out[6] += fishCount[0]
			out[8] += fishCount[0]
		} else {
			out[i-1] += fishCount[i]
		}
	}

	return out
}
