package main

// City множество "элементов", которые можно посетить
type City struct {
	places []Place
}

// Add добавление новых элементов в множество City
func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

// Accept имплементация метода посещения для всех элементов
func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}
