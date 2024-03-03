package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var INPUT_FILE = "../../../input-files/2020/02/input.txt"

type PasswordPolicy struct {
	min      int
	max      int
	letter   rune
	password string
}

func main() {
	input := utils.ProcessStringLinesFile(INPUT_FILE)
	count1 := 0
	count2 := 0

	for _, line := range input {
		policy := parsePasswordPolicy(line)
		if policy.validate() {
			count1++
		}
		if policy.validate2() {
			count2++
		}
	}
	fmt.Println("Part 1:", count1)
	fmt.Println("Part 2:", count2)
}

func parsePasswordPolicy(input string) PasswordPolicy {
	var p PasswordPolicy

	_, err := fmt.Sscanf(input, "%d-%d %c: %s", &p.min, &p.max, &p.letter, &p.password)
	if err != nil {
		fmt.Println("Error parsing input:", input)
		panic(err)
	}
	return p
}

func (p PasswordPolicy) validate() bool {
	count := 0

	for _, c := range p.password {
		if c == p.letter {
			count++
		}
	}

	return count >= p.min && count <= p.max
}

func (p PasswordPolicy) validate2() bool {
	return (rune(p.password[p.min-1]) == p.letter) != (rune(p.password[p.max-1]) == p.letter)
}
