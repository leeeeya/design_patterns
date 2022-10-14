package main

// Device интерфейс получателя команд
type Device interface {
	on()
	off()
}
