package main

// OffCommand описание команды выключения с получателем внутри
type OffCommand struct {
	device Device
}

// запуск команды off
func (c *OffCommand) execute() {
	c.device.off()
}
