package main

// Button структура-отправитель команды
type Button struct {
	command Command
}

// основной метод - делегирование выполнения действия команде
func (b *Button) press() {
	b.command.execute()
}
