/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import (
	"runtime"
	"sync"
)

func Parallelise[T any](acc func(T, T) T, fn func(int) T, maxLength int) T {
	var results T
	workerCount := runtime.NumCPU() - 1
	ch := make(chan T, workerCount)
	wg := sync.WaitGroup{}
	for i := 0; i < workerCount; i++ {
		start := maxLength / workerCount * i
		end := maxLength / workerCount * (i + 1)
		if i == workerCount-1 {
			end = maxLength
		}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var result T
			for j := start; j < end; j++ {
				result = acc(result, fn(j))
			}
			ch <- result
		}(i)
	}
	wg.Wait()
	for i := 0; i < workerCount; i++ {
		results = acc(results, <-ch)
	}
	return results
}
