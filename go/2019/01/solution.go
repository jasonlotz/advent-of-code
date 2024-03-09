package main

import (
	"fmt"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2019/01/input.txt"

func main() {
	input := utils.ProcessIntLinesFile(inputFile)

	totalFuel := 0
	totalFuelRecursive := 0

	for _, mass := range input {
		totalFuel += calcFuel(mass)
		totalFuelRecursive += calcFuelRecursive(mass)
	}

	fmt.Println("Part 1:", totalFuel)
	fmt.Println("Part 2:", totalFuelRecursive)
}

func calcFuel(mass int) int {
	return mass/3 - 2
}

func calcFuelRecursive(mass int) int {
	fuel := calcFuel(mass)
	if fuel <= 0 {
		return 0
	}
	return fuel + calcFuelRecursive(fuel)
}
