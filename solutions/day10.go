/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"AdventOfCode2023/utils"
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

func GetCharForStart(x int, y int, loopMap [][]int) rune {
	neighbours := charMap['S']
	for _, neighbour := range neighbours {
		nx, ny := x+neighbour[0], y+neighbour[1]
		if nx < 0 || nx >= len(loopMap[0]) || ny < 0 || ny >= len(loopMap) {
			continue
		}
	}
	return ' '
}

var DIRS = map[string][]int{
	"up":    {0, -1},
	"down":  {0, 1},
	"left":  {-1, 0},
	"right": {1, 0},
}

func HandlePossibleUpdate(possibleStarts []rune, newPossibleStarts []rune) []rune {
	if len(possibleStarts) == 0 {
		return newPossibleStarts
	}
	result := []rune{}
	for _, char := range possibleStarts {
		if utils.IndexOf(newPossibleStarts, char) != -1 {
			result = append(result, char)
		}
	}
	return result
}

func UpdatePossibleStarts(possibleStarts []rune, neighbour []int) []rune {
	switch neighbour[0] {
	case 0:
		switch neighbour[1] {
		case -1:
			return HandlePossibleUpdate(possibleStarts, []rune{'J', 'L', '|'})
		case 1:
			return HandlePossibleUpdate(possibleStarts, []rune{'7', 'F', '|'})
		}
		break
	case -1:
		return HandlePossibleUpdate(possibleStarts, []rune{'J', '7', '-'})
	case 1:
		return HandlePossibleUpdate(possibleStarts, []rune{'L', 'F', '-'})
	}
	return possibleStarts
}

func FloodFill(x int, y int, input []string) (int, [][]int, rune) {
	visited := make([][]int, len(input))
	for i := range visited {
		visited[i] = make([]int, len(input[0]))
	}
	neighbours := charMap['S']
	neighbourToUse := []int{}
	visited[y][x] = 1
	possibleStarts := []rune{}
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
			neighbourToUse = neighbour
			possibleStarts = UpdatePossibleStarts(possibleStarts, neighbour)
		}
	}
	nx, ny := x+neighbourToUse[0], y+neighbourToUse[1]
	maxDist, isLoop := FloodRecursive(nx, ny, input, 2, visited)
	if isLoop {
		return maxDist / 2, visited, possibleStarts[0]
	}
	return -1, [][]int{}, 'S'
}

func SolveDay10Part1(input []string) (int, [][]int, rune) {
	x, y := FindStart(input)
	return FloodFill(x, y, input)
}

var cornerPairs = map[rune]rune{
	'L': '7',
	'F': 'J',
	'J': 'F',
	'7': 'L',
}

func CountInsideForLine(loopMap []int, line string, startSymbol rune) int {
	inside := false
	count := 0
	cornerToFind := '.'
	for i, val := range loopMap {
		if val == 0 {
			if inside {
				count++
			}
			continue
		}

		char := rune(line[i])
		if char == 'S' {
			char = startSymbol
		}

		if cornerToFind == '.' {
			if char == '|' {
				inside = !inside
				continue
			}
			cornerToFind = cornerPairs[rune(char)]
			continue
		}

		if char != '-' {
			if char == cornerToFind {
				inside = !inside
			}
			cornerToFind = '.'
			continue
		}
	}
	return count
}

func SolveDay10Part2(loopMap [][]int, input []string, startSymbol rune) int {
	fn := func(j int) int {
		return CountInsideForLine(loopMap[j], input[j], '|')
	}
	return utils.Parallelise(utils.IntAcc, fn, len(input)) - 1
}

func Day10(input []string) []string {
	part1, visited, startSymbol := SolveDay10Part1(input)
	part2 := SolveDay10Part2(visited, input, startSymbol) // 413
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
