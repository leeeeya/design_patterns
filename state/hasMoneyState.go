package main

import "fmt"

// GetMoneyState состояние - автомат получил деньги
type GetMoneyState struct {
	vendingMachine *VendingMachine
}

// запрос на получение предмета невозможен, тк ещё обрабатывается предыдущий такой же запрос
func (i *GetMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress\n")
}

// добавление предметов в автомат в этом состоянии невозможно
func (i *GetMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress\n")
}

// внесение оплаты в данном состоянии невозможно
func (i *GetMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item dispense in progress\n")
}

// метод выдачи предмета из автомата.
// устанавливает следующее состояние в зависимости от того, остались предметы в автомате или нет
func (i *GetMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}
