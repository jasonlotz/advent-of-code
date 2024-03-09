package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2018/03/input.txt"

type Claim struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func main() {
	input := utils.ProcessStringLinesFile(inputFile)
	claims := processInput(input)

	fabric := make([][]int, 1000)
	for i := range fabric {
		fabric[i] = make([]int, 1000)
	}

	processClaims(claims, fabric)

	overlapping := countOverlap(fabric)

	fmt.Println("Part 1 (overlapping):", overlapping)

	nonOverlapping := findNonOverlapping(claims, fabric)
	fmt.Println("Part 2 (nonOverlapping):", nonOverlapping)
}

func processInput(input []string) []Claim {
	claims := make([]Claim, len(input))

	for i, line := range input {
		claims[i] = createClaim(line)
	}

	return claims
}

func createClaim(s string) Claim {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(s, -1)

	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}

	if len(numbers) != 5 {
		panic("Invalid claim")
	}

	return Claim{numbers[0], numbers[1], numbers[2], numbers[3], numbers[4]}
}

func processClaims(claims []Claim, fabric [][]int) {
	for _, claim := range claims {
		for x := claim.x; x < claim.x+claim.width; x++ {
			for y := claim.y; y < claim.y+claim.height; y++ {
				fabric[x][y]++
			}
		}
	}
}

func countOverlap(fabric [][]int) int {
	overlapping := 0
	for _, row := range fabric {
		for _, cell := range row {
			if cell > 1 {
				overlapping++
			}
		}
	}
	return overlapping
}

func findNonOverlapping(claims []Claim, fabric [][]int) int {
	for _, claim := range claims {
		overlapping := false
		for x := claim.x; x < claim.x+claim.width; x++ {
			for y := claim.y; y < claim.y+claim.height; y++ {
				if fabric[x][y] > 1 {
					// break out of the inner loop for this claim
					overlapping = true
					break
				}
			}
			// break out of the outer loop for this claim
			if overlapping {
				break
			}
		}

		// if we didn't break out of the loops, then this claim is not overlapping
		if !overlapping {
			return claim.id
		}
	}

	return -1
}
