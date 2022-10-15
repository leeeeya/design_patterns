package main

// SushiBar "элемент", который можно посетить
type SushiBar struct {
}

// Accept имплементация метода посещения
func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}
