/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import "testing"

var testDataDay2 = [][]string{
	{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	},
}

func TestDay2Part1(t *testing.T) {
	input := testDataDay2[0]
	expected := "8"
	actual := Day2(input)[0]
	if actual != expected {
		t.Errorf("Expected %s but was %s", expected, actual)
	}
}

func TestDay2Part2(t *testing.T) {
	input := testDataDay2[0]
	expected := "2286"
	actual := Day2(input)[1]
	if actual != expected {
		t.Errorf("Expected %s but was %s", expected, actual)
	}
}
