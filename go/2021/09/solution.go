package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2021/09/input.txt"
var testInputFile = "../../../input-files/2021/09/input-sample.txt"
var isTestMode = false

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func getInput() []string {
	file := inputFile

	if isTestMode {
		file = testInputFile
	}

	return utils.ProcessStringLinesFile(file)
}

func main() {
	part1()
	part2()
}

func part1() {
	straightLines := []Line{}

	for _, inputLine := range getInput() {
		line, err := parseLine(inputLine)

		if err != nil {
			log.Fatal("Could not parse the line", err)
		}

		if isStraight(line) {
			straightLines = append(straightLines, *line)
		}
	}

	grid := processLines(straightLines)

	if isTestMode {
		printGrid(grid)
	}

	overlapping := countOverlapping(grid)

	fmt.Println("Part 1 (overlapping):", overlapping)
}

func part2() {
	straightOrDiagLines := []Line{}

	for _, inputLine := range getInput() {
		line, err := parseLine(inputLine)

		if err != nil {
			log.Fatal("Could not parse the line", err)
		}

		if isStraight(line) || isDiagonal(line) {
			straightOrDiagLines = append(straightOrDiagLines, *line)
		}
	}

	grid := processLines(straightOrDiagLines)

	if isTestMode {
		printGrid(grid)
	}

	overlapping := countOverlapping(grid)

	fmt.Println("Part 2 (overlapping):", overlapping)
}

func parsePoint(str string) (*Point, error) {
	splitLine := strings.Split(str, ",")

	x, err := strconv.Atoi(splitLine[0])

	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(splitLine[1])

	if err != nil {
		return nil, err
	}

	return &Point{x: x, y: y}, nil
}

func parseLine(line string) (*Line, error) {
	splitLine := strings.Split(line, " -> ")

	start, err := parsePoint(splitLine[0])

	if err != nil {
		return nil, err
	}

	end, err := parsePoint(splitLine[1])

	if err != nil {
		return nil, err
	}

	return &Line{start: *start, end: *end}, nil
}

func isStraight(line *Line) bool {
	return line.start.x == line.end.x || line.start.y == line.end.y
}

func isDiagonal(line *Line) bool {
	xDiff := line.start.x - line.end.x
	yDiff := line.start.y - line.end.y

	return xDiff == yDiff || xDiff == -yDiff
}

func processLines(lines []Line) [][]int {
	size := 1000

	if isTestMode {
		size = 10
	}

	grid := createGrid(size)

	for _, line := range lines {
		if line.start.x == line.end.x {
			// Vertical
			start := min(line.start.y, line.end.y)
			end := max(line.start.y, line.end.y)
			for y := start; y <= end; y++ {
				grid[y][line.start.x]++
			}
		} else if line.start.y == line.end.y {
			// Horizontal
			start := min(line.start.x, line.end.x)
			end := max(line.start.x, line.end.x)
			for x := start; x <= end; x++ {
				grid[line.start.y][x]++
			}
		} else {
			// Diagonal
			points := bresenhamLine(line)
			for _, point := range points {
				grid[point.y][point.x]++
			}
		}
	}

	return grid
}

func createGrid(size int) [][]int {
	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}

	return grid
}

func printGrid(grid [][]int) {
	fmt.Println("Grid:")
	for _, x := range grid {
		for _, y := range x {
			fmt.Printf("%d", y)
		}
		fmt.Println()
	}
}

func countOverlapping(grid [][]int) int {
	overlapping := 0
	for _, x := range grid {
		for _, y := range x {
			if y > 1 {
				overlapping++
			}
		}
	}
	return overlapping
}

// Use Breseham's line algorithm to calculate all the points in a diagonal line
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
func bresenhamLine(line Line) []Point {
	var points []Point

	x1 := line.start.x
	y1 := line.start.y
	x2 := line.end.x
	y2 := line.end.y

	dx := utils.Abs(x2 - x1)
	dy := utils.Abs(y2 - y1)

	sx := utils.Sign(x2 - x1)
	sy := utils.Sign(y2 - y1)

	var err float64
	if dx > dy {
		err = float64(dx) / 2.0
	} else {
		err = float64(-dy) / 2.0
	}

	x, y := x1, y1

	for {
		points = append(points, Point{x: x, y: y})

		if x == x2 && y == y2 {
			break
		}

		e2 := err
		if e2 > -float64(dx) {
			err -= float64(dy)
			x += sx
		}
		if e2 < float64(dy) {
			err += float64(dx)
			y += sy
		}
	}

	return points
}
