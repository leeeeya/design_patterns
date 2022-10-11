package main

// Builder определяет методы для сборки структуры House
type Builder interface {
	setHouseType()
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House // возвращает готовый объект
}

// getBuilder возвращает билдера в зависимости от вида объекта
func getBuilder(builderType string) Builder {
	switch builderType {
	case "wooden":
		return newWoodenBuilder()
	case "igloo":
		return newIglooBuilder()
	default:
		return nil
	}
}
