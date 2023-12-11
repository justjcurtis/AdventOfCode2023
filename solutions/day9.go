/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"AdventOfCode2023/utils"
	"strconv"
	"strings"
)

func ParseDay9Line(line string) []int {
	arr := strings.Split(line, " ")
	nums := make([]int, len(arr))
	for i, v := range arr {
		nums[i], _ = strconv.Atoi(v)
	}

	return nums
}

func SolveDay9Line(nums []int, z int) []int {
	rows := [][]int{}
	rows = append(rows, nums)
	for j := 1; j < len(nums); j++ {
		allZero := true
		row := []int{}
		prevRow := rows[j-1]
		for i := 0; i < len(prevRow)-1; i++ {
			num := prevRow[i+1] - prevRow[i]
			if num != 0 {
				allZero = false
			}
			row = append(row, num)
		}
		if allZero {
			break
		}
		rows = append(rows, row)
	}

	part1 := 0
	part2 := 0
	for j := len(rows) - 1; j >= 0; j-- {
		part1 += rows[j][len(rows[j])-1]
		part2 = rows[j][0] - part2
	}
	return []int{part1, part2}
}

func SolveDay9(input []string) []int {
	fn := func(j int) []int {
		nums := ParseDay9Line(input[j])
		total := SolveDay9Line(nums, j)
		return total
	}
	return utils.Parallelise(utils.IntPairAcc, fn, len(input))
}

func Day9(input []string) []string {
	arr := SolveDay9(input)
	part1 := arr[0]
	part2 := arr[1]
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
