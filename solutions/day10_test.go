/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import "testing"

var testDataDay10 = [][]string{
	{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	},
	{
		"...........",
		".S-------7.",
		".|F-----7|.",
		".||.....||.",
		".||.....||.",
		".|L-7.F-J|.",
		".|..|.|..|.",
		".L--J.L--J.",
		"...........",
	},
	{
		"..........",
		".S------7.",
		".|F----7|.",
		".||OOOO||.",
		".||OOOO||.",
		".|L-7F-J|.",
		".|II||II|.",
		".L--JL--J.",
		"..........",
	},
}

func TestDay10Part1(t *testing.T) {
	input := testDataDay10[0]
	result := Day10(input)[0]
	expected := "8"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestDay10Part2_1(t *testing.T) {
	input := testDataDay10[1]
	result := Day10(input)[1]
	expected := "4"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
func TestDay10Part2_2(t *testing.T) {
	input := testDataDay10[2]
	result := Day10(input)[1]
	expected := "4"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
