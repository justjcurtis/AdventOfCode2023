/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"AdventOfCode2023/utils"
	"strconv"
	"unicode"
)

var WORD_NUMS = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func SolveDay1Line(line string, part int) int {
	numString := ""
	for i := 0; i < len(line); i++ {
		char := rune(line[i])
		if unicode.IsDigit(char) {
			numString += string(char)
			break
		}
		foundWord := false
		if part > 1 {
			for j, word := range WORD_NUMS {
				if len(line[i:]) < len(word) {
					continue
				}
				if line[i:i+len(word)] == word {
					foundWord = true
					numString += strconv.Itoa(j + 1)
					break
				}
			}
			if foundWord {
				break
			}
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		char := rune(line[i])
		if unicode.IsDigit(char) {
			numString += string(char)
			break
		}
		foundWord := false
		if part > 1 {
			for j, word := range WORD_NUMS {
				if len(line[i:]) < len(word) {
					continue
				}
				if line[i:i+len(word)] == word {
					foundWord = true
					numString += strconv.Itoa(j + 1)
					break
				}
			}
			if foundWord {
				break
			}
		}
	}
	if numString == "" {
		numString = "0"
	}
	if len(numString) == 1 {
		numString += numString
	}
	num, err := strconv.Atoi(numString)
	if err != nil {
		num = 0
	}
	return num
}

func SolveDay1(input []string) []int {
	fn := func(j int, input []string) []int {
		line := input[j]
		return []int{SolveDay1Line(line, 1), SolveDay1Line(line, 2)}
	}
	return utils.Parallelise(utils.IntPairAcc, fn, input)
}

func Day1(input []string) []string {
	results := SolveDay1(input)
	return []string{strconv.Itoa(results[0]), strconv.Itoa(results[1])}
}
