package main

import (
	"fmt"
)

func main() {
	// создание нового объекта контекста с исходными параметрами
	vendingMachine := newVendingMachine(0, 10)

	fmt.Println()
	// запрос на выдачу предмета
	err := vendingMachine.requestItem()
	if err != nil {
		fmt.Println(err.Error())
	}

	// внесение оплаты
	err = vendingMachine.insertMoney(10)
	if err != nil {
		fmt.Println(err.Error())
	}

	// выдача предмета
	err = vendingMachine.dispenseItem()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println()

	// добавление определённого количества предметов в автомат
	err = vendingMachine.addItem(1)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = vendingMachine.insertMoney(9)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		fmt.Println(err.Error())
	}
}
