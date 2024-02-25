package main

import (
	"fmt"
	"strconv"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var INPUT_FILE = "../../../input-files/2017/01/input.txt"

func main() {
	input := utils.ProcessSingleStringLineFile(INPUT_FILE)

	part1Answer := sumCaptchaNext(input)
	part2Answer := sumCaptchaHalfway(input)

	fmt.Printf("Part 1 answer: %d\n", part1Answer)
	fmt.Printf("Part 2 answer: %d\n", part2Answer)
}

func sumCaptchaNext(input string) int {
	sum := 0
	for i := 1; i < len(input); i++ {
		first, err := strconv.Atoi(string(input[i-1]))
		if err != nil {
			fmt.Println(err)
		}
		second, err := strconv.Atoi(string(input[i]))
		if err != nil {
			fmt.Println(err)
		}

		if first == second {
			sum += second
		}
	}
	last, err := strconv.Atoi(string(input[len(input)-1]))
	if err != nil {
		fmt.Println(err)
	}
	first, err := strconv.Atoi(string(input[0]))
	if err != nil {
		fmt.Println(err)
	}

	if last == first {
		sum += last
	}

	return sum
}

func sumCaptchaHalfway(input string) int {
	sum := 0
	totalElements := len(input)
	fmt.Printf("Total Elements: %d\n", totalElements)

	for i := 0; i < len(input); i++ {
		first, err := strconv.Atoi(string(input[i]))
		if err != nil {
			fmt.Println(err)
		}
		nextNum := i + (totalElements / 2)
		if nextNum >= totalElements {
			nextNum -= totalElements
		}
		second, err := strconv.Atoi(string(input[nextNum]))
		if err != nil {
			fmt.Println(err)
		}

		if first == second {
			sum += second
		}
	}

	return sum
}
