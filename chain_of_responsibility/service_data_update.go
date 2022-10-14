package main

import "fmt"

// UpdateDataService сервис для обновления данных
type UpdateDataService struct {
	name string
	next Service // следующий обработчик
}

// реализация обработки данных
func (upd *UpdateDataService) execute(data *Data) {
	// если данные уже обработаны, они передаются следующему сервису
	if data.updSource {
		fmt.Printf("Data in service %s is already updated\n", upd.name)
		upd.next.execute(data)
		return
	}
	// обработка данных, полсе которых происходит передача следующему сервису
	fmt.Printf("Update data in service %s\n", upd.name)
	data.updSource = true
	upd.next.execute(data)
}

// передача обработки следующему звену
func (upd *UpdateDataService) setNext(svc Service) {
	upd.next = svc
}
