/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package main

import (
	"AdventOfCode2023/solutions"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	solutions.Day1()
	elapsed := time.Since(start)
	fmt.Printf("Day 1 took %s\n", elapsed)
	println()

	start = time.Now()
	solutions.Day2()
	elapsed = time.Since(start)
	fmt.Printf("Day 2 took %s\n", elapsed)
	println()

}
