package main

import (
	"fmt"
)

// структура для конкретного типа объекта
type hubbaBubba struct {
	Gum
	numInPack int
}

func (h hubbaBubba) printDetails() {
	fmt.Printf("name: %s, taste: %s, number in package: %d\n", h.name, h.taste, h.numInPack)
}

// возвращает готовую структуру, удовлетворяющую интерфейсу IGum
func newHubbaBubba() IGum {
	return &hubbaBubba{
		Gum: Gum{
			name:  "Hubba Bubba Gum",
			taste: "strawberry",
		},
		numInPack: 5,
	}
}
