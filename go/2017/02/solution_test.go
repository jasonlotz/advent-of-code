package main

import (
	"testing"
)

func TestFindMinMax(t *testing.T) {
	test := []struct {
		input       []int
		expectedMin int
		expectedMax int
	}{
		{[]int{5, 1, 9, 5}, 1, 9},
		{[]int{7, 5, 3}, 3, 7},
		{[]int{2, 4, 6, 8}, 2, 8},
	}

	for _, test := range test {
		resultMin, resultMax := findMinMax(test.input)

		if test.expectedMin != resultMin {
			t.Error(test.input, "-> expected min of", test.expectedMin, "but got", resultMin)
		}
		if test.expectedMax != resultMax {
			t.Error(test.input, "-> expected max of", test.expectedMax, "but got", resultMax)
		}
	}
}

func TestFindDivisible(t *testing.T) {
	test := []struct {
		input     []int
		expectedI int
		expectedJ int
	}{
		{[]int{5, 9, 2, 8}, 8, 2},
		{[]int{9, 4, 7, 3}, 9, 3},
		{[]int{3, 8, 6, 5}, 6, 3},
	}

	for _, test := range test {
		resultI, resultJ := findDivisible(test.input)

		if test.expectedI != resultI {
			t.Error(test.input, "-> expected i of", test.expectedI, "but got", resultI)
		}
		if test.expectedJ != resultJ {
			t.Error(test.input, "-> expected j of", test.expectedJ, "but got", resultJ)
		}
	}
}
