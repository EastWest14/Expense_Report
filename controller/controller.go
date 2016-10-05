package controller

import (
	_ "Expense_Tracker/dbaccessor"
	"Expense_Tracker/service"
	dl "github.com/deeplogger"
)

var ContrInpHandler dl.InputHandler = dl.NewBlankInputHandler()

type DBAccess interface {
}

type Service interface {
	WaitForCommand(func(c *service.Command, err error))
}

func NewController() *ControllerModule {
	ContrInpHandler.LogMessage(`Initializing Controller.`)
	return &ControllerModule{}
}

type ControllerModule struct {
	Serv Service
	Dal  DBAccess
}

//Control is transfered from main to Controller. Application begins operating.
func (c *ControllerModule) AcceptControl() {
	ContrInpHandler.LogMessage(`Controller accepts control. Normal operation begins.`)
	c.waitForCommandFromServiceM()
}

func (c *ControllerModule) SetService(serv Service) {
	c.Serv = serv
	ContrInpHandler.LogMessage(`Service link set.`)
}

func (c *ControllerModule) SetDAL(dal DBAccess) {
	c.Dal = dal
	ContrInpHandler.LogMessage(`DAL link set.`)
}

func (c *ControllerModule) waitForCommandFromServiceM() {
	ContrInpHandler.LogMessage(`Waiting for command from Service Module.`)
	c.Serv.WaitForCommand(takeInCommand)
}

func takeInCommand(command *service.Command, err error) {
	ContrInpHandler.LogMessage(`Taking in command.`)
}
