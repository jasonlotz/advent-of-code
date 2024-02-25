package main

import (
	"fmt"
	"github.com/jasonlotz/advent-of-code/go/utils"
)

var INPUT_FILE = "../../../input-files/2017/05/input.txt"

func main() {
	part1()
	part2()
}

func part1() {
	jumpOffsets := utils.ProcessIntLinesFile(INPUT_FILE)
	jumpCount := processJumpsPart1(jumpOffsets)
	fmt.Println("Jumps to escape (part 1):", jumpCount)
}

func part2() {
	jumpOffsets := utils.ProcessIntLinesFile(INPUT_FILE)
	jumpCount := processJumpsPart2(jumpOffsets)
	fmt.Println("Jumps to escape: (part 2)", jumpCount)
}

func processJumpsPart1(jumps []int) int {
	steps := 0
	currentIndex := 0

	for currentIndex >= 0 && currentIndex < len(jumps) {
		jump := jumps[currentIndex]
		jumps[currentIndex]++
		currentIndex += jump
		steps++
	}

	return steps
}

func processJumpsPart2(jumps []int) int {
	steps := 0
	currentIndex := 0

	for currentIndex >= 0 && currentIndex < len(jumps) {
		jump := jumps[currentIndex]
		if jump >= 3 {
			jumps[currentIndex]--
		} else {
			jumps[currentIndex]++
		}
		currentIndex += jump
		steps++
	}

	return steps
}
