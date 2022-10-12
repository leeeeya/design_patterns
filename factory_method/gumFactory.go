package main

import (
	"log"
)

// IGum основной интерфейс поведения фабрики
type IGum interface {
	getName() string
	getTaste() string
	setName(name string)
	setTaste(taste string)
	printDetails()
}

// фабричный метод - создаёт новый интерфейс в зависимости от вида (входных данных)
func getGum(gumType string) IGum {
	switch gumType {
	case "Hubba Bubba":
		return newHubbaBubba()
	case "Love Is":
		return newLoveIs()
	default:
		log.Printf("%s: Wrong type passed", gumType)
		return nil
	}
}
