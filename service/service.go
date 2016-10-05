package service

import (
	dl "github.com/deeplogger"
)

var ServInpHandler dl.InputHandler = dl.NewBlankInputHandler()

func NewService() *ServiceModule {
	ServInpHandler.LogMessage(`Initializing a Service Module.`)
	return &ServiceModule{}
}

type ServiceModule struct {
}

func (sm *ServiceModule) WaitForCommand(commandOutputter func(c *Command, err error)) {
	ServInpHandler.LogMessage(`Waiting for command.`)
	commandOutputter(&Command{}, nil)
}

type Command struct {
}
