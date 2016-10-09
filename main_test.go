package main

import (
	"Expense_Tracker/controller"
	"Expense_Tracker/dbaccessor"
	"Expense_Tracker/service"
	dl "github.com/deeplogger"
	"testing"
)

const DL_CONFIG_PATH = "./conf_files/deep_logger_config/production_dl_config.json"

//**************** Test Deep Logger Setup ****************

func TestConstructDeepLoggerSystem(t *testing.T) {
	controller.ContrInpHandler = nil
	dbaccessor.DBAccessorInpHandler = nil
	service.ServInpHandler = nil

	constructDeepLoggerSystem(DL_CONFIG_PATH)

	if controller.ContrInpHandler == nil {
		t.Error("Controller input handler not set.")
	}
	if dbaccessor.DBAccessorInpHandler == nil {
		t.Error("DB accessor input handler not set.")
	}
	if service.ServInpHandler == nil {
		t.Error("Service input handler not set.")
	}
}

//**************** Utilities ****************

func disableDLLogging() {
	mainInpHandler = dl.NewBlankInputHandler()
	service.ServInpHandler = dl.NewBlankInputHandler()
	dbaccessor.DBAccessorInpHandler = dl.NewBlankInputHandler()
	controller.ContrInpHandler = dl.NewBlankInputHandler()
}
