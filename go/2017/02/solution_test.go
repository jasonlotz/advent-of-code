package main

import (
	"testing"
)

func TestFindMinMax(t *testing.T) {
	input := []int{5, 1, 9, 5}
	expectedMin := 1
	expectedMax := 9
	resultMin, resultMax := findMinMax(input)
	testMinMaxResult(input, expectedMin, expectedMax, resultMin, resultMax, t)

	input = []int{7, 5, 3}
	expectedMin = 3
	expectedMax = 7
	resultMin, resultMax = findMinMax(input)
	testMinMaxResult(input, expectedMin, expectedMax, resultMin, resultMax, t)

	input = []int{2, 4, 6, 8}
	expectedMin = 2
	expectedMax = 8
	resultMin, resultMax = findMinMax(input)

	testMinMaxResult(input, expectedMin, expectedMax, resultMin, resultMax, t)
}

func testMinMaxResult(input []int, expectedMin int, expectedMax int, resultMin int, resultMax int, t *testing.T) {
	if expectedMin != resultMin {
		t.Error(input, "-> expected min of", expectedMin, "but got", resultMin)
	}
	if expectedMax != resultMax {
		t.Error(input, "-> expected max of", expectedMax, "but got", resultMax)
	}
}

func TestFindDivisible(t *testing.T) {
	input := []int{5, 9, 2, 8}
	expectedI := 8
	expectedJ := 2
	resultI, resultJ := findDivisible(input)
	testDivisibleResult(input, expectedI, expectedJ, resultI, resultJ, t)

	input = []int{9, 4, 7, 3}
	expectedI = 9
	expectedJ = 3
	resultI, resultJ = findDivisible(input)
	testDivisibleResult(input, expectedI, expectedJ, resultI, resultJ, t)

	input = []int{3, 8, 6, 5}
	expectedI = 6
	expectedJ = 3
	resultI, resultJ = findDivisible(input)
	testDivisibleResult(input, expectedI, expectedJ, resultI, resultJ, t)
}

func testDivisibleResult(input []int, expectedI int, expectedJ int, resultI int, resultJ int, t *testing.T) {
	if expectedI != resultI {
		t.Error(input, "-> expected i of", expectedI, "but got", resultI)
	}
	if expectedJ != resultJ {
		t.Error(input, "-> expected j of", expectedJ, "but got", resultJ)
	}
}
