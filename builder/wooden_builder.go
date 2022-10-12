package main

// WoodenBuilder струтура для определённого вида объекта
type WoodenBuilder struct {
	houseType  string
	windowType string
	doorType   string
	floor      int
}

// билдер для конкретного типа дома
func newWoodenBuilder() *WoodenBuilder {
	return &WoodenBuilder{}
}

// импелементация методов для построения определённого объекта
func (b *WoodenBuilder) setHouseType() {
	b.houseType = Wooden
}

func (b *WoodenBuilder) setWindowType() {
	b.windowType = "Plastic Window"
}

func (b *WoodenBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *WoodenBuilder) setNumFloor() {
	b.floor = 2
}

// getHouse получение готового объекта типа House
func (b *WoodenBuilder) getHouse() *House {
	return &House{
		houseType:  b.houseType,
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
