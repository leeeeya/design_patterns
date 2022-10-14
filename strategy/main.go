package main

import (
	"fmt"
)

func main() {
	// инициализация данных для сортировки
	data1 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}
	fmt.Print("Unsorted data: ", data1)
	fmt.Println()

	// инициализация контекста
	ctx := new(Context)

	// добавление конкретной стратегии сортировки
	ctx.Algorithm(&BubbleSort{})
	ctx.Sort(data1)

	fmt.Println("Sorted data: ", data1)
	fmt.Println()

	data2 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}
	fmt.Print("Unsorted data: ", data2)
	fmt.Println()

	ctx.Algorithm(&InsertionSort{})
	ctx.Sort(data2)

	fmt.Println("Sorted data: ", data1)
	fmt.Println()
}
