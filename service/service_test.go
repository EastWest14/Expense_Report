package service

import (
	"testing"
)

//**************** Test Service Module Setup ****************

func TestNewService(t *testing.T) {
	serv := NewService()
	if serv == nil {
		t.Error("Failed to initialize service")
	}
}
