package main

import "fmt"

func main() {
	// инициализация билдеров для разных видов типа House
	woodenBuilder := getBuilder("wooden")
	iglooBuilder := getBuilder("igloo")

	// создание новой струтуры Director с одним из билдеров
	director := newDirector(woodenBuilder)
	// сборка готового объекта типа "wooden"
	woodenHouse := director.buildHouse()

	woodenHouse.Print()

	fmt.Println()
	// конфигурация нового сборщика для другого вида объекта House
	director.setBuilder(iglooBuilder)
	// сборка готового объекта типа "igloo"
	iglooHouse := director.buildHouse()

	iglooHouse.Print()
	fmt.Println()

	// инициализация несуществующего билдера
	otherBuilder := getBuilder("Candy")
	director.setBuilder(otherBuilder)
	candyHouse := director.buildHouse()
	candyHouse.Print()
}
