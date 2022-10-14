package main

import (
	"fmt"
	"log"
)

var (
	bank = Bank{
		name:  "Some Bank",
		cards: []Card{},
	}
	card1 = Card{
		name:    "CRD-1",
		balance: 200,
		bank:    &bank,
	}
	card2 = Card{
		name:    "CRD-2",
		balance: 5,
		bank:    &bank,
	}
	customer1 = Account{
		name: "Customer-1",
		card: &card1,
	}
	customer2 = Account{
		name: "Customer-2",
		card: &card2,
	}
	product = Product{
		name:  "Nemoloko",
		price: 150,
	}
	shop = Shop{
		name: "24",
		products: []Product{
			product,
		},
	}
)

func main() {
	fmt.Println("[Bank] Issuance of cards")
	bank.cards = append(bank.cards, card1, card2)
	fmt.Println()

	// покупка товара первым покупателем
	fmt.Printf("[%s]\n\n", customer1.name)
	if err := shop.Selling(customer1, product.name); err != nil {
		log.Print(err)
	}

	fmt.Println()

	// покупка товара вторым покупателем
	fmt.Printf("[%s]\n\n", customer2.name)
	if err := shop.Selling(customer2, product.name); err != nil {
		log.Print(err)
	}
}
