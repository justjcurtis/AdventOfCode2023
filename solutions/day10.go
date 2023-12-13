/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"fmt"
	"strconv"
)

var charMap = map[rune][][]int{
	'|': {{0, -1}, {0, 1}},                  // up down
	'-': {{-1, 0}, {1, 0}},                  // left right
	'L': {{0, -1}, {1, 0}},                  // up right
	'J': {{0, -1}, {-1, 0}},                 // up left
	'7': {{0, 1}, {-1, 0}},                  // down left
	'F': {{0, 1}, {1, 0}},                   // down right
	'S': {{0, -1}, {0, 1}, {-1, 0}, {1, 0}}, // all
}

func FindStart(input []string) (int, int) {
	for y, line := range input {
		for x, char := range line {
			if char == 'S' {
				return x, y
			}
		}
	}
	return -1, -1
}

func FloodRecursive(x int, y int, input []string, distance int, visited [][]int) (int, bool) {
	neighbours := charMap[rune(input[y][x])]
	visited[y][x] = distance
	for _, neighbour := range neighbours {
		nx, ny := x+neighbour[0], y+neighbour[1]
		if nx < 0 || nx >= len(input[0]) || ny < 0 || ny >= len(input) {
			continue
		}
		visitVal := visited[ny][nx]
		if visitVal != 0 {
			if visitVal == 1 && distance > 2 {
				// for i := range visited {
				// 	fmt.Println(visited[i])
				// }
				return distance, true
			}
		}
		char := input[ny][nx]
		if char == '.' {
			continue
		}
		maxDist, isLoop := FloodRecursive(nx, ny, input, distance+1, visited)
		if isLoop {
			return maxDist, true
		}
	}
	return -1, false
}

func FloodFill(x int, y int, input []string) int {
	results := []int{}
	visited := make([][]int, len(input))
	for i := range visited {
		visited[i] = make([]int, len(input[0]))
	}
	neighbours := charMap['S']
	visited[y][x] = 1
	for _, neighbour := range neighbours {
		nx, ny := x+neighbour[0], y+neighbour[1]
		if nx < 0 || nx >= len(input[0]) || ny < 0 || ny >= len(input) {
			continue
		}
		char := input[ny][nx]
		if char == '.' {
			continue
		}
		maxDist, isLoop := FloodRecursive(nx, ny, input, 2, visited)
		if isLoop {
			results = append(results, maxDist/2)
		}
	}
	fmt.Println(results)
	fmt.Println(visited)
	return -1
}

func SolveDay10Part1(input []string) int {
	x, y := FindStart(input)
	result := FloodFill(x, y, input)
	return result
}

func Day10(input []string) []string {
	part1 := SolveDay10Part1(input)
	return []string{strconv.Itoa(part1)}
}
