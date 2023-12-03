/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

func CombineMaps[T comparable, U any](a map[T]U, b map[T]U) map[T]U {
	combined := make(map[T]U)
	for k, v := range a {
		combined[k] = v
	}
	for k, v := range b {
		combined[k] = v
	}
	return combined
}
