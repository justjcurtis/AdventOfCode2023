/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
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
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] != '.' {
				galaxies = append(galaxies, []int{j, i})
			}
		}
	}
	return input, galaxies
}

func MinDist(start []int, end []int) int {
	return int(math.Abs(float64(start[0]-end[0])) + math.Abs(float64(start[1]-end[1])))
}

func SolveDay11Part1(input []string, galaxies [][]int) int {
	totalPathLength := 0
	for i := 0; i < len(galaxies)-1; i++ {
		a := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			b := galaxies[j]
			pathlength := MinDist(a, b)
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
