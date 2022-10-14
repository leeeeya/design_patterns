package main

// State интерфейс, описыващий поведение автомата в зависимости от конкретного состояния
type State interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}
