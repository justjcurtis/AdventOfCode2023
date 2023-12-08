/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import "testing"

var testDataDay8 = [][]string{
	{
		"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	},
	{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	},
	{
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	},
}

func TestDay8Part1(t *testing.T) {
	input := testDataDay8[0]
	result := Day8(input)[0]
	expected := "2"
	if result != expected {
		t.Errorf("TestDay8Part1() failed, expected %v, got %v", expected, result)
	}
	input = testDataDay8[1]
	result = Day8(input)[0]
	expected = "6"
	if result != expected {
		t.Errorf("TestDay8Part1() - 2 failed, expected %v, got %v", expected, result)
	}
}

func TestDay8Part2(t *testing.T) {
	input := testDataDay8[2]
	parsed := ParseDay8(input)
	result := SolveDay8Part2(parsed)
	expected := 6
	if result != expected {
		t.Errorf("TestDay8Part2() failed, expected %v, got %v", expected, result)
	}
}
