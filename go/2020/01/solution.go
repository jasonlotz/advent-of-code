package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var INPUT_FILE = "../../../input-files/2020/01/input.txt"

func main() {
	part1()
	part2()
}

func part1() {
	input := utils.ProcessIntLinesFile(INPUT_FILE)

	entry1, entry2 := findSumEntries2(input, 2020)

	fmt.Println("Part 1:", entry1, entry2, entry1*entry2)
}

func part2() {
	input := utils.ProcessIntLinesFile(INPUT_FILE)

	entry1, entry2, entry3 := findSumEntries3(input, 2020)

	fmt.Println("Part 2:", entry1, entry2, entry3, entry1*entry2*entry3)
}

func findSumEntries2(numbers []int, target int) (num1, num2 int) {
	seen := make(map[int]bool, 0)

	for _, number := range numbers {
		fmt.Println("Number", number)
		complement := target - number
		fmt.Println("Lookup for", complement)

		if _, ok := seen[complement]; ok {
			fmt.Println("Found complement")
			return number, complement
		}

		seen[number] = true
	}

	return -1, -1
}

func findSumEntries3(numbers []int, target int) (num1, num2, num3 int) {
	for _, number := range numbers {
		complement := target - number
		entry1, entry2 := findSumEntries2(numbers, complement)

		if entry1 != -1 {
			return number, entry1, entry2
		}
	}

	return -1, -1, -1
}
