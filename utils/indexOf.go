/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

func IndexOf(arr []int, val int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}
