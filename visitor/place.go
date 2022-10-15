package main

// Place интерфейс принятия посетителей для каждого места, которое можно посетить
type Place interface {
	Accept(v Visitor) string
}
