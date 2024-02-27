package main

import (
	"testing"
)

func TestCalcFuel(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, test := range tests {
		if result := calcFuel(test.input); result != test.expected {
			t.Errorf("calcFuel(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}

func TestCalcFuelRecursive(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, test := range tests {
		if result := calcFuelRecursive(test.input); result != test.expected {
			t.Errorf("calcFuelRecursive(%v) = %v, want %v", test.input, result, test.expected)
		}
	}
}
