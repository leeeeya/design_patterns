package main

// Service интерфейс сервиса
type Service interface {
	execute(*Data)   // метод выполнения какого-либо обработчика
	setNext(Service) // задача следующего обрабочика
}

// Data данные, которые необходимо обработать в цепочке
type Data struct {
	getSource bool // флаг - приняты данные от источника или нет
	updSource bool // флаг - обработаны ли данные в обработчике
}
