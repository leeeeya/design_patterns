package main

// OnCommand описание команды включения с получателем внутри
type OnCommand struct {
	device Device
}

// запуск команды on
func (c *OnCommand) execute() {
	c.device.on()
}
