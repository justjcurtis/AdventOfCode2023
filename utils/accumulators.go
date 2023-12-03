/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

func IntAcc(a int, b int) int {
	return a + b
}

func IntPairAcc(a []int, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	if len(a) != len(b) {
		panic("Arrays must be the same length")
	}
	return []int{a[0] + b[0], a[1] + b[1]}
}
