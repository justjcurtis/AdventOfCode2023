/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import "testing"

var testDataDay16 = []string{
	".|...\\....",
	"|.-.\\.....",
	".....|-...",
	"........|.",
	"..........",
	".........\\",
	"..../.\\\\..",
	".-.-/..|..",
	".|....-|.\\",
	"..//.|....",
}

func TestDay16Part1(t *testing.T) {
	input := testDataDay16
	expected := "46"
	actual := Day16(input)[0]
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestDay16Part2(t *testing.T) {
	input := testDataDay16
	expected := "51"
	actual := Day16(input)[1]
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
