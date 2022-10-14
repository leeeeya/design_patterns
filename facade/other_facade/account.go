package main

// Account описание покупателя
type Account struct {
	name string
	card *Card // карта для оплачивания
}

// GetBalance получение баланса карты покупателя
func (a Account) GetBalance() float64 {
	return a.card.balance
}
