/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import "testing"

var testDataDay9 = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
}

func TestDay9(t *testing.T) {
	input := testDataDay9
	expected := "114"
	actual := Day9(input)[0]
	if actual != expected {
		t.Errorf("TestDay9Part1() returned %s, expected %s", actual, expected)
	}
}

func TestDay9Part2(t *testing.T) {
	input := testDataDay9
	expected := "2"
	actual := Day9(input)[1]
	if actual != expected {
		t.Errorf("TestDay9Part2() returned %s, expected %s", actual, expected)
	}
}
