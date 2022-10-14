package main

import "fmt"

// DataStorageService сервис для сохранения обработанных данных
type DataStorageService struct {
	next Service // следующий обработчик
}

// реализация сохранения данных
func (storSvc *DataStorageService) execute(data *Data) {
	// если данные ещё не обработаны, они не сохраняются
	if !data.updSource {
		fmt.Println("Data not updated")
		return
	}
	// сохранение данных
	fmt.Println("Data saved")
}

// передача обработки следующему звену
func (storSvc *DataStorageService) setNext(svc Service) {
	storSvc.next = svc
}
