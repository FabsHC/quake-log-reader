package main

import (
	"quake-log-reader/cmd/handler"
	"quake-log-reader/cross_cut"
)

type (
	InputHandler interface {
		Execute()
	}

	Main struct {
		inputHandler InputHandler
	}
)

func main() {
	reg := cross_cut.NewRegister()
	terminalHandler := &Main{
		inputHandler: handler.NewTerminalHandler(reg.ProcessEventUseCase),
	}
	terminalHandler.inputHandler.Execute()
}
