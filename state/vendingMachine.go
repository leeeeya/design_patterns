package main

import (
	"fmt"
	"log"
)

// VendingMachine струкутра контекста - вендингового автомата
type VendingMachine struct {
	// все возможные состояния автомата
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State

	// текущее состояние
	currentState State

	itemCount int
	itemPrice int
}

// создание нового объекта контекста
func newVendingMachine(itemCount, itemPrice int) *VendingMachine {
	if itemCount < 0 || itemPrice < 0 {
		log.Fatalln("incorrect number or price of item!")
	}
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &GetMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

// реализация методов состояний для конкретного контекста
func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendingMachine) incrementItemCount(count int) {
	if count <= 0 {
		log.Fatalln("Incorrect number of items entered!")
	}
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}
