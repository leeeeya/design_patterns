package main

// Pizzeria "элемент", который можно посетить
type Pizzeria struct {
}

// Accept имплементация метода посещения
func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}
