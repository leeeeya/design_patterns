package main

import "fmt"

// NoItemState состояние - нет предметов в автомате
type NoItemState struct {
	vendingMachine *VendingMachine
}

// выдача предмета невозможна, тк нет предметов в автомате
func (i *NoItemState) requestItem() error {
	return fmt.Errorf("Item out of stock\n")
}

// добавление предметов в пустой автомат
func (i *NoItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem) // установление следующего состояния
	return nil
}

// невозможно принять оплату, тк предметы для покупки отсутствуют
func (i *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock\n")
}

// в данном состоянии выдача предметов невозможна
func (i *NoItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock\n")
}
