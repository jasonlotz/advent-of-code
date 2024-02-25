package main

import "testing"

func TestGetCoordForSpiralGrid(t *testing.T) {
	tests := []struct {
		input     int
		expectedX int
		expectedY int
	}{
		{1, 0, 0},
		{12, 2, 1},
		{23, 0, -2},
		{1024, -15, 16},
	}

	for _, test := range tests {
		resultX, resultY := getCoordForSpiralGrid(test.input)
		if resultX != test.expectedX || resultY != test.expectedY {
			t.Errorf("%d -> expected %d, %d but got %d, %d", test.input, test.expectedX, test.expectedY,
				resultX, resultY)
		}
	}
}
