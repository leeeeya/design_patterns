package main

import "fmt"

// People конкретный посетитель, реализующий интерфейс Visitor
type People struct {
	name string
}

// VisitSushiBar реализация метода посещения SushiBar.
func (v *People) VisitSushiBar(p *SushiBar) string {
	return fmt.Sprintf("%s buys suhhi\n", v.name)
}

// VisitPizzeria реализация метода посещения Pizzeria.
func (v *People) VisitPizzeria(p *Pizzeria) string {
	return fmt.Sprintf("%s buys pizza\n", v.name)
}

// VisitBurgerBar реализация метода посещения BurgerBar.
func (v *People) VisitBurgerBar(p *BurgerBar) string {
	return fmt.Sprintf("%s buys burger\n", v.name)
}
