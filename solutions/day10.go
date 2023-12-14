/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
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
	'!': {{0, -1}, {0, 1}, {-1, 0}, {1, 0},
		{-1, -1}, {1, 1}, {-1, 1}, {1, -1}}, // all + Diagonals
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
				return distance, true
			}
			continue
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

func FloodFill(x int, y int, input []string) (int, [][]int) {
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
		nextNeighbours := charMap[rune(char)]
		matchA := (nextNeighbours[0][0]+nx == x && nextNeighbours[0][1]+ny == y)
		matchB := (nextNeighbours[1][0]+nx == x && nextNeighbours[1][1]+ny == y)
		if matchA || matchB {
			maxDist, isLoop := FloodRecursive(nx, ny, input, 2, visited)
			if isLoop {
				return maxDist / 2, visited
			}
		}
	}
	return -1, [][]int{}
}

func SolveDay10Part1(input []string) (int, [][]int) {
	x, y := FindStart(input)
	return FloodFill(x, y, input)
}

var cornerPairs = map[rune]rune{
	'L': '7',
	'F': 'J',
	'J': 'F',
	'7': 'L',
}

func CanReachOutsideViaCrack(x int, y int, loopMap [][]int, input []string, outside [][]int) bool {
	cardinal := charMap['S']
	results := [4]int{}
	lastVal := -1
	waitFor := '.'
	for i, dir := range cardinal {
		nx, ny := x+dir[0], y+dir[1]
		for nx >= 0 && nx < len(loopMap[0]) &&
			ny >= 0 && ny < len(loopMap) {
			val := loopMap[ny][nx]
			if val == 0 {
				if lastVal != -1 {
					results[i]++
					lastVal = -1
				}
				outVal := outside[ny][nx]
				if outVal > 1 {
					if results[i]%2 == 0 {
						return outVal == 3
					}
					return outVal == 2
				}
			}
			char := rune(input[ny][nx])
			if val != lastVal-1 && val != lastVal+1 {
				lastVal = val
				if char != '|' && char != '-' {
					waitFor = cornerPairs[char]
				} else {
					results[i]++
				}
				if nx == 0 || nx == len(loopMap[0])-1 ||
					ny == 0 || ny == len(loopMap)-1 {
					return results[i]%2 == 0
				}
			} else {
				if waitFor != '.' {
					if char != '|' && char != '-' {
						waitFor = '.'
						if char == waitFor {
							results[i]++
						}
					}
				}
			}
			nx += dir[0]
			ny += dir[1]
		}
		results[i] /= 2
	}
	minimum := results[0]
	for i := 1; i < 4; i++ {
		if results[i] < minimum {
			minimum = results[i]
		}
	}
	return minimum%2 == 0
}
func CanReachEdge(x int, y int, loopMap [][]int, input []string, outside [][]int) bool {
	if outside[y][x] == 1 {
		return false
	}
	if outside[y][x] == 2 {
		return false
	}
	if outside[y][x] == 3 ||
		x == 0 || x == len(loopMap[0])-1 ||
		y == 0 || y == len(loopMap)-1 {
		outside[y][x] = 3
		return true
	}
	outside[y][x] = 1
	found := 2
	neighbours := charMap['!']
	for _, neighbour := range neighbours {
		nx, ny := x+neighbour[0], y+neighbour[1]
		if nx < 0 || nx >= len(loopMap[0]) ||
			ny < 0 || ny >= len(loopMap) {
			continue
		}
		if loopMap[ny][nx] == 0 {
			if CanReachEdge(nx, ny, loopMap, input, outside) {
				found = 3
				break
			}
		}

	}
	if found == 99 {
		if CanReachOutsideViaCrack(x, y, loopMap, input, outside) {
			found = 3
		}
	}
	outside[y][x] = found
	for _, neighbour := range neighbours {
		outside[y+neighbour[1]][x+neighbour[0]] = found
	}
	return found == 3
}

func SolveDay10Part2(loopMap [][]int, input []string) int {
	w := len(loopMap[0])
	outside := make([][]int, len(loopMap))
	for i := range outside {
		outside[i] = make([]int, w)
	}
	insideCount := 0
	for y, line := range loopMap {
		for x, val := range line {
			if val != 0 {
				continue
			}
			if !CanReachEdge(x, y, loopMap, input, outside) {
				insideCount++
			}
		}
	}
	return insideCount
}

func Day10(input []string) []string {
	part1, visited := SolveDay10Part1(input)
	part2 := SolveDay10Part2(visited, input)
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
