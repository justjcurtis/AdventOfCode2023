/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package main

import (
	"AdventOfCode2023/solutions"
	"AdventOfCode2023/utils"
	"fmt"
	"time"
)

type solution struct {
	day int
	fn  func([]string) []string
}

var SOLUTIONS = []solution{
	{1, solutions.Day1},
	{2, solutions.Day2},
}

func main() {
	var totalTime time.Duration
	for _, solution := range SOLUTIONS {
		input := utils.GetInput(solution.day)
		start := time.Now()
		results := solution.fn(input)
		elapsed := time.Since(start)
		totalTime += elapsed
		utils.PrintResults(solution.day, results)
		fmt.Printf("Day %d took %s\n", solution.day, elapsed)
		println()
	}

	println("=------ Total ------=")
	fmt.Printf("Total time: %s\n", totalTime)
	println("=-------------------=")
}
