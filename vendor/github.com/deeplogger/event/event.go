//Event package is a lower level package used to construct event objects.
package event

import ()

type Event interface {
	InputHandlerName() string
	SetInputHandlerName(string)
	EventMessage() string
}

type SimpleEvent struct {
	inputHandlerName string
	message          string
}

func New(message string) *SimpleEvent {
	return &SimpleEvent{message: message}
}

func (se *SimpleEvent) InputHandlerName() string {
	return se.inputHandlerName
}

func (se *SimpleEvent) EventMessage() string {
	return se.message
}

func (se *SimpleEvent) SetInputHandlerName(inputHandlerName string) {
	se.inputHandlerName = inputHandlerName
	return
}
