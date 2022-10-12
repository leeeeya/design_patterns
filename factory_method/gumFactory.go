package main

import (
	"log"
)

const (
	HubbaBubba = "Hubba Bubba Gum"
	LoveIs     = "Love Is Gum"
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
	case HubbaBubba:
		return newHubbaBubba()
	case LoveIs:
		return newLoveIs()
	default:
		log.Printf("%s: Wrong type passed!", gumType)
		return nil
	}
}
