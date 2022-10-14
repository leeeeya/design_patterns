package main

// Command общий интерфейс команд с методом запуска
type Command interface {
	execute()
}
