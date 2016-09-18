package basicinputhandler

import (
	"github.com/deeplogger/dispatcher"
	"github.com/deeplogger/event"
)

type BasicInputHandler struct {
	Dispatcher       *dispatcher.Dispatcher
	InputHandlerName string
}

func New(dl *dispatcher.Dispatcher, inputName string) *BasicInputHandler {
	return &BasicInputHandler{Dispatcher: dl, InputHandlerName: inputName}
}

func (bih *BasicInputHandler) SetDispatcher(d *dispatcher.Dispatcher) {
	bih.Dispatcher = d
}

func (bih *BasicInputHandler) LogEvent(ev event.Event) {
	if bih.Dispatcher == nil {
		panic("No dispatcher registered.")
		return
	}
	ev.SetInputHandlerName(bih.InputHandlerName)
	bih.Dispatcher.InputEvent(ev)
}

func (bih *BasicInputHandler) LogMessage(message string) {
	bih.LogEvent(event.New(message))
}
