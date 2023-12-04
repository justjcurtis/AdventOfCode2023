/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"AdventOfCode2023/utils"
	"strconv"
	"unicode"
)

func GetNumFromCoords(i int, j int, input []string) int {
	line := input[j]
	start := i - 1
	for start >= 0 && unicode.IsDigit(rune(line[start])) {
		start--
	}
	start++
	end := i + 1
	for end < len(line) && unicode.IsDigit(rune(line[end])) {
		end++
	}
	end--
	num, _ := strconv.Atoi(line[start : end+1])
	return num
}

func GetNumbersFromSymbol(i int, j int, input []string) [][]int {
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

	numMap := make(map[int]bool)
	nums := [][]int{}
	for _, coord := range coords {
		x, y := coord[0], coord[1]
		char := rune(input[y][x])
		if unicode.IsDigit(char) {
			num := GetNumFromCoords(x, y, input)
			if numMap[num] {
				continue
			}
			numMap[num] = true
			nums = append(nums, []int{num, i, j})
		}
	}

	return nums
}

func GetRelevenNumbersFromLine(j int, input []string) [][]int {
	line := input[j]
	nums := [][]int{}
	for i := 0; i < len(line); i++ {
		char := rune(line[i])
		if char == '.' || unicode.IsDigit(char) {
			continue
		}
		curr := GetNumbersFromSymbol(i, j, input)
		nums = append(nums, curr...)
	}
	return nums
}

func GetReleventNumbers(input []string) [][]int {
	fn := func(j int) [][]int {
		return GetRelevenNumbersFromLine(j, input)
	}
	return utils.Parallelise(utils.Arr2DAcc[int], fn, len(input))
}

func GetGears(input []string, nums [][]int) [][]int {
	w := len(input[0])
	gears := make([][]int, len(input)*w)
	for _, arr := range nums {
		i, j := arr[1], arr[2]
		char := rune(input[j][i])
		if char == '*' {
			id := utils.TwoDToOneD(i, j, w)
			gears[id] = append(gears[id], arr[0])
		}
	}
	result := [][]int{}
	for i := 0; i < len(gears); i++ {
		if len(gears[i]) != 2 {
			continue
		}
		result = append(result, gears[i])
	}
	return result
}

func Day3(input []string) []string {
	nums := GetReleventNumbers(input)
	part1 := 0
	for _, arr := range nums {
		part1 += arr[0]
	}

	gears := GetGears(input, nums)
	part2 := 0
	for _, arr := range gears {
		if len(arr) != 2 {
			continue
		}
		part2 += arr[0] * arr[1]
	}

	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
