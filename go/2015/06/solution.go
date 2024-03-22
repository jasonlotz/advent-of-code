package main

import (
	"fmt"
	"strings"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2015/06/input.txt"

type point struct {
	x int
	y int
}

type instruction struct {
	action string
	start  point
	end    point
}

type grid [][]int

func getInput() []string {
	file := inputFile

	return utils.ProcessStringLinesFile(file)
}

func main() {
	part1()
	part2()
}

func part1() {
	input := getInput()

	instructions := parseInstructions(input)

	grid := make(grid, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	grid.applyInstructions(instructions)

	lightsOn := grid.countLights()

	fmt.Println("Part 1:", lightsOn)
}

func part2() {
	input := getInput()

	instructions := parseInstructions(input)

	grid := make(grid, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	grid.applyInstructions2(instructions)

	lightsOn := grid.countLights2()

	fmt.Println("Part 2:", lightsOn)
}

func parseInstructions(input []string) []instruction {
	var instructions []instruction

	for _, line := range input {
		var inst instruction

		if strings.HasPrefix(line, "turn on") {
			inst.action = "on"
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &inst.start.x, &inst.start.y, &inst.end.x, &inst.end.y)
		} else if strings.HasPrefix(line, "turn off") {
			inst.action = "off"
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &inst.start.x, &inst.start.y, &inst.end.x, &inst.end.y)
		} else if strings.HasPrefix(line, "toggle") {
			inst.action = "toggle"
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &inst.start.x, &inst.start.y, &inst.end.x, &inst.end.y)
		} else {
			panic("Invalid instruction")
		}

		instructions = append(instructions, inst)
	}

	return instructions
}

func (g *grid) applyInstructions(instructions []instruction) {
	for _, inst := range instructions {
		for x := inst.start.x; x <= inst.end.x; x++ {
			for y := inst.start.y; y <= inst.end.y; y++ {
				switch inst.action {
				case "on":
					(*g)[x][y] = 1
				case "off":
					(*g)[x][y] = 0
				case "toggle":
					(*g)[x][y] = 1 - (*g)[x][y]
				}
			}
		}
	}
}

func (g *grid) countLights() int {
	count := 0

	for _, row := range *g {
		for _, light := range row {
			if light == 1 {
				count++
			}
		}
	}

	return count
}

func (g *grid) applyInstructions2(instructions []instruction) {
	for _, inst := range instructions {
		for x := inst.start.x; x <= inst.end.x; x++ {
			for y := inst.start.y; y <= inst.end.y; y++ {
				switch inst.action {
				case "on":
					(*g)[x][y] += 1
				case "off":
					if (*g)[x][y] >= 1 {
						(*g)[x][y] -= 1
					}
				case "toggle":
					(*g)[x][y] += 2
				}
			}
		}
	}
}

func (g *grid) countLights2() int {
	count := 0

	for _, row := range *g {
		for _, light := range row {
			count += light
		}
	}

	return count
}
