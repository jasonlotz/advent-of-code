package main

import (
	"testing"
)

func TestFindMinMax(t *testing.T) {
	input := []int{5, 1, 9, 5}
	expectedMin := 1
	expectedMax := 9
	min, max := findMinMax(input)
	testMinMaxResult(input, expectedMin, expectedMax, min, max, t)

	input = []int{7, 5, 3}
	expectedMin = 3
	expectedMax = 7
	min, max = findMinMax(input)
	testMinMaxResult(input, expectedMin, expectedMax, min, max, t)

	input = []int{2, 4, 6, 8}
	expectedMin = 2
	expectedMax = 8
	min, max = findMinMax(input)
	testMinMaxResult(input, expectedMin, expectedMax, min, max, t)
}

func testMinMaxResult(input []int, expectedMin int, expectedMax int, min int, max int, t *testing.T) {
	if expectedMin != min {
		t.Error(input, "-> expected min of", expectedMin, "but got", min)
	}
	if expectedMax != max {
		t.Error(input, "-> expected max of", expectedMax, "but got", max)
	}
}

func TestFindDivisible(t *testing.T) {
	input := []int{5, 9, 2, 8}
	expectedI := 8
	expectedJ := 2
	i, j := findDivisible(input)
	testDivisibleResult(input, expectedI, expectedJ, i, j, t)

	input = []int{9, 4, 7, 3}
	expectedI = 9
	expectedJ = 3
	i, j = findDivisible(input)
	testDivisibleResult(input, expectedI, expectedJ, i, j, t)

	input = []int{3, 8, 6, 5}
	expectedI = 6
	expectedJ = 3
	i, j = findDivisible(input)
	testDivisibleResult(input, expectedI, expectedJ, i, j, t)

}

func testDivisibleResult(input []int, expectedI int, expectedJ int, i int, j int, t *testing.T) {
	if expectedI != i {
		t.Error(input, "-> expected i of", expectedI, "but got", i)
	}
	if expectedJ != j {
		t.Error(input, "-> expected j of", expectedJ, "but got", j)
	}
}
