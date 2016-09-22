package service

import (
	dl "github.com/deeplogger"
)

var ServInpHandler dl.InputHandler = dl.NewBlankInputHandler()

type Service interface {
}

func NewService() Service {
	ServInpHandler.LogMessage(`Initializing a Service Module.`)
	return &serviceModule{}
}

type serviceModule struct {
}
