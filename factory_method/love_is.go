package main

import "fmt"

// структура для конкретного типа объекта
type loveIs struct {
	Gum
}

func (l loveIs) printDetails() {
	fmt.Printf("name: %s, taste: %s\n", l.name, l.taste)
}

// возвращает готовую структуру, удовлетворяющую интерфейсу IGum
func newLoveIs() IGum {
	return &loveIs{
		Gum: Gum{
			name:  LoveIs,
			taste: "pineapple",
		},
	}
}
