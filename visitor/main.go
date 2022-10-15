package main

import "fmt"

func main() {

	// инициализация множества элементов, которые надо посетить
	city := new(City)

	city.Add(&SushiBar{})
	city.Add(&Pizzeria{})
	city.Add(&BurgerBar{})

	// посещение всех элементов
	result := city.Accept(&People{name: "Masha"})

	fmt.Println(result)
}
