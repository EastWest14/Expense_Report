package controller

import (
	_ "Expense_Tracker/dbaccessor"
	"Expense_Tracker/service"
	"testing"
)

type mockDbAccessor struct {
}

func newMockDbAccessor() *mockDbAccessor {
	return &mockDbAccessor{}
}

func TestNewController(t *testing.T) {
	contr := NewController()
	if contr == nil {
		t.Error("Failed to initialize controller")
	}
}

func TestSetDAL(t *testing.T) {
	contr := controllerModule{}
	dal := newMockDbAccessor()
	contr.SetDAL(dal)
	if contr.dal == nil {
		t.Error("Failed to set DAL reference")
	}
}

func TestSetService(t *testing.T) {
	contr := controllerModule{}
	serv := service.NewService()
	contr.SetService(serv)
	if contr.serv == nil {
		t.Error("Failed to set Service reference")
	}
}

func TestConnectToComponents(t *testing.T) {

}

func TestWaitForCommand(t *testing.T) {

}
