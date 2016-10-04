package controller

import (
	_ "Expense_Tracker/dbaccessor"
	"Expense_Tracker/service"
	dl "github.com/deeplogger"
)

var ContrInpHandler dl.InputHandler = dl.NewBlankInputHandler()

type DBAccess interface {
}

type Controller interface {
	AcceptControl()
	SetService(service.Service)
	SetDAL(DBAccess)
}

func NewController() Controller {
	ContrInpHandler.LogMessage(`Initializing Controller.`)
	return &controllerModule{}
}

type controllerModule struct {
	serv service.Service
	dal  DBAccess
}

//Control is transfered from main to Controller. Application begins operating.
func (c *controllerModule) AcceptControl() {
	ContrInpHandler.LogMessage(`Controller accepts control. Normal operation begins.`)
	c.waitForCommandFromServiceM()
}

func (c *controllerModule) SetService(serv service.Service) {
	c.serv = serv
	ContrInpHandler.LogMessage(`Service link set.`)
}

func (c *controllerModule) SetDAL(dal DBAccess) {
	c.dal = dal
	ContrInpHandler.LogMessage(`DAL link set.`)
}

func (c *controllerModule) waitForCommandFromServiceM() {
	ContrInpHandler.LogMessage(`Waiting for command from Service Module.`)
	c.serv.WaitForCommand(takeInCommand)
}

func takeInCommand(command *service.Command, err error) {
	ContrInpHandler.LogMessage(`Taking in command.`)
}
