package main

import "fmt"

func main() {
	hubbaBubba := getGum(HubbaBubba)
	loveIs := getGum(LoveIs)

	if hubbaBubba != nil {
		hubbaBubba.printDetails()
	}

	fmt.Println()

	if loveIs != nil {
		loveIs.printDetails()
	}

	fmt.Println()
	// инициализация несуществующего объекта
	otherGum := getGum("Orbit") // nil
	if otherGum != nil {
		otherGum.printDetails()
	}
}
