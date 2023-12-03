/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"AdventOfCode2023/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func GetNumFromCoords(i int, j int, input []string) (int, int) {
	line := input[j]
	start := i
	for start >= 0 && unicode.IsDigit(rune(line[start])) {
		start--
	}
	start++
	end := i
	for end < len(line) && unicode.IsDigit(rune(line[end])) {
		end++
	}
	end--
	num, _ := strconv.Atoi(line[start : end+1])
	return num, start
}

func GetNumbersFromSymbol(i int, j int, input []string) map[string][]int {
	coords := [][]int{}
	if i > 0 {
		coords = append(coords, []int{i - 1, j})
		if j > 0 {
			coords = append(coords, []int{i - 1, j - 1})
			coords = append(coords, []int{i, j - 1})
		}
		if j < len(input)-1 {
			coords = append(coords, []int{i - 1, j + 1})
			coords = append(coords, []int{i, j + 1})
		}
	}
	if i < len(input[j])-1 {
		coords = append(coords, []int{i + 1, j})
		if j > 0 {
			coords = append(coords, []int{i + 1, j - 1})
		}
		if j < len(input)-1 {
			coords = append(coords, []int{i + 1, j + 1})
		}
	}

	nums := make(map[string][]int)
	for _, coord := range coords {
		x, y := coord[0], coord[1]
		char := rune(input[y][x])
		if unicode.IsDigit(char) {
			num, start := GetNumFromCoords(x, y, input)
			nums[strings.Join([]string{strconv.Itoa(start), strconv.Itoa(y)}, ",")] = []int{num, i, j}
		}
	}

	return nums
}

func GetRelevenNumbersFromLine(j int, input []string, isSymbol func(rune) bool) map[string][]int {
	line := input[j]
	nums := make(map[string][]int)
	for i := 0; i < len(line); i++ {
		char := rune(line[i])
		if char == '.' {
			continue
		}
		if isSymbol(char) {
			curr := GetNumbersFromSymbol(i, j, input)
			nums = utils.MapAcc(nums, curr)
		}
	}
	return nums
}

func GetReleventNumbers(input []string) map[string][]int {
	fn := func(j int) map[string][]int {
		return GetRelevenNumbersFromLine(j, input, isSymbol)
	}
	return utils.Parallelise(utils.MapAcc[string, []int], fn, len(input))
}

func isSymbol(char rune) bool {
	return !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '.'
}

func GetGears(input []string, nums map[string][]int) map[string][]int {
	gears := make(map[string][]int)
	for _, arr := range nums {
		i := arr[1]
		j := arr[2]
		char := rune(input[j][i])
		if char == '*' {
			id := strings.Join([]string{strconv.Itoa(i), strconv.Itoa(j)}, ",")
			gears[id] = append(gears[id], arr[0])
		}
	}
	for key, arr := range gears {
		if len(arr) != 2 {
			delete(gears, key)
		}
	}
	return gears
}

func Day3(input []string) []string {
	start := time.Now()
	nums := GetReleventNumbers(input)
	elapsed := time.Since(start)
	fmt.Printf("Time to get nums: %s\n", elapsed)
	start = time.Now()
	part1 := 0
	for _, arr := range nums {
		part1 += arr[0]
	}
	elapsed = time.Since(start)
	fmt.Printf("Time to get part 1: %s\n", elapsed)

	start = time.Now()
	gears := GetGears(input, nums)
	elapsed = time.Since(start)
	fmt.Printf("Time to get gears: %s\n", elapsed)
	start = time.Now()
	part2 := 0
	for _, arr := range gears {
		part2 += arr[0] * arr[1]
	}
	elapsed = time.Since(start)
	fmt.Printf("Time to get part 2: %s\n\n", elapsed)

	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
