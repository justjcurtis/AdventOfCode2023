/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import "testing"

var testDataDay13 = []string{
	"#.##..##.",
	"..#.##.#.",
	"##......#",
	"##......#",
	"..#.##.#.",
	"..##..##.",
	"#.#.##.#.",
	"",
	"#...##..#",
	"#....#..#",
	"..##..###",
	"#####.##.",
	"#####.##.",
	"..##..###",
	"#....#..#",
}

func TestDay13Part1(t *testing.T) {
	input := testDataDay13
	expected := "405"
	actual := Day13(input)[0]
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestDay13Part2(t *testing.T) {
	input := testDataDay13
	expected := "400"
	actual := Day13(input)[1]
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
