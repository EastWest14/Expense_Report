package controller

import (
	_ "Expense_Tracker/dbaccessor"
	"Expense_Tracker/service"
	"testing"
)

//**************** Mock Structures ****************

type mockDbAccessor struct {
}

func newMockDbAccessor() *mockDbAccessor {
	return &mockDbAccessor{}
}

//**************** Test Controller Module Setup ****************

func TestNewController(t *testing.T) {
	contr := NewController()
	if contr == nil {
		t.Error("Failed to initialize controller")
	}
}

func TestSetDAL(t *testing.T) {
	contr := ControllerModule{}
	dal := newMockDbAccessor()
	contr.SetDAL(dal)
	if contr.Dal == nil {
		t.Error("Failed to set DAL reference")
	}
}

func TestSetService(t *testing.T) {
	contr := ControllerModule{}
	serv := service.NewService()
	contr.SetService(serv)
	if contr.Serv == nil {
		t.Error("Failed to set Service reference")
	}
}

//**************** Test Command Handling ****************

func TestConnectToComponents(t *testing.T) {

}

func TestWaitForCommand(t *testing.T) {

}
