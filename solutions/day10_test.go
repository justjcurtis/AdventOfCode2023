/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import "testing"

var testDataDay10 = []string{
	"..F7.",
	".FJ|.",
	"SJ.L7",
	"|F--J",
	"LJ...",
}

func TestDay10Part1(t *testing.T) {
	input := testDataDay10
	result := Day10(input)[0]
	expected := "8"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
