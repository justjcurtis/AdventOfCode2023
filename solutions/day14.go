package solutions

import (
	"AdventOfCode2023/utils"
	"fmt"
)

func parseDay14(input []string) [][]int {
	output := make([][]int, len(input))
	fn := func(i int) {
		output[i] = make([]int, len(input[i]))
		for j, char := range input[i] {
			if char == 'O' {
				output[i][j] = 1
			} else if char == '#' {
				output[i][j] = -1
			}
		}
	}
	utils.ParalleliseVoid(fn, len(input))
	return output
}

func recursiveSearchUp(input [][]int, i int, j int) int {
	if i > 0 && input[i-1][j] == 0 {
		return recursiveSearchUp(input, i-1, j)
	}
	return i
}

func tiltUp(input [][]int) [][]int {
	fn := func(j int) {
		for i := 1; i < len(input); i++ {
			if input[i][j] == 1 {
				moveTo := recursiveSearchUp(input, i, j)
				if moveTo != i {
					input[moveTo][j] = 1
					input[i][j] = 0
				}
			}
		}
	}
	utils.ParalleliseVoid(fn, len(input[0]))
	return input
}

func recursiveSearchDown(input [][]int, i int, j int) int {
	if i < len(input)-1 && input[i+1][j] == 0 {
		return recursiveSearchDown(input, i+1, j)
	}
	return i
}

func tiltDown(input [][]int) [][]int {
	fn := func(j int) {
		for i := len(input) - 2; i >= 0; i-- {
			if input[i][j] == 1 {
				moveTo := recursiveSearchDown(input, i, j)
				if moveTo != i {
					input[moveTo][j] = 1
					input[i][j] = 0
				}
			}
		}
	}
	utils.ParalleliseVoid(fn, len(input[0]))
	return input
}

func recursiveSearchLeft(input [][]int, i int, j int) int {
	if j > 0 && input[i][j-1] == 0 {
		return recursiveSearchLeft(input, i, j-1)
	}
	return j
}

func tiltLeft(input [][]int) [][]int {
	fn := func(i int) {
		for j := 1; j < len(input[i]); j++ {
			if input[i][j] == 1 {
				moveTo := recursiveSearchLeft(input, i, j)
				if moveTo != j {
					input[i][moveTo] = 1
					input[i][j] = 0
				}
			}
		}
	}
	utils.ParalleliseVoid(fn, len(input))
	return input
}

func recursiveSearchRight(input [][]int, i int, j int) int {
	if j < len(input[i])-1 && input[i][j+1] == 0 {
		return recursiveSearchRight(input, i, j+1)
	}
	return j
}

func tiltRight(input [][]int) [][]int {
	fn := func(i int) {
		for j := len(input[i]) - 2; j >= 0; j-- {
			if input[i][j] == 1 {
				moveTo := recursiveSearchRight(input, i, j)
				if moveTo != j {
					input[i][moveTo] = 1
					input[i][j] = 0
				}
			}
		}
	}
	utils.ParalleliseVoid(fn, len(input))
	return input
}

func cycle(input [][]int) [][]int {
	return tiltRight(tiltDown(tiltLeft(tiltUp(input))))
}

func calculateLoad(input [][]int) int {
	weight := 0
	for i, line := range input {
		for _, char := range line {
			if char == 1 {
				weight += len(input) - i
			}
		}
	}
	return weight
}

func SolveDay14Part1(input [][]int) int {
	tilted := tiltUp(input)
	return calculateLoad(tilted)
}

func copyArray(input [][]int) [][]int {
	output := make([][]int, len(input))
	for i, line := range input {
		output[i] = make([]int, len(line))
		copy(output[i], line)
	}
	return output
}

func hashInput(input [][]int) string {
	hash := ""
	for _, line := range input {
		hash += fmt.Sprint(line)
	}
	return hash
}

func SolveDay14Part2(input [][]int) int {
	var cache = make(map[string]int)
	cache[hashInput(input)] = 0
	total := 1000000000
	for n := 0; n < total; n++ {
		input = cycle(input)
		hash := hashInput(input)
		if val, ok := cache[hash]; ok {
			period := n - val
			remaining := total - n
			remaining %= period
			for i := 0; i < remaining-1; i++ {
				input = cycle(input)
			}
			break
		} else {
			cache[hash] = n
		}
	}
	return calculateLoad(input)
}

func Day14(input []string) []string {
	parsed := parseDay14(input)
	part1 := SolveDay14Part1(copyArray(parsed))
	part2 := func() int {
		var input [][]int = parsed
		var cache = make(map[string]int)
		cache[hashInput(input)] = 0
		total := 1000000000
		for n := 0; n < total; n++ {
			input = cycle(input)
			hash := hashInput(input)
			if val, ok := cache[hash]; ok {
				period := n - val
				remaining := total - n
				remaining %= period
				for i := 0; i < remaining-1; i++ {
					input = cycle(input)
				}
				break
			} else {
				cache[hash] = n
			}
		}
		return calculateLoad(input)
	}()
	return []string{fmt.Sprint(part1), fmt.Sprint(part2)}
}
