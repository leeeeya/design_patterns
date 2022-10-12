package main

// Gum базовая структура, от которой будут наследоваться объекты
type Gum struct {
	name  string
	taste string
}

func (g *Gum) setName(name string) {
	g.name = name
}

func (g *Gum) getName() string {
	return g.name
}

func (g *Gum) setTaste(taste string) {
	g.taste = taste
}

func (g *Gum) getTaste() string {
	return g.taste
}
