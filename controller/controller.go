package controller

import (
	_ "Expense_Tracker/dbaccessor"
	"Expense_Tracker/service"
	dl "github.com/deeplogger"
	"github.com/gAssert"
)

//Setting up logging input point
var ContrInpHandler dl.InputHandler = dl.NewBlankInputHandler()

//**************** Controller Module Setup ****************

type ControllerModule struct {
	Serv Service
	Dal  DBAccess
}

type DBAccess interface {
}

type Service interface {
	WaitForCommand(service.CommandInputter)
}

func NewController() *ControllerModule {
	ContrInpHandler.LogMessage(`Initializing Controller.`)
	return &ControllerModule{}
}

func (c *ControllerModule) SetService(serv Service) {
	c.Serv = serv
	ContrInpHandler.LogMessage(`Service link set.`)
}

func (c *ControllerModule) SetDAL(dal DBAccess) {
	c.Dal = dal
	ContrInpHandler.LogMessage(`DAL link set.`)
}

//**************** Command Handling ****************

//Control is transfered from main to Controller. Application begins operating.
func (c *ControllerModule) AcceptControl() {
	gAssert.AssertHard(c.Serv != nil, "Service link of controller not set. Controller cannot accept control.")
	gAssert.AssertHard(c.Dal != nil, "Database access link of controller not set. Controller cannot accept control.")
	ContrInpHandler.LogMessage(`Controller accepts control. Normal operation begins.`)
	c.waitForCommandFromServiceM()
}

func (c *ControllerModule) waitForCommandFromServiceM() {
	ContrInpHandler.LogMessage(`Waiting for command from Service Module.`)
	c.Serv.WaitForCommand(takeInCommand)
}

//TODO: fill command interface

//takeInCommand is passed to the service module as a variable
func takeInCommand(command interface{}, err error) {
	ContrInpHandler.LogMessage(`Taking in command.`)
}
