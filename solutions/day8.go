/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"AdventOfCode2023/utils"
	"strconv"

	"github.com/patrickmn/go-cache"
)

func ParseDay8(input []string) ([]int, *cache.Cache) {
	instructions := make([]int, len(input[0]))
	for i, c := range input[0] {
		if c == 'R' {
			instructions[i] = 1
			continue
		}
		instructions[i] = 0
	}
	var camelMap = cache.New(cache.NoExpiration, cache.NoExpiration)
	fn := func(i int) {
		line := input[i+2]
		key := line[:3]
		a := line[7:10]
		b := line[12:15]
		camelMap.Add(key, []string{a, b}, cache.NoExpiration)
	}
	utils.ParalleliseVoid(fn, len(input)-2)

	return instructions, camelMap
}

func SolveDay8Part1(instructions []int, camelMap *cache.Cache) int {
	stepCount := 0
	position := "AAA"
	for true {
		for i := 0; i < len(instructions); i++ {
			instruction := instructions[i]
			choices, found := camelMap.Get(position)
			if !found {
				return stepCount
			}
			nextPosition := choices.([]string)[instruction]
			if nextPosition == "ZZZ" {
				return stepCount + 1
			}
			position = nextPosition
			stepCount++
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

func GetLCM(nums []int) int {
	lcm := nums[0]
	for i := 1; i < len(nums); i++ {
		lcm = lcm * nums[i] / GetHCF(lcm, nums[i])
	}
	return lcm
}

func SolveDay8Part2(instructions []int, camelMap *cache.Cache) int {
	starts := []string{}
	for key, _ := range camelMap.Items() {
		if key[2] == 'A' {
			starts = append(starts, key)
		}
	}
	result := make(chan int, len(starts))
	worker := func(start string, result chan<- int) {
		stepCount := 0
		position := start
		history := map[string]int{}
		for true {
			if position[2] == 'Z' {
				if val, found := history[position]; found {
					result <- stepCount - val
					break
				}
				history[position] = stepCount
			}
			instruction := instructions[stepCount%len(instructions)]
			choices, _ := camelMap.Get(position)
			nextPosition := choices.([]string)[instruction]
			stepCount++
			position = nextPosition
		}
	}
	for _, start := range starts {
		go worker(start, result)
	}

	periods := []int{}
	for i := 0; i < len(starts); i++ {
		periods = append(periods, <-result)
	}
	return GetLCM(periods)
}

func Day8(input []string) []string {
	instruction, camelMap := ParseDay8(input)
	part1 := SolveDay8Part1(instruction, camelMap)
	part2 := SolveDay8Part2(instruction, camelMap)
	return []string{strconv.Itoa(part1), strconv.Itoa(part2)}
}
