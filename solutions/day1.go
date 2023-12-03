/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"AdventOfCode2023/utils"
	"runtime"
	"strconv"
	"unicode"
)

var WORD_NUMS = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func SolutionPerLine(line string, part int) int {
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

func Solve(input []string, part int, c chan<- int) {
	result := 0
	ch := make(chan int)
	workerCount := runtime.NumCPU() / 2
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			result := 0
			for j := (len(input) / workerCount * i); j < (len(input) / workerCount * (i + 1)); j++ {
				line := input[j]
				result += SolutionPerLine(line, part)
			}
			ch <- result
		}(i)
	}
	for i := 0; i < workerCount; i++ {
		result += <-ch
	}
	c <- result
}

func Day1() {
	input := utils.GetInput(1)

	part1 := 0
	part2 := 0

	a := make(chan int)
	go Solve(input, 1, a)

	b := make(chan int)
	go Solve(input, 2, b)

	for i := 0; i < 2; i++ {
		select {
		case num := <-a:
			part1 += num
		case num := <-b:
			part2 += num
		}
	}

	println("=------ Day 1 ------=")
	println("Part 1: " + strconv.Itoa(part1))
	println("Part 2: " + strconv.Itoa(part2))
	println("=-------------------=")
}
