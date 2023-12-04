/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package main

import (
	"AdventOfCode2023/solutions"
	"AdventOfCode2023/utils"
	"flag"
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
	{3, solutions.Day3},
}

func main() {
	runCount := flag.Int("n", 1, "Number of times to run each solution")
	flag.Parse()

	var totalTime time.Duration
	for _, solution := range SOLUTIONS {
		input := utils.GetInput(solution.day)
		start := time.Now()
		for i := 0; i < *runCount-1; i++ {
			solution.fn(input)
		}
		results := solution.fn(input)
		elapsed := time.Since(start)
		totalTime += elapsed
		utils.PrintResults(solution.day, results)
		fmt.Printf("Day %d took %s\n", solution.day, elapsed/time.Duration(*runCount))
		println()
	}

	println("=------ Total ------=")
	fmt.Printf("Total time: %s\n", totalTime/time.Duration(*runCount))
	println("=-------------------=")
}
