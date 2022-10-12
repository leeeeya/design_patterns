package main

// IglooBuilder струтура для определённого вида объекта
type IglooBuilder struct {
	houseType  string
	windowType string
	doorType   string
	floor      int
}

// билдер для конкретного типа дома
func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

// импелементация методов для построения определённого объекта
func (b *IglooBuilder) setHouseType() {
	b.houseType = "Igloo"
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "Ice Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

// getHouse получение готового объекта типа House
func (b *IglooBuilder) getHouse() *House {
	return &House{
		houseType:  b.houseType,
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
