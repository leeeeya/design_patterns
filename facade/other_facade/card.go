package main

import (
	"fmt"
	"time"
)

// Card описание карты
type Card struct {
	number  string
	balance float64
	bank    *Bank // банк, в котором была оформалена карта
}

// CheckBalance проверка баланса карты
func (c Card) CheckBalance() error {
	fmt.Println("[Карта]	Запрос в банк для проверки остатка")
	time.Sleep(time.Millisecond * 800)
	return c.bank.CheckBalance(c.number)
}
