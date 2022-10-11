package main

import "fmt"

func main() {
	woodenBuilder := getBuilder("wooden")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(woodenBuilder)
	woodenHouse := director.buildHouse()

	woodenHouse.Print()

	fmt.Println()
	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	iglooHouse.Print()

}
