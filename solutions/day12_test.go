/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package solutions

import "testing"

var testDataDay12 = []string{
	"???.### 1,1,3",
	".??..??...?##. 1,1,3",
	"?#?#?#?#?#?#?#? 1,3,1,6",
	"????.#...#... 4,1,1",
	"????.######..#####. 1,6,5",
	"?###???????? 3,2,1",
}

func TestDay12Part1(t *testing.T) {
	input := testDataDay12
	expected := "21"
	actual := Day12(input)[0]
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
