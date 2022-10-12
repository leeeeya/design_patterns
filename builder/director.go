package main

// Director структура для сборки объекта с интерфейсом сборщика
type Director struct {
	builder Builder
}

// инициализация новой структуры с конкретным сборщиком
func newDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

// метод для возможности изменить вид возвращаемого объекта для уже существующей структуры Director
func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

// пошаговое выполнение сборки конкретного вида структуры House
// можно удалять или изменять последовательность шагов
func (d *Director) buildHouse() *House {
	if d.builder == nil {
		return nil
	}
	d.builder.setHouseType()
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
