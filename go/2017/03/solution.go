package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

type DirectionType string

var Direction = struct {
	RIGHT DirectionType
	UP    DirectionType
	LEFT  DirectionType
	DOWN  DirectionType
}{
	RIGHT: "right",
	UP:    "up",
	LEFT:  "left",
	DOWN:  "down",
}

func main() {
	input := 23
	x, y := getCoordForSpiralGrid(input)
	distance := utils.Abs(x) + utils.Abs(y)
	fmt.Printf("%d  -> %d, %d (%d)\n", input, x, y, distance)
}

func getCoordForSpiralGrid(input int) (int, int) {
	x, y := 0, 0
	rightBound := 1
	upBound := 1
	leftBound := -1
	downBound := -1
	value := 1
	direction := Direction.RIGHT

	for value < input {
		if direction == Direction.RIGHT {
			if x < rightBound {
				x++
				value++
			} else {
				direction = Direction.UP
				rightBound++
			}
		}

		if direction == Direction.UP {
			if y < upBound {
				y++
				value++
			} else {
				direction = Direction.LEFT
				upBound++
			}
		}

		if direction == Direction.LEFT {
			if x > leftBound {
				x--
				value++
			} else {
				direction = Direction.DOWN
				leftBound--
			}
		}

		if direction == Direction.DOWN {
			if y > downBound {
				y--
				value++
			} else {
				direction = Direction.RIGHT
				downBound--
			}
		}
	}

	return x, y
}
