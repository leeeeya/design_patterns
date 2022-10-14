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
	fmt.Println("[Магазин] Запрос к пользователю для получения остата по карте")
	time.Sleep(time.Millisecond * 500)
	if err := a.card.CheckBalance(); err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка, может ли пользователь %s купить товар %s\n", a.name, product)
	time.Sleep(time.Millisecond * 500)
	for _, p := range s.products {
		if p.name != product {
			continue
		}
		if p.price >= a.GetBalance() {
			return fmt.Errorf("[Магазин] Недостаточно средств для покупки товара\n")
		}
		fmt.Printf("[Магазин] Товар %s куплен\n", p.name)
	}
	return nil
}
