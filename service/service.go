package service

import (
	"Expense_Tracker/expense"
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

type CommandInputter func(command interface{}, err error)

func (sm *ServiceModule) WaitForCommand(commandInput CommandInputter) {
	ServInpHandler.LogMessage(`Waiting for command.`)
	commandInput(expense.New(123.5), nil)
}
