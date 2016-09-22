package controller

import (
	"Expense_Tracker/service"
	dl "github.com/deeplogger"
)

var ContrInpHandler dl.InputHandler = dl.NewBlankInputHandler()

type Controller interface {
	AcceptControl()
	SetService(service.Service)
}

func NewController() Controller {
	ContrInpHandler.LogMessage(`Initializing Controller.`)
	return &controllerModule{}
}

type controllerModule struct {
	serv service.Service
}

//Control is transfered from main to Controller. Application begins operating.
func (c *controllerModule) AcceptControl() {

}

func (c *controllerModule) SetService(serv service.Service) {
	c.serv = serv
	ContrInpHandler.LogMessage(`Service link set.`)
}

func (c *controllerModule) waitForCommand() {

}
