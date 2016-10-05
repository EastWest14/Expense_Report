package service

import (
	dl "github.com/deeplogger"
)

//Setting up logging input point
var ServInpHandler dl.InputHandler = dl.NewBlankInputHandler()

//**************** Service Module Setup ****************

type ServiceModule struct {
}

func NewService() *ServiceModule {
	ServInpHandler.LogMessage(`Initializing a Service Module.`)
	return &ServiceModule{}
}

//**************** Incomming Command Handling ****************

func (sm *ServiceModule) WaitForCommand(commandOutputter func(c *Command, err error)) {
	ServInpHandler.LogMessage(`Waiting for command.`)
	commandOutputter(&Command{}, nil)
}

type Command struct {
}
