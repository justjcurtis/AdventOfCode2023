/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"strconv"
)

func ExpandDay11(input []string) ([][]int, [][]int) {
	galaxies := [][]int{}
	colsWithGalaxies := make([]bool, len(input[0]))
	expandMap := [][]int{{}, {}}
	for j := len(input) - 1; j >= 0; j-- {
		allDots := true
		for i, char := range input[j] {
			if char != '.' {
				colsWithGalaxies[i] = true
				allDots = false
			}
		}
		if allDots {
			expandMap[0] = append(expandMap[0], j)
		}
	}
	for i := len(input[0]) - 1; i >= 0; i-- {
		isEmpty := !colsWithGalaxies[i]
		if isEmpty {
			expandMap[1] = append(expandMap[1], i)
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

func MinDist(start []int, end []int, expandMap [][]int, efA int, efB int) (int, int) {
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
	pathLengthA := (maxI - minI) + (maxJ - minJ)
	pathLengthB := pathLengthA
	for _, j := range expandMap[0] {
		if j >= minJ && j < maxJ {
			pathLengthA += efA - 1
			pathLengthB += efB - 1
		}
	}
	for _, i := range expandMap[1] {
		if i >= minI && i < maxI {
			pathLengthA += efA - 1
			pathLengthB += efB - 1
		}
	}
	return pathLengthA, pathLengthB
}

func SolveDay11(expandMap [][]int, galaxies [][]int, efA int, efB int) (int, int) {
	totalPathLengthA := 0
	totalPathLengthB := 0
	for i := 0; i < len(galaxies)-1; i++ {
		a := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			b := galaxies[j]
			pathLengthA, pathLengthB := MinDist(a, b, expandMap, efA, efB)
			totalPathLengthA += pathLengthA
			totalPathLengthB += pathLengthB
		}
	}
	return totalPathLengthA, totalPathLengthB
}

func Day11(input []string) []string {
	expandMap, galaxies := ExpandDay11(input)
	part1, part2 := SolveDay11(expandMap, galaxies, 2, 1000000)

	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
