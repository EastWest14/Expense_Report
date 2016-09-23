package service

import (
	dl "github.com/deeplogger"
)

var ServInpHandler dl.InputHandler = dl.NewBlankInputHandler()

type Service interface {
	WaitForCommand(func(c *Command, err error))
}

func NewService() Service {
	ServInpHandler.LogMessage(`Initializing a Service Module.`)
	return &serviceModule{}
}

type serviceModule struct {
}

func (sm *serviceModule) WaitForCommand(commandOutputter func(c *Command, err error)) {
	ServInpHandler.LogMessage(`Waiting for command.`)
	commandOutputter(&Command{}, nil)
}

type Command struct {
}
