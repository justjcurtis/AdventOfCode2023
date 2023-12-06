/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import (
	"reflect"
	"testing"
)

var testDataDay6 = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func TestDay6Part1(t *testing.T) {
	input := testDataDay6
	expected := "288"
	actual := Day6(input)[0]
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v to equal %v", actual, expected)
	}
}

func TestDay6Part2(t *testing.T) {
	input := testDataDay6
	expected := "71503"
	actual := Day6(input)[1]
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v to equal %v", actual, expected)
	}
}
