package solutions

import "fmt"

func recursiveSearchUp(input []string, i int, j int) int {
	if i > 0 && input[i-1][j] == '.' {
		return recursiveSearchUp(input, i-1, j)
	}
	return i
}

func tiltUp(input []string) []string {
	for i := 1; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'O' {
				moveTo := recursiveSearchUp(input, i, j)
				if moveTo != i {
					input[moveTo] = input[moveTo][:j] + "O" + input[moveTo][j+1:]
					input[i] = input[i][:j] + "." + input[i][j+1:]
				}
			}
		}
	}
	return input
}

func SolveDay14Part1(input []string) int {
	weight := 0
	tilted := tiltUp(input)
	for i, line := range tilted {
		for _, char := range line {
			if char == 'O' {
				weight += len(input) - i
			}
		}
	}
	return weight
}

func Day14(input []string) []string {
	part1 := SolveDay14Part1(input)
	return []string{fmt.Sprint(part1)}
}
