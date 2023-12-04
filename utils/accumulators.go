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
	if len(a) != len(b) || len(a) != 2 {
		panic("IntPairAcc: invalid input")
	}
	return []int{a[0] + b[0], a[1] + b[1]}
}

func Arr2DAcc[T any](a [][]T, b [][]T) [][]T {
	return append(a, b...)
}

func MapAcc[T comparable, U any](a map[T]U, b map[T]U) map[T]U {
	combined := make(map[T]U)
	for k, v := range a {
		combined[k] = v
	}
	for k, v := range b {
		combined[k] = v
	}
	return combined
}
