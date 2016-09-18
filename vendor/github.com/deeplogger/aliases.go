package deeplogger

import (
	"github.com/deeplogger/dispatcher"
	"github.com/deeplogger/event"
	"github.com/deeplogger/handlers"
)

//Event is a minimum loggable unit.
type Event interface {
	event.Event
}

//Creates a new event object.
func NewEvent(message string) Event {
	return event.New(message)
}

//InputHandler takes in events to be logged and passes them on the dispatcher.
//Has a name shown in the output log.
type InputHandler interface {
	handlers.InputHandler
}

//NewInputHanlder creates a new instance of output handler with a given name.
func NewInputHandler(name string) InputHandler {
	return handlers.NewInputHandler(name)
}

//NewBlankInputHandler creates a dummy input handler, which prevents accidental nil pointer dereferencing.
func NewBlankInputHandler() InputHandler {
	return handlers.NewBlankInputHandler()
}

//OutputHandler takes in events from dispatcher and passes them on to io.Writer. By default outputs to stdout.
type OutputHandler interface {
	handlers.OutputHandler
}

//NewOutputHandler creates new instance of output handler and attaches it to the dispatcher.
func NewOutputHandler(disp *dispatcher.Dispatcher, name string) OutputHandler {
	return handlers.NewOutputHandler(disp, name)
}
