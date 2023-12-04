/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import "testing"

var testDataDay3 = [][]string{
	{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	},
}

func TestDay3Part1(t *testing.T) {
	input := testDataDay3[0]
	expected := "4361"
	actual := Day3(input)[0]
	if actual != expected {
		t.Errorf("Expected %s but was %s", expected, actual)
	}
}

func TestDay3Part2(t *testing.T) {
	input := testDataDay3[0]
	expected := "467835"
	actual := Day3(input)[1]
	if actual != expected {
		t.Errorf("Expected %s but was %s", expected, actual)
	}
}
