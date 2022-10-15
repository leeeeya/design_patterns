package main

// Visitor интерфейс посетителя с методами посещения для каждой конкретной структуры
type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}
