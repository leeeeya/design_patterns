package main

type IglooBuilder struct {
	houseType  string
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

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

func (b *IglooBuilder) getHouse() House {
	return House{
		houseType:  b.houseType,
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
