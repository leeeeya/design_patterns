package main

func main() {
	// инициализация объекта-получателя
	tv := &Tv{}

	// инициализация команд для отправления определённому получателю
	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	// инициализация отправителя
	onButton := &Button{
		command: onCommand,
	}

	// исполнение команды
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
