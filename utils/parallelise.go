/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import "runtime"

func Parallelise[T any, U any](acc func(U, U) U, fn func(T) U, input []T) U {
	var results U
	workerCount := runtime.NumCPU()
	ch := make(chan U)
	for i := 0; i < workerCount; i++ {
		start := len(input) / workerCount * i
		end := len(input) / workerCount * (i + 1)
		if i == workerCount-1 {
			end = len(input)
		}
		go func(i int) {
			var result U
			for j := start; j < end; j++ {
				result = acc(result, fn(input[j]))
			}
			ch <- result
		}(i)
	}
	for i := 0; i < workerCount; i++ {
		current := <-ch
		results = acc(results, current)
	}
	return results
}
