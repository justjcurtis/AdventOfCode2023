/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"testing"
)

var testDataDay11 = []string{
	"...#......",
	".......#..",
	"#.........",
	"..........",
	"......#...",
	".#........",
	".........#",
	"..........",
	".......#..",
	"#...#.....",
}

func TestDay11Part1(t *testing.T) {
	input := testDataDay11
	expected := "374"
	actual := Day11(input)[0]
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
