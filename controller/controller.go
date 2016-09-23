package controller

import (
	"Expense_Tracker/dbaccessor"
	"Expense_Tracker/service"
	dl "github.com/deeplogger"
)

var ContrInpHandler dl.InputHandler = dl.NewBlankInputHandler()

type Controller interface {
	AcceptControl()
	SetService(service.Service)
	SetDAL(dbaccessor.DBAccess)
}

func NewController() Controller {
	ContrInpHandler.LogMessage(`Initializing Controller.`)
	return &controllerModule{}
}

type controllerModule struct {
	serv service.Service
	dal  dbaccessor.DBAccess
}

//Control is transfered from main to Controller. Application begins operating.
func (c *controllerModule) AcceptControl() {

}

func (c *controllerModule) SetService(serv service.Service) {
	c.serv = serv
	ContrInpHandler.LogMessage(`Service link set.`)
}

func (c *controllerModule) SetDAL(dal dbaccessor.DBAccess) {
	c.dal = dal
	ContrInpHandler.LogMessage(`DAL link set.`)
}

func (c *controllerModule) waitForCommand() {

}
