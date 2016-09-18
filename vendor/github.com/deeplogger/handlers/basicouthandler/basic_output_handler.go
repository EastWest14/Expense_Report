package basicouthandler

import (
	"fmt"
	"github.com/deeplogger/dispatcher"
	"github.com/deeplogger/event"
	"io"
	"os"
)

type BasicOutputHandler struct {
	Dispatcher        *dispatcher.Dispatcher
	OutputHandlerName string
	OutputWriter      io.Writer
}

func New(disp *dispatcher.Dispatcher, name string) *BasicOutputHandler {
	boh := BasicOutputHandler{Dispatcher: disp, OutputHandlerName: name}
	disp.AddOutputHandler(boh.OutputHandlerName, boh.TakeInEvent)
	//By default writes to stdout
	boh.SetOutputWriter(os.Stdout)
	return &boh
}

func (boh *BasicOutputHandler) SetOutputWriter(writer io.Writer) {
	boh.OutputWriter = writer
}

func (boh *BasicOutputHandler) TakeInEvent(ev event.Event) {
	evString := fmt.Sprintln("[" + ev.InputHandlerName() + "]: " + ev.EventMessage())
	boh.outputData([]byte(evString))
}

func (boh *BasicOutputHandler) outputData(data []byte) {
	boh.OutputWriter.Write(data)
}
