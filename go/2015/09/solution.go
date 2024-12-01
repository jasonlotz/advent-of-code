package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/jasonlotz/advent-of-code/go/utils"
)

var inputFile = "../../../input-files/2015/09/input.txt"
var sampleInputFile = "../../../input-files/2015/09/input-sample.txt"
var isSampleMode = false

type path struct {
	From     string
	To       string
	Distance int
}

type mapGraph map[string]map[string]int

func getInput() []string {
	flag.BoolVar(&isSampleMode, "sample", false, "Use the sample input")
	flag.Parse()

	file := inputFile
	if isSampleMode {
		file = sampleInputFile
	}

	return utils.ProcessStringLinesFile(file)
}

func parseInput(input []string) []path {
	paths := []path{}

	for _, line := range input {
		var from, to string
		var distance int
		fmt.Sscanf(line, "%s to %s = %d", &from, &to, &distance)

		paths = append(paths, path{from, to, distance})
	}

	return paths
}

func buildGraph(paths []path) map[string]map[string]int {
	graph := map[string]map[string]int{}

	for _, path := range paths {
		if _, ok := graph[path.From]; !ok {
			graph[path.From] = map[string]int{}
		}

		if _, ok := graph[path.To]; !ok {
			graph[path.To] = map[string]int{}
		}

		graph[path.From][path.To] = path.Distance
		graph[path.To][path.From] = path.Distance
	}

	return graph
}

func travelingSalesman(graph mapGraph) (int, int) {
	min := math.MaxInt32
	max := 0

	for k := range graph {
		dfsMin, dfsMax := dfsTotalDistance(graph, k, map[string]bool{k: true})
		if dfsMin < min {
			min = dfsMin
		}
		if dfsMax > max {
			max = dfsMax
		}
	}

	return min, max
}

func dfsTotalDistance(graph mapGraph, entry string, visited map[string]bool) (min, max int) {
	if len(visited) == len(graph) {
		return 0, 0
	}

	minDistance := math.MaxInt32
	maxDistance := 0

	for k := range graph {
		if !visited[k] {
			visited[k] = true

			weight := graph[entry][k]
			minRecurse, maxRecurse := dfsTotalDistance(graph, k, visited)
			weightedMin := minRecurse + weight
			weightedMax := maxRecurse + weight

			if weightedMin < minDistance {
				minDistance = weightedMin
			}
			if weightedMax > maxDistance {
				maxDistance = weightedMax
			}

			delete(visited, k)
		}
	}

	return minDistance, maxDistance
}

func main() {
	input := getInput()
	paths := parseInput(input)

	graph := buildGraph(paths)
	fmt.Println(graph)
	shortest, longest := travelingSalesman(graph)

	fmt.Println("Part 1:", shortest)
	fmt.Println("Part 2:", longest)
}
