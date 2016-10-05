package main

import (
	"Expense_Tracker/controller"
	"Expense_Tracker/dbaccessor"
	"Expense_Tracker/service"
	"errors"
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

//**************** Test Setup Full Application ****************

func TestSetupApplication(t *testing.T) {

}

//**************** Test Setup Database Module ****************

type mockDBAccess struct {
	connectError         error
	checkConnectionError error
}

func (m *mockDBAccess) SetDBConfig(dbConf *dbaccessor.DBConfig) {
	return
}

func (m *mockDBAccess) Connect(driver string) error {
	return m.connectError
}

func (m *mockDBAccess) CheckConnection() error {
	return m.checkConnectionError
}

func TestVerifyDBConnection(t *testing.T) {
	disableDLLogging()

	mockAccess := &mockDBAccess{connectError: nil, checkConnectionError: nil}
	err := verifyDBConnection(mockAccess)
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	mockAccess = &mockDBAccess{connectError: errors.New(""), checkConnectionError: nil}
	err = verifyDBConnection(mockAccess)
	if err == nil {
		t.Errorf("Expected error connecting, got none")
	}
	mockAccess = &mockDBAccess{connectError: nil, checkConnectionError: errors.New("")}
	err = verifyDBConnection(mockAccess)
	if err == nil {
		t.Errorf("Expected error pinging, got none")
	}
}

//**************** Test Setup Service Module ****************

func TestSetupServiceModule(t *testing.T) {
	servM := setupServiceModule()
	if servM == nil {
		t.Error("Failed to setup service module")
	}
}

//**************** Test Setup Controller ****************

func TestSetupControllerModule(t *testing.T) {
	disableDLLogging()

	contM := setupControllerModule()
	if contM == nil {
		t.Error("Failed to setup controller module")
	}
}

//**************** Test Utilities ****************

func disableDLLogging() {
	mainInpHandler = dl.NewBlankInputHandler()
	service.ServInpHandler = dl.NewBlankInputHandler()
	dbaccessor.DBAccessorInpHandler = dl.NewBlankInputHandler()
	controller.ContrInpHandler = dl.NewBlankInputHandler()
}
