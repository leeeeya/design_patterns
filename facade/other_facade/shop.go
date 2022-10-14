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

// Selling метод, имитирующий процесс покупки
func (s Shop) Selling(a Account, product string) error {
	fmt.Println("[Магазин] Запрос к пользователю для получения остата по карте")
	time.Sleep(time.Millisecond * 500)
	if err := a.card.CheckBalance(); err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка, может ли пользователь %s купить товар %s\n", a.name, product)
	time.Sleep(time.Millisecond * 500)

}
