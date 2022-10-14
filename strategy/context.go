package main

// Context контекст, содержащий конкретную стратегию
type Context struct {
	strategy StrategySort
}

// Algorithm задаёт конкретную стратегию сортировки
func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

// Sort сортирует данные по заданной стратегии
func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}
