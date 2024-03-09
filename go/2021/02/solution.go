package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2021/02/input.txt"

type Point struct {
	x   int
	y   int
	aim int
}

type Command struct {
	direction string
	amount    int
}

func main() {
	part1()
	part2()
}

func part1() {
	input := utils.ProcessStringLinesFile(inputFile)
	position := Point{x: 0, y: 0}

	for _, line := range input {
		command := parseCommand(line)
		position = move(position, command)
	}

	fmt.Println("Part 1:", position.x, position.y, position.x*position.y)
}

func part2() {
	input := utils.ProcessStringLinesFile(inputFile)
	position := Point{x: 0, y: 0, aim: 0}

	for _, line := range input {
		command := parseCommand(line)
		position = moveWithAim(position, command)
	}

	fmt.Println("Part 2:", position.x, position.y, position.x*position.y)
}

func parseCommand(line string) Command {
	splitLine := strings.Split(line, " ")

	direction := splitLine[0]
	amount, err := strconv.Atoi(splitLine[1])

	if err != nil {
		panic(err)
	}

	return Command{direction: direction, amount: amount}
}

func move(start Point, command Command) Point {
	var change Point

	if command.direction == "forward" {
		change = Point{x: command.amount, y: 0}
	} else if command.direction == "up" {
		change = Point{x: 0, y: -command.amount}
	} else if command.direction == "down" {
		change = Point{x: 0, y: command.amount}
	}

	return Point{x: start.x + change.x, y: start.y + change.y}
}

func moveWithAim(start Point, command Command) Point {
	var change Point

	if command.direction == "forward" {
		change = Point{x: command.amount, y: (start.aim * command.amount), aim: 0}
	} else if command.direction == "up" {
		change = Point{x: 0, y: 0, aim: -command.amount}
	} else if command.direction == "down" {
		change = Point{x: 0, y: 0, aim: command.amount}
	}

	return Point{x: start.x + change.x, y: start.y + change.y, aim: start.aim + change.aim}
}
