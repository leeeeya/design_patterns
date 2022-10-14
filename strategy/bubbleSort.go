package main

import "log"

// BubbleSort стратегия пузырьковой сортировки
type BubbleSort struct {
}

// Sort реализация метода сортировки у конкретной стратегии
func (s *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		log.Println("Too few items in slice")
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}
