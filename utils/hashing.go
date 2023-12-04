/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

func SzudzikPairing(x int, y int) int {
	if x < y {
		return y*y + x
	}
	return x*x + x + y
}
