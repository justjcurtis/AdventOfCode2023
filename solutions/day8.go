/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"strconv"
)

type day8Parsed struct {
	start        int
	instructions []int
	names        []string
	choices      [][]int
	isEnd        [][]bool
}

func ParseDay8(input []string) day8Parsed {
	instructions := make([]int, len(input[0]))
	for i, c := range input[0] {
		if c == 'R' {
			instructions[i] = 1
			continue
		}
		instructions[i] = 0
	}
	indexMap := make(map[string]int)
	names := make([]string, len(input)-2)
	choices := make([][]int, len(input)-2)
	isEnd := make([][]bool, len(input)-2)
	start := 0
	for i := 0; i < len(input)-2; i++ {
		line := input[i+2]
		key := line[:3]
		if key == "AAA" {
			start = i
		}
		isEnd[i] = []bool{key == "ZZZ", key[2] == 'Z'}
		indexMap[key] = i
		names[i] = key
	}
	for i := 0; i < len(input)-2; i++ {
		line := input[i+2]
		a := line[7:10]
		b := line[12:15]
		choices[i] = []int{indexMap[a], indexMap[b]}
	}
	return day8Parsed{start, instructions, names, choices, isEnd}
}

func SolveDay8Part1(parsed day8Parsed) int {
	stepCount := 0
	position := parsed.start
	instructionIndex := 0
	for true {
		instruction := parsed.instructions[instructionIndex]
		choices := parsed.choices[position]
		nextPosition := choices[instruction]
		if parsed.isEnd[position][0] {
			return stepCount
		}
		position = nextPosition
		stepCount++
		instructionIndex++
		if instructionIndex == len(parsed.instructions) {
			instructionIndex = 0
		}
	}
	return -1
}

func GetHCF(a int, b int) int {
	if b == 0 {
		return a
	}
	return GetHCF(b, a%b)
}

func GetLCM(a int, b int) int {
	return a * b / GetHCF(a, b)
}

func SolveDay8Part2(parsed day8Parsed) int {
	starts := []int{}
	for i, name := range parsed.names {
		if name[2] == 'A' {
			starts = append(starts, i)
		}
	}
	result := make(chan int, len(starts))
	worker := func(start int, result chan<- int) {
		stepCount := 0
		instructionIndex := 0
		position := start
		for true {
			if parsed.isEnd[position][1] {
				result <- stepCount
				return
			}
			instruction := parsed.instructions[instructionIndex]
			choices := parsed.choices[position]
			nextPosition := choices[instruction]
			instructionIndex++
			if instructionIndex == len(parsed.instructions) {
				instructionIndex = 0
			}
			stepCount++
			position = nextPosition
		}
	}
	for _, start := range starts {
		go worker(start, result)
	}

	lcm := 1
	for i := 0; i < len(starts); i++ {
		lcm = GetLCM(lcm, <-result)
	}
	return lcm
}

func Day8(input []string) []string {
	parsed := ParseDay8(input)
	part1 := SolveDay8Part1(parsed)
	part2 := SolveDay8Part2(parsed)
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
