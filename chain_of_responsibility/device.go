package main

import "fmt"

// Device утройство для передачи данных из источника
type Device struct {
	name string
	next Service // следующий обработчик
}

// реализация получения данных
func (device *Device) execute(data *Data) {
	// если данные уже получены, они передаются следующему звену
	if data.getSource {
		fmt.Printf("Data from device %s has already been got\n", device.name)
		device.next.execute(data)
		return
	}
	// получение данных, полсе которых происходит передача следующему звену
	fmt.Printf("Get data from device %s\n", device.name)
	data.getSource = true
	device.next.execute(data)
}

// передача обработки следующему звену
func (device *Device) setNext(svc Service) {
	device.next = svc
}
