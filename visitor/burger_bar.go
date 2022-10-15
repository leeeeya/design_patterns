package main

// BurgerBar "элемент", который можно посетить
type BurgerBar struct {
}

// Accept имплементация метода посещения
func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}
