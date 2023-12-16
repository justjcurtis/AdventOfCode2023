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

func Day12(input []string) []string {
	ParseDay12(input)
	return []string{"", ""}
}
