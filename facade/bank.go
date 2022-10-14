package main

import (
	"fmt"
	"time"
)

// Bank описание банка
type Bank struct {
	name  string
	cards []Card // список карт, оформленных в конкретном банке
}

// CheckBalance получение информации о балансе
func (b Bank) CheckBalance(cardNum string) error {
	fmt.Printf("[Банк] Получение остатка по карте %s\n", cardNum)
	time.Sleep(time.Millisecond * 300)
	for _, card := range b.cards {
		if card.name != cardNum {
			continue
		}
		if card.balance <= 0 {
			return fmt.Errorf("[Bank] Insufficient funds\n")
		}
	}
	fmt.Println("[Bank] The balance is positive")
	return nil
}
