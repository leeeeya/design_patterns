package main

import "fmt"

// Tv - получатель, выставляет флаг в зависимости от результата команды
type Tv struct {
	isRunning bool
}

// имплементация интерфейса Device
func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
