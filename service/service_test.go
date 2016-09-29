package service

import (
	"testing"
)

func TestNewService(t *testing.T) {
	serv := NewService()
	if serv == nil {
		t.Error("Failed to initialize service")
	}
}
