package main

import (
	"fmt"
	"math"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2015/02/input.txt"

type Present struct {
	length int
	width  int
	height int
}

func (p Present) SurfaceArea() int {
	return 2*p.length*p.width + 2*p.width*p.height + 2*p.height*p.length
}

func (p *Present) ParseDimensions(dimensions string) {
	fmt.Sscanf(dimensions, "%dx%dx%d", &p.length, &p.width, &p.height)
}

func (p Present) SmallestSideArea() int {
	lw := float64(p.length * p.width)
	wh := float64(p.width * p.height)
	hl := float64(p.height * p.length)

	return int(math.Min(lw, math.Min(wh, hl)))
}

func (p Present) Volume() int {
	return p.length * p.width * p.height
}

func (p Present) SmallestPerimeter() int {
	lw := float64(2 * (p.length + p.width))
	wh := float64(2 * (p.width + p.height))
	hl := float64(2 * (p.height + p.length))

	return int(math.Min(lw, math.Min(wh, hl)))
}

func getInput() []string {
	file := inputFile

	return utils.ProcessStringLinesFile(file)
}

func main() {
	input := getInput()
	var totalPaper int
	var totalRibbon int

	for _, line := range input {
		present := Present{}
		present.ParseDimensions(line)

		totalPaper += present.SurfaceArea() + present.SmallestSideArea()
		totalRibbon += present.SmallestPerimeter() + present.Volume()
	}

	fmt.Println("Part 1:", totalPaper)
	fmt.Println("Part 2:", totalRibbon)
}
