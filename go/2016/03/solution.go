package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var INPUT_FILE = "../../../input-files/2016/03/input.txt"

func main() {
	part1()
	part2()
}

func part1() {
	input := utils.ProcessIntsLinesFile(INPUT_FILE)
	validTriangles := countTrianglesHorizontal(input)
	fmt.Println("Part 1:", validTriangles)
}

func part2() {
	input := utils.ProcessIntsLinesFile(INPUT_FILE)
	validTriangles := countTrianglesVertical(input)
	fmt.Println("Part 2:", validTriangles)
}

func isTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

func countTrianglesHorizontal(input [][]int) int {
	validTriangles := 0
	for _, triangle := range input {
		if isTriangle(triangle[0], triangle[1], triangle[2]) {
			validTriangles++
		}
	}
	return validTriangles
}

func countTrianglesVertical(input [][]int) int {
	validTriangles := 0

	for i := 0; i < len(input); i += 3 {
		for j := 0; j < 3; j++ {
			if isTriangle(input[i][j], input[i+1][j], input[i+2][j]) {
				validTriangles++
			}
		}
	}
	return validTriangles
}
