/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import "testing"

var testDataDay15 = []string{
	"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
}

func HashTestDay15String(t *testing.T) {
	input := []string{"HASH"}
	expected := "52"
	actual := Day15(input)[0]
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestDay15Part1(t *testing.T) {
	input := testDataDay15
	expected := "1320"
	actual := Day15(input)[0]
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestDay15Part2(t *testing.T) {
	input := testDataDay15
	expected := "145"
	actual := Day15(input)[1]
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
