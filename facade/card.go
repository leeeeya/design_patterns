package main

import (
	"fmt"
	"time"
)

// Card описание карты
type Card struct {
	name    string
	balance float64
	bank    *Bank // банк, в котором была оформалена карта
}

// CheckBalance проверка баланса карты
func (c Card) CheckBalance() error {
	fmt.Println("[Card] Request to the bank to check the balance")
	time.Sleep(time.Millisecond * 800)
	return c.bank.CheckBalance(c.name)
}
