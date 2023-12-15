/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"strconv"
)

func ExpandDay11(input []string) ([]map[int]bool, [][]int) {
	galaxies := [][]int{}
	colsWithGalaxies := make([]bool, len(input[0]))
	expandMap := []map[int]bool{
		make(map[int]bool),
		make(map[int]bool),
	}
	for j := len(input) - 1; j >= 0; j-- {
		allDots := true
		for i, char := range input[j] {
			if char != '.' {
				colsWithGalaxies[i] = true
				allDots = false
			}
		}
		if allDots {
			expandMap[0][j] = true
		}
	}
	for i := len(input[0]) - 1; i >= 0; i-- {
		isEmpty := !colsWithGalaxies[i]
		if isEmpty {
			expandMap[1][i] = true
		}
	}
	for j := 0; j < len(input); j++ {
		for i := 0; i < len(input[0]); i++ {
			if input[j][i] != '.' {
				galaxies = append(galaxies, []int{j, i})
			}
		}
	}
	return expandMap, galaxies
}

func MinDist(start []int, end []int, expandMap []map[int]bool, expansionFactor int) int {
	minJ := start[0]
	maxJ := end[0]
	if start[0] > end[0] {
		minJ = end[0]
		maxJ = start[0]
	}
	minI := start[1]
	maxI := end[1]
	if start[1] > end[1] {
		minI = end[1]
		maxI = start[1]
	}
	pathLength := (maxI - minI) + (maxJ - minJ)
	for j := minJ; j < maxJ; j++ {
		if expandMap[0][j] {
			pathLength += expansionFactor - 1
		}
	}
	for i := minI; i < maxI; i++ {
		if expandMap[1][i] {
			pathLength += expansionFactor - 1
		}
	}
	return pathLength
}

func SolveDay11Part1(expandMap []map[int]bool, galaxies [][]int, expansionFactor int) int {
	totalPathLength := 0
	for i := 0; i < len(galaxies)-1; i++ {
		a := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			b := galaxies[j]
			pathlength := MinDist(a, b, expandMap, expansionFactor)
			totalPathLength += pathlength
		}
	}
	return totalPathLength
}

func Day11(input []string) []string {
	expandMap, galaxies := ExpandDay11(input)
	part1 := SolveDay11Part1(expandMap, galaxies, 2)
	part2 := SolveDay11Part1(expandMap, galaxies, 1000000)

	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
