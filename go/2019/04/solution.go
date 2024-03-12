package main

import (
	"fmt"
	"strconv"
)

var inputMin = 128392
var inputMax = 643281

func main() {
	countPart1 := 0
	countPart2 := 0

	for i := inputMin; i <= inputMax; i++ {
		if isValidPassword(i, inputMin, inputMax, false) {
			countPart1++
		}
		if isValidPassword(i, inputMin, inputMax, true) {
			countPart2++
		}
	}

	fmt.Println("Part 1:", countPart1)
	fmt.Println("Part 2:", countPart2)
}

func isValidPassword(password int, min int, max int, limitDoubles bool) bool {

	// It is a six-digit number.
	if password < 100000 || password > 999999 {
		return false
	}

	// The value is within the range given in your puzzle input.
	if password < min || password > max {
		return false
	}

	passwordStr := strconv.Itoa(password)

	hasDouble := false

	if limitDoubles {
		// Part 2: Check if there is a double that is not part of a larger group
		// HACK: Go regexp package does not support lookbehind/capture groups, so we have to do this manually
		for i := 0; i < len(passwordStr)-1; i++ {
			if passwordStr[i] == passwordStr[i+1] {
				if i == 0 || passwordStr[i-1] != passwordStr[i] {
					if i == len(passwordStr)-2 || passwordStr[i+2] != passwordStr[i] {
						hasDouble = true
						break
					}
				}
			}
		}
	} else {
		// Part 1: Two adjacent digits are the same (like 22 in 122345).
		for i := 0; i < len(passwordStr)-1; i++ {
			if passwordStr[i] == passwordStr[i+1] {
				hasDouble = true
				break
			}
		}
	}
	if !hasDouble {
		return false
	}

	// Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
	for i := 0; i < len(passwordStr)-1; i++ {
		if passwordStr[i] > passwordStr[i+1] {
			return false
		}
	}

	return true
}
