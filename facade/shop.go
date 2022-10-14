package main

import (
	"fmt"
	"time"
)

// Shop описание магазина
type Shop struct {
	name     string
	products []Product // продукты, реализуемые в магазине
}

// Selling метод, имитирующий процесс покупки - фасад для взаимодействия с другими сервисами
func (s Shop) Selling(a Account, product string) error {
	fmt.Println("[Shop] Request to the user to receive the balance on the card")
	time.Sleep(time.Millisecond * 500)
	if err := a.card.CheckBalance(); err != nil {
		return err
	}
	fmt.Printf("[Shop] Checking if user %s can buy an item %s\n", a.name, product)
	time.Sleep(time.Millisecond * 500)
	for _, p := range s.products {
		if p.name != product {
			continue
		}
		if p.price >= a.GetBalance() {
			return fmt.Errorf("[Shop] Insufficient funds to buy product\n")
		}
		fmt.Printf("[Shop] Product %s is bought\n", p.name)
	}
	return nil
}
