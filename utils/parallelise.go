/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import "runtime"

func Parallelise[T any](acc func(T, T) T, fn func(int) T, maxLength int) T {
	var results T
	workerCount := runtime.NumCPU() - 1
	ch := make(chan T)
	for i := 0; i < workerCount; i++ {
		start := maxLength / workerCount * i
		end := maxLength / workerCount * (i + 1)
		if i == workerCount-1 {
			end = maxLength
		}
		go func(i int) {
			var result T
			for j := start; j < end; j++ {
				result = acc(result, fn(j))
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
