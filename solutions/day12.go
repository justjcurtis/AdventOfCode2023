/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"AdventOfCode2023/utils"
	"strconv"
	"strings"
)

func ParseLineDay12(line string) (string, []int, int) {
	arr := strings.Split(line, " ")
	wilds := 0
	for _, char := range arr[0] {
		if char == '?' {
			wilds++
		}
	}
	numStrs := strings.Split(arr[1], ",")
	nums := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		nums[i], _ = strconv.Atoi(numStr)
	}
	return arr[0], nums, wilds
}

func ParseDay12(input []string) ([]string, [][]int, []int) {
	chars := make([]string, len(input))
	nums := make([][]int, len(input))
	wilds := make([]int, len(input))
	fn := func(j int) {
		char, num, wild := ParseLineDay12(input[j])
		chars[j] = char
		nums[j] = num
		wilds[j] = wild
	}
	utils.ParalleliseVoid(fn, len(input))
	return chars, nums, wilds
}

func SolveDay12Line(
	wilds int, chars string, nums []int,
	lastChar rune, index int, groupIndex int, groups []int) int {
	if index == len(chars) {
		if len(groups) < len(nums) {
			return 0
		}
		for i := groupIndex; i < len(nums); i++ {
			if groups[i] != nums[i] {
				return 0
			}
		}
		return 1
	}
	if len(groups) < groupIndex+1 {
		groups = append(groups, 0)
	}
	if groupIndex >= len(nums) {
		for i := index; i < len(chars); i++ {
			if rune(chars[i]) == '#' {
				return 0
			}
		}
		return 1
	}
	result := 0
	switch rune(chars[index]) {
	case '#':
		if groups[groupIndex] == nums[groupIndex] {
			return 0
		}
		nextGroups := make([]int, len(groups))
		copy(nextGroups, groups)
		nextGroups[groupIndex]++
		return SolveDay12Line(wilds, chars, nums, '#', index+1, groupIndex, nextGroups)
	case '.':
		nextGroupIndex := groupIndex
		if index > 0 && lastChar == '#' {
			if groups[groupIndex] < nums[groupIndex] {
				return 0
			}
			nextGroupIndex++
		}
		nextGroups := make([]int, len(groups))
		copy(nextGroups, groups)
		return SolveDay12Line(wilds, chars, nums, '.', index+1, nextGroupIndex, nextGroups)
	case '?':
		if groups[groupIndex] < nums[groupIndex] {
			nextGroups := make([]int, len(groups))
			copy(nextGroups, groups)
			nextGroups[groupIndex]++
			result = SolveDay12Line(wilds-1, chars, nums, '#', index+1, groupIndex, nextGroups)
		}
		nextGroupIndex := groupIndex
		if index > 0 && lastChar == '#' {
			if groups[groupIndex] < nums[groupIndex] {
				return result
			}
			nextGroupIndex++
		}
		otherNextGroups := make([]int, len(groups))
		copy(otherNextGroups, groups)
		result += SolveDay12Line(wilds-1, chars, nums, '.', index+1, nextGroupIndex, otherNextGroups)
		return result
	}
	return -1
}

func SolveDay12Part1(chars []string, nums [][]int, wilds []int) int {
	fn := func(j int) int {
		return SolveDay12Line(wilds[j], chars[j], nums[j], '.', 0, 0, []int{})
	}
	return utils.Parallelise(utils.IntAcc, fn, len(chars))
}

func Day12(input []string) []string {
	chars, nums, wilds := ParseDay12(input)
	part1 := SolveDay12Part1(chars, nums, wilds)
	return []string{strconv.Itoa(part1)}
}
