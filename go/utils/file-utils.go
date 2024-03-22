package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func openFile(inputFile string) *os.File {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return file
}

func ProcessFullFile(inputFile string) string {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func ProcessSingleStringLineFile(inputFile string) string {
	file := openFile(inputFile)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func ProcessStringLinesFile(inputFile string) []string {
	file := openFile(inputFile)

	scanner := bufio.NewScanner(file)
	processedFile := make([]string, 0)

	for scanner.Scan() {
		processedFile = append(processedFile, scanner.Text())
	}

	return processedFile
}

func ProcessIntLinesFile(inputFile string) []int {
	file := openFile(inputFile)

	scanner := bufio.NewScanner(file)
	processedFile := make([]int, 0)

	for scanner.Scan() {
		intVal, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		processedFile = append(processedFile, intVal)
	}

	return processedFile
}
func ProcessIntsLinesFile(inputFile string) [][]int {
	file := openFile(inputFile)

	scanner := bufio.NewScanner(file)
	processedFile := make([][]int, 0)

	for scanner.Scan() {
		stringVals := strings.Fields(scanner.Text())
		intVals := make([]int, len(stringVals))

		for i, s := range stringVals {
			intVal, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			intVals[i] = intVal
		}
		processedFile = append(processedFile, intVals)
	}

	return processedFile
}
