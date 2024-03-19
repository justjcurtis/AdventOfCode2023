package solutions

import (
	"AdventOfCode2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseDay15(input []string) []string {
	result := []string{}
	buffer := ""
	for _, s := range input {
		for _, c := range s {
			if c == ',' {
				result = append(result, buffer)
				buffer = ""
			} else {
				buffer += string(c)
			}
		}
	}
	result = append(result, buffer)
	return result
}

func hashDay15String(input string) int {
	hash := 0
	for _, c := range input {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return hash
}

func SolveDay15Part1(input []string) int {
	fn := func(i int) int {
		return hashDay15String(input[i])
	}
	return utils.Parallelise(utils.IntAcc, fn, len(input))
}

type instruction struct {
	operation   string
	label       string
	box         int
	focalLength int
}

type lense struct {
	label       string
	focalLength int
}

func getLenseIndex(box []lense, label string) int {
	return slices.IndexFunc(box, func(l lense) bool {
		return l.label == label
	})
}

func removeLense(box []lense, label string) []lense {
	index := getLenseIndex(box, label)
	if index == -1 {
		return box
	}
	return append(box[:index], box[index+1:]...)
}

func addLense(box []lense, label string, focalLength int) []lense {
	index := getLenseIndex(box, label)
	if index == -1 {
		box = append(box, lense{label: label, focalLength: focalLength})
	} else {
		box[index].focalLength = focalLength
	}
	return box
}

func SolveDay15Part2(input []string) int {
	boxes := make([][]lense, 256)
	instructions := make([]instruction, len(input))

	fn := func(i int) {
		raw := input[i]
		operation := "="
		label := raw[:len(raw)-1]
		if raw[len(raw)-1] == '-' {
			operation = "-"
		}
		focalLength := -1
		if operation == "=" {
			parts := strings.Split(raw, "=")
			label = parts[0]
			focalLength, _ = strconv.Atoi(parts[1])
		}
		instructions[i] = instruction{
			operation:   operation,
			label:       label,
			box:         hashDay15String(label),
			focalLength: focalLength,
		}
	}

	utils.ParalleliseVoid(fn, len(input))
	for _, i := range instructions {
		if i.operation == "=" {
			boxes[i.box] = addLense(boxes[i.box], i.label, i.focalLength)
		} else {
			boxes[i.box] = removeLense(boxes[i.box], i.label)
		}
	}

	result := 0
	for i, b := range boxes {
		for j, l := range b {
			result += (i + 1) * (j + 1) * l.focalLength
		}
	}
	return result
}

func Day15(input []string) []string {
	parsed := parseDay15(input)
	part1 := SolveDay15Part1(parsed)
	part2 := SolveDay15Part2(parsed)
	return []string{fmt.Sprint(part1), fmt.Sprint(part2)}
}
