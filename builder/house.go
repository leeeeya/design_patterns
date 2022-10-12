package main

import "fmt"

// House структура, которую требуется создать
type House struct {
	houseType  string
	windowType string
	doorType   string
	floor      int
}

// Print выводит информацию о готовом объекте
func (h *House) Print() {
	if h == nil {
		return
	}
	fmt.Printf("House type: %s\n", h.houseType)
	fmt.Printf("Door type: %s\n", h.doorType)
	fmt.Printf("Window type: %s\n", h.windowType)
	fmt.Printf("Number of floors: %d\n", h.floor)
}
