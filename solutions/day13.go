/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"strconv"
)

func ParseDay13Input(input []string) [][]string {
	patterns := [][]string{{}}
	for _, line := range input {
		if line == "" {
			patterns = append(patterns, []string{})
			continue
		}
		patterns[len(patterns)-1] = append(patterns[len(patterns)-1], line)
	}
	return patterns
}

func RotatePattern(pattern []string) []string {
	rotated := make([]string, len(pattern[0]))
	for i := 0; i < len(pattern[0]); i++ {
		for j := len(pattern) - 1; j >= 0; j-- {
			rotated[i] += string(pattern[j][i])
		}
	}
	return rotated
}

func StringCompare(a, b string, allowDiff bool) (bool, bool) {
	if !allowDiff || a == b {
		return a == b, false
	}
	hasDiff := false
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if hasDiff {
				return false, false
			}
			hasDiff = true
		}
	}
	return true, true
}

func GetReflectedRow(pattern []string, allowDiff bool, skip int) int {
	for i := 0; i < len(pattern)-1; i++ {
		usedDiff := false
		found := true
		for j := 0; j <= i && j+i+1 < len(pattern); j++ {
			equal, diff := StringCompare(pattern[i-j], pattern[i+j+1], allowDiff && !usedDiff)
			if !equal {
				found = false
				break
			}
			usedDiff = usedDiff || diff
		}
		if found && i != skip {
			return i
		}
	}
	return -1
}

func getEmptySkips(n int) [][]int {
	emptySkips := make([][]int, n)
	for i := range emptySkips {
		emptySkips[i] = []int{-1, -1}
	}
	return emptySkips
}

func SolveDay13(patterns [][]string, allowDiff bool, prevSkips [][]int) (int, [][]int) {
	horizontal := 0
	vertical := 0
	skips := getEmptySkips(len(patterns))
	for i, pattern := range patterns {
		upperRows := GetReflectedRow(pattern, allowDiff, prevSkips[i][0]) + 1
		if upperRows == 0 {
			leftRows := GetReflectedRow(RotatePattern(pattern), allowDiff, prevSkips[i][1]) + 1
			skips[i][1] = leftRows - 1
			vertical += leftRows
			continue
		}
		horizontal += upperRows * 100
		skips[i][0] = upperRows - 1
	}
	return (horizontal + vertical), skips
}

func Day13(input []string) []string {
	patterns := ParseDay13Input(input)
	emptySkips := getEmptySkips(len(patterns))
	part1, skips := SolveDay13(patterns, false, emptySkips)
	part2, _ := SolveDay13(patterns, true, skips)
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
