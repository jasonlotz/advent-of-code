package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("input-files/2017/01/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	input := strings.TrimSpace(string(f))
	part1(input)
	part2(input)
}

func part1(input string) {
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

	fmt.Printf("Part 1 answer is: %d\n", sum)
}

func part2(input string) {
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

	fmt.Printf("Part 2 answer is: %d\n", sum)
}
