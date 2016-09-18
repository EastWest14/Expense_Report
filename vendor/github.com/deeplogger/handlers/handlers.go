package handlers

import (
	"github.com/deeplogger/dispatcher"
	"github.com/deeplogger/event"
	bih "github.com/deeplogger/handlers/basicinputhandler"
	boh "github.com/deeplogger/handlers/basicouthandler"
	"io"
)

type InputHandler interface {
	SetDispatcher(*dispatcher.Dispatcher)
	LogEvent(event.Event)
	LogMessage(string)
}

//TODO: add type enumeration.

func NewInputHandler(name string) InputHandler {
	return bih.New(nil, name)
}

type BlankInputHandler struct{}

//Panic
func (blih *BlankInputHandler) SetDispatcher(*dispatcher.Dispatcher) {
	panic("Attempting to set dispatcher on a blank Input Handler.")
}

//Do nothing
func (blih *BlankInputHandler) LogEvent(event.Event) {
	return
}

//Do nothing
func (bih *BlankInputHandler) LogMessage(message string) {
	return
}

func NewBlankInputHandler() InputHandler {
	return &BlankInputHandler{}
}

type OutputHandler interface {
	TakeInEvent(event.Event)
	SetOutputWriter(io.Writer)
}

func NewOutputHandler(disp *dispatcher.Dispatcher, name string) OutputHandler {
	return boh.New(disp, name)
}
