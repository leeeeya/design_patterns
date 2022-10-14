package main

import "fmt"

// HasItemState состояние - автомат имеет предмет
type HasItemState struct {
	vendingMachine *VendingMachine
}

// метод для выбора предмета из автомата
func (i *HasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present\n")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested) // установление следующего состояния
	return nil
}

// метод добавления нового предмета в автомат
func (i *HasItemState) addItem(count int) error {
	fmt.Printf("%d Items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

// метод для внесения денег в автомат - в данном состоянии всегда будет возварщать ошибку, тк предмет ещё не выбран
func (i *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first\n")
}

// метод для выдачи предмета - в данном состоянии всегда будет возварщать ошибку, тк предмет ещё не выбран и не оплачен
func (i *HasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first\n")
}
