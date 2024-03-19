/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import "testing"

var testDataDay14 = []string{
	"O....#....",
	"O.OO#....#",
	".....##...",
	"OO.#O....O",
	".O.....O#.",
	"O.#..O.#.#",
	"..O..#O..O",
	".......O..",
	"#....###..",
	"#OO..#....",
}

func TestDay14Part1(t *testing.T) {
	input := testDataDay14
	expected := "136"
	actual := Day14(input)[0]
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestDay14Part2(t *testing.T) {
	input := testDataDay14
	expected := "64"
	actual := Day14(input)[1]
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
