/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import "strconv"

func PrintResults(day int, results []string) {
	println("=------ Day " + strconv.Itoa(day) + " ------=")
	println("Part 1: " + results[0])
	println("Part 2: " + results[1])
	println("=-------------------=")
}
