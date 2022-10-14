package main

import "fmt"

func main() {
	// инициализация обработчиков
	device := &Device{name: "Device-1"}
	updService := &UpdateDataService{name: "Update-1"}
	storageService := &DataStorageService{}

	// установление порядка выполнения обработчиков
	device.setNext(updService)
	updService.setNext(storageService)

	// инициализация данных
	data := &Data{}

	fmt.Println()

	// вызов цепочки обработчиков
	device.execute(data)
}
