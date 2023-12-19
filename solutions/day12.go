/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"AdventOfCode2023/utils"
	"strconv"
	"strings"
)

func ParseLineDay12(line string) (string, []int) {
	arr := strings.Split(line, " ")
	numStrs := strings.Split(arr[1], ",")
	nums := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		nums[i], _ = strconv.Atoi(numStr)
	}
	return arr[0], nums
}

func ParseDay12(input []string) ([]string, [][]int) {
	chars := make([]string, len(input))
	nums := make([][]int, len(input))
	fn := func(j int) {
		char, num := ParseLineDay12(input[j])
		chars[j] = char
		nums[j] = num
	}
	utils.ParalleliseVoid(fn, len(input))
	return chars, nums
}

func GetIdDay12(chars string, nums []int) string {
	id := chars
	for _, num := range nums {
		id += "," + strconv.Itoa(num)
	}
	return id
}

var day12Cache = make(map[string]int)

func SolveDay12Line(chars string, nums []int) int {
	id := GetIdDay12(chars, nums)
	if val, ok := day12Cache[id]; ok {
		return val
	}
	if len(chars) == 0 {
		if len(nums) == 0 {
			return 1
		}
		return 0
	}
	if len(nums) == 0 {
		if strings.Contains(chars, "#") {
			return 0
		}
		return 1
	}
	char := rune(chars[0])
	result := 0
	switch char {
	case '#':
		if nums[0] <= len(chars) &&
			!strings.Contains(chars[:nums[0]], ".") &&
			(nums[0] == len(chars) || chars[nums[0]] != '#') {
			nextChars := ""
			if nums[0]+1 < len(chars) {
				nextChars = chars[nums[0]+1:]
			}
			result = SolveDay12Line(nextChars, nums[1:])
			break
		}
		result = 0
		break
	case '.':
		result = SolveDay12Line(chars[1:], nums)
		break
	case '?':
		if nums[0] <= len(chars) &&
			!strings.Contains(chars[:nums[0]], ".") &&
			(nums[0] == len(chars) || chars[nums[0]] != '#') {
			nextChars := ""
			if nums[0]+1 < len(chars) {
				nextChars = chars[nums[0]+1:]
			}
			result += SolveDay12Line(nextChars, nums[1:])
		}
		result += SolveDay12Line(chars[1:], nums)
		break
	}
	day12Cache[id] = result
	return result
}

func SolveDay12Part1(chars []string, nums [][]int) int {
	total := 0
	for i := 0; i < len(chars); i++ {
		total += SolveDay12Line(chars[i], nums[i])
	}
	return total
}

func UnfoldDay12Line(chars string, nums []int) (string, []int) {
	resultChars := strings.Join([]string{
		chars, chars, chars, chars, chars,
	}, "?")
	resultNums := make([]int, len(nums)*5)
	for i := 0; i < 5; i++ {
		for j, num := range nums {
			resultNums[i*len(nums)+j] = num
		}
	}
	return resultChars, resultNums
}

func SolveDay12Part2Line(chars string, nums []int) int {
	unfoldedChars, unfoldedNums := UnfoldDay12Line(chars, nums)
	return SolveDay12Line(unfoldedChars, unfoldedNums)
}

func SolveDay12Part2(chars []string, nums [][]int) int {
	total := 0
	for i := 0; i < len(chars); i++ {
		total += SolveDay12Part2Line(chars[i], nums[i])
	}
	return total
}

func Day12(input []string) []string {
	chars, nums := ParseDay12(input)
	part1 := SolveDay12Part1(chars, nums)
	part2 := SolveDay12Part2(chars, nums)
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
