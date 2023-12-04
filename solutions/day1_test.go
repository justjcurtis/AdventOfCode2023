/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"testing"
)

var testDataDay1 = [][]string{
	{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	},
	{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	},
}

func TestDay1Part1(t *testing.T) {
	input := testDataDay1[0]
	expected := "142"
	actual := Day1(input)[0]
	if actual != expected {
		t.Errorf("Expected %s but was %s", expected, actual)
	}
}

func TestDay1Part2(t *testing.T) {
	input := testDataDay1[1]
	expected := "281"
	actual := Day1(input)[1]
	if actual != expected {
		t.Errorf("Expected %s but was %s", expected, actual)
	}
}
