package main

import (
	"fmt"
	"log"
)

// ItemRequestedState состояние - запрос на выдачу предмета
type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

// метод для выбора предмета - возвращает ошибку, тк предмет уже выбран и находится в состоянии получения
func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("Item already requested\n")
}

// метод добавления предмета в автомат - возвращает ошибку, тк в данном состоянии добавление невозможно
func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("Item Dispense in progress\n")
}

// метод для внесения денег в автомат
func (i *ItemRequestedState) insertMoney(money int) error {
	if money <= 0 {
		log.Fatalln("Incorrect amount of money entered")
	}
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d\n", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

// метод выдачи предмета - возвращает ошибку, тк выдача невозможна, пока не внесены деньги
func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first\n")
}
