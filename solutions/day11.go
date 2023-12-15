/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"AdventOfCode2023/utils"
	"math"
)

func ExpandDay11(input []string) []string {
	colsWithGalaxies := make([]bool, len(input[0]))
	for i := len(input) - 1; i >= 0; i-- {
		allDots := true
		for j, char := range input[i] {
			if char != '.' {
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
	return input
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
			return gScore[currentStr]
		}
		openSet.Pop()
		for _, neighbour := range GetNeighboursDay11(current, input) {
			neighbourId := utils.SzudzikPairing(neighbour[0], neighbour[1])
			tentativeGScore := gScore[currentStr] + 1
			if tentativeGScore < gScore[neighbourId] {
				cameFrom[neighbourId] = current
				gScore[neighbourId] = tentativeGScore
				fScore[neighbourId] = tentativeGScore + h(neighbour, end)
				if !openSet.Contains(neighbour) {
					openSet.Push(QueueItem{neighbour, fScore[neighbourId]})
				}
			}
		}
	}
}

func Day11(input []string) []string {
	expanded := ExpandDay11(input)

	return []string{}
}
