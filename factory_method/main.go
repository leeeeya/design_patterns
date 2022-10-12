package main

import "fmt"

func main() {
	hubbaBubba := getGum("Hubba Bubba")
	loveIs := getGum("Love Is")

	if hubbaBubba != nil {
		hubbaBubba.printDetails()
	}

	fmt.Println()

	if loveIs != nil {
		loveIs.printDetails()
	}

	fmt.Println()
	// инициализация несуществующего объекта
	otherGum := getGum("Orbit")
	if otherGum != nil {
		otherGum.printDetails()
	}
}
