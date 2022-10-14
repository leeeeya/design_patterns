package main

import "log"

// InsertionSort стратегия пузырьковой сортировки
type InsertionSort struct {
}

// Sort реализация метода сортировки у конкретной стратегии
func (s *InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		log.Println("Too few items in slice")
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}
