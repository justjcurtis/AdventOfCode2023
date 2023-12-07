/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import "testing"

var testDataDay7 = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func TestDay7Part1(t *testing.T) {
	input := testDataDay7
	expected := "6440"
	actual := Day7(input)[0]
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDay7Part2(t *testing.T) {
	input := testDataDay7
	expected := "5905"
	actual := Day7(input)[1]
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
