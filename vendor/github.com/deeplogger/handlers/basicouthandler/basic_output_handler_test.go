package basicouthandler

import (
	"deeplogger/event"
	"os"
	"testing"
)

func TestTakeInEvent(t *testing.T) {
	boh := BasicOutputHandler{Dispatcher: nil, OutputHandlerName: "ABC", OutputWriter: os.Stdout}

	ev := event.New("Hello world!")
	ev.SetInputHandlerName("XYZ")
	boh.TakeInEvent(ev)
}
