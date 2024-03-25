package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2021/04/input.txt"
var testInputFile = "../../../input-files/2021/04/input-sample.txt"
var isTestMode = false

type Board [][]int

func getInput() (picks []int, boards []Board) {
	file := inputFile

	if isTestMode {
		file = testInputFile
	}

	splitInput := strings.Split(strings.TrimSpace(utils.ProcessFullFile(file)), "\n\n")

	for _, pickStr := range strings.Split(splitInput[0], ",") {
		pick, err := strconv.Atoi(string(pickStr))
		if err != nil {
			log.Fatal(err)
		}

		picks = append(picks, pick)
	}

	for _, boardStr := range splitInput[1:] {
		var board [][]int
		for _, rowStr := range strings.Split(boardStr, "\n") {
			var row []int
			for _, cellStr := range strings.Fields(rowStr) {
				cell, err := strconv.Atoi(cellStr)
				if err != nil {
					log.Fatal(err)
				}
				row = append(row, cell)
			}

			board = append(board, row)
		}

		boards = append(boards, board)
	}

	return picks, boards
}

func main() {
	picks, boards := getInput()

	for _, pick := range picks {
		boards = processPick(pick, boards)
	}
}

func processPick(pick int, boards []Board) []Board {
	nonwinningBoards := make([]Board, 0)

	for _, board := range boards {
		for i, row := range board {
			for j, cell := range row {
				if cell == pick {
					board[i][j] = -1
				}
			}
		}

		if checkBoard(board) {
			score := scoreBoard(board)

			fmt.Println("Bingo!")
			printBoard(board)
			fmt.Println("Score:", score)
			fmt.Println("Pick:", pick)
			fmt.Println("Score * Pick:", score*pick)
		} else {
			nonwinningBoards = append(nonwinningBoards, board)
		}
	}

	return nonwinningBoards
}

func checkBoard(board Board) bool {
	// Check rows
	for _, row := range board {
		rowComplete := true
		for _, cell := range row {
			if cell != -1 {
				rowComplete = false
			}
		}

		if rowComplete {
			return true
		}
	}

	// Check colums
	for i := 0; i < len(board); i++ {
		columnComplete := true
		for j := 0; j < len(board); j++ {
			if board[j][i] != -1 {
				columnComplete = false
			}
		}

		if columnComplete {
			return true
		}
	}

	return false
}

func scoreBoard(board Board) int {
	score := 0
	for _, row := range board {
		for _, cell := range row {
			if cell != -1 {
				score += cell
			}
		}
	}

	return score
}

func printBoard(board Board) {
	fmt.Println("-------------------------------------")
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%d\t", cell)
		}
		fmt.Println()
	}
	fmt.Println("-------------------------------------")
}
