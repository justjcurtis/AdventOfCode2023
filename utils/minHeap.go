/*
Copyright © 2023 Jacson Curtis <justjcurtis@gmail.com>
*/
package utils

import "math"

type MinHeap[T any] struct {
	heap       []T
	comparator func(T, T) bool
	isEqual    func(T, T) bool
}

func NewMinHeap[T any](comparator func(T, T) bool, isEqual func(T, T) bool) *MinHeap[T] {
	return &MinHeap[T]{comparator: comparator}
}

func (h *MinHeap[T]) Push(value T) {
	h.heap = append(h.heap, value)
	h.bubbleUp(len(h.heap) - 1)
}

func (h *MinHeap[T]) Pop() T {
	last := len(h.heap) - 1
	h.swap(0, last)
	value := h.heap[last]
	h.heap = h.heap[:last]
	h.bubbleDown(0)
	return value
}

func (h *MinHeap[T]) Peek() T {
	return h.heap[0]
}

func (h *MinHeap[T]) bubbleUp(index int) {
	parent := int(math.Floor(float64(index-1) / float64(2)))
	if parent < 0 || h.comparator(h.heap[parent], h.heap[index]) {
		return
	}
	h.swap(index, parent)
	h.bubbleUp(parent)
}

func (h *MinHeap[T]) bubbleDown(index int) {
	left := index*2 + 1
	right := index*2 + 2
	if left >= len(h.heap) {
		return
	}
	smallest := left
	if right < len(h.heap) && h.comparator(h.heap[right], h.heap[left]) {
		smallest = right
	}
	if h.comparator(h.heap[index], h.heap[smallest]) {
		return
	}
	h.swap(index, smallest)
	h.bubbleDown(smallest)
}

func (h *MinHeap[T]) swap(i int, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *MinHeap[T]) Len() int {
	return len(h.heap)
}

func (h *MinHeap[T]) IsEmpty() bool {
	return h.Len() == 0
}

func (h *MinHeap[T]) Clear() {
	h.heap = []T{}
}

func (h *MinHeap[T]) Contains(value T) bool {
	for _, item := range h.heap {
		if h.isEqual(item, value) {
			return true
		}
	}
	return false
}
