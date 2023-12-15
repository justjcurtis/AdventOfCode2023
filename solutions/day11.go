/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"AdventOfCode2023/utils"
	"math"
	"strconv"
)

func ExpandDay11(input []string) ([]string, [][]int) {
	galaxies := [][]int{}
	colsWithGalaxies := make([]bool, len(input[0]))
	for i := len(input) - 1; i >= 0; i-- {
		allDots := true
		for j, char := range input[i] {
			if char != '.' {
				galaxies = append(galaxies, []int{j, i})
				colsWithGalaxies[j] = true
				allDots = false
			}
		}
		if allDots {
			input = append(input[:i+1], input[i:]...)
		}
	}
	for j := len(input[0]) - 1; j >= 0; j-- {
		isEmpty := !colsWithGalaxies[j]
		if isEmpty {
			for i := range input {
				input[i] = input[i][:j+1] + input[i][j:]
			}
		}
	}
	return input, galaxies
}

func HeuristicDay11(start []int, end []int) int {
	return int(math.Abs(float64(start[0]-end[0])) + math.Abs(float64(start[1]-end[1])))
}

type QueueItem struct {
	coords []int
	score  int
}

func QueueItemComparator(a QueueItem, b QueueItem) bool {
	return a.score < b.score
}
func QueueItemIsEqual(a QueueItem, b QueueItem) bool {
	return a.coords[0] == b.coords[0] && a.coords[1] == b.coords[1]
}

var dirs = [4][2]int{
	{0, -1}, {0, 1}, {-1, 0}, {1, 0},
}

func GetNeighboursDay11(coords []int, input []string) [][]int {
	neighbours := [][]int{}
	for _, dir := range dirs {
		x, y := coords[0]+dir[0], coords[1]+dir[1]
		if x < 0 || x >= len(input[0]) || y < 0 || y >= len(input) {
			continue
		}
		neighbours = append(neighbours, []int{x, y})
	}
	return neighbours
}

func reconstructPathDay11(cameFrom map[int][]int, current []int) [][]int {
	totalPath := [][]int{current}
	for {
		currentId := utils.SzudzikPairing(current[0], current[1])
		current = cameFrom[currentId]
		if current == nil {
			break
		}
		totalPath = append([][]int{current}, totalPath...)
	}
	return totalPath
}

func AstarDay11(start []int, end []int, input []string, h func([]int, []int) int) int {
	openSet := utils.NewMinHeap[QueueItem](QueueItemComparator, QueueItemIsEqual)
	openSet.Push(QueueItem{start, 0})
	cameFrom := map[int][]int{}
	gScore := map[int]int{}
	fScore := map[int]int{}
	gScore[utils.SzudzikPairing(start[0], start[1])] = 0
	fScore[utils.SzudzikPairing(start[0], start[1])] = h(start, end)
	for openSet.Len() > 0 {
		current := openSet.Peek().coords
		currentStr := utils.SzudzikPairing(current[0], current[1])
		if current[0] == end[0] && current[1] == end[1] {
			return len(reconstructPathDay11(cameFrom, current))
		}
		openSet.Pop()
		for _, neighbour := range GetNeighboursDay11(current, input) {
			neighbourId := utils.SzudzikPairing(neighbour[0], neighbour[1])
			tentativeGScore := gScore[currentStr] + 1
			if tentativeGScore < gScore[neighbourId] {
				cameFrom[neighbourId] = current
				gScore[neighbourId] = tentativeGScore
				fScore[neighbourId] = tentativeGScore + h(neighbour, end)
				neighbourItem := QueueItem{neighbour, fScore[neighbourId]}
				if !openSet.Contains(neighbourItem) {
					openSet.Push(neighbourItem)
				}
			}
		}
	}
	return -1
}

func SolveDay11Part1(input []string, galaxies [][]int) int {
	totalPathLength := 0
	for i := 0; i < len(galaxies)-1; i++ {
		a := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			b := galaxies[j]
			pathlength := AstarDay11(a, b, input, HeuristicDay11)
			totalPathLength += pathlength
		}
	}
	return totalPathLength
}

func Day11(input []string) []string {
	expanded, galaxies := ExpandDay11(input)
	part1 := SolveDay11Part1(expanded, galaxies)

	return []string{strconv.Itoa(part1)}
}
