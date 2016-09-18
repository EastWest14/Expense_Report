package concurrency

/*import (
	"deeplogger"
	"deeplogger/dispatcher"
	"deeplogger/event"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

const config = `{"dispatcher_name": "Dispatcher",
	"isOn": true,
	"inputHandlers": ["Input", "Input2"],
	"outputHandlers": ["Output"],
	"dispatchRules": [
		{"input":"Input", "output": "Output"},
		{"input":"Input2", "output": "Output"}
	]
}`

var inpHandler deeplogger.InputHandler = deeplogger.NewBlankInputHandler()
var inpHandler2 deeplogger.InputHandler = deeplogger.NewBlankInputHandler()
var disp *dispatcher.Dispatcher
var outHandler deeplogger.OutputHandler

func setupWithConfigString() {
	inpHandlers, d, outHandlers, err := deeplogger.ConstructLoggerFromConfig(config)
	if err != nil {
		panic("Failed loading DL from config. " + err.Error())
	}
	disp = d
	inpHandler = inpHandlers["Input"]
	inpHandler2 = inpHandlers["Input2"]
	outHandler = outHandlers["Output"]
}

func setupManual() {
	disp = dispatcher.New("Dispatcher")
	disp.AddInputHandler("Input", true)
	disp.AddInputHandler("Input2", true)
	disp.AddRule(dispatcher.NewRule(dispatcher.NewMatchCondition("Input"), "Output"))
	disp.AddRule(dispatcher.NewRule(dispatcher.NewMatchCondition("Input2"), "Output"))
	inpHandler = deeplogger.NewInputHandler("Input")
	inpHandler2 = deeplogger.NewInputHandler("Input2")
	inpHandler.SetDispatcher(disp)
	inpHandler2.SetDispatcher(disp)
	outHandler = deeplogger.NewOutputHandler(disp, "Output")
}

func TestMain(m *testing.M) {
	setupWithConfigString()
	res1 := m.Run()
	setupManual()
	res2 := m.Run()
	if res1 == 0 && res2 == 0 {
		os.Exit(0)
	} else if res1 != 0 {
		os.Exit(res1)
	} else {
		os.Exit(res2)
	}
}

func bombardDeepLogger() {
	for i := 0; i < 5; i++ {
		go hitDL(inpHandler)
		go hitDL(inpHandler2)
	}
}

func hitDL(inpH deeplogger.InputHandler) {
	for i := 0; i < 10; i++ {
		inpH.LogEvent(&timestampEvent{})
		time.Sleep(time.Microsecond * time.Duration(1*rand.Intn(10)))
	}
}

var _ event.Event = &timestampEvent{}

type timestampEvent struct {
	inputHandlerName      string
	message               string
	loggedIntoDispatcherT time.Time
}

func (te *timestampEvent) InputHandlerName() string {
	return te.inputHandlerName
}

func (te *timestampEvent) SetInputHandlerName(name string) {
	te.inputHandlerName = name
	te.loggedIntoDispatcherT = time.Now()
	te.message = fmt.Sprintf("%d", time.Now().UnixNano())
}

func (te *timestampEvent) EventMessage() string {
	return te.message
}

type writeRedirector struct {
	redirChan chan string
}

func (wr *writeRedirector) Write(input []byte) (n int, err error) {
	strInput := string(input)
	parts := strings.Split(strInput, " ")
	if len(parts) != 2 {
		panic("Output format incorrect")
	}
	timestamp := parts[1]
	wr.redirChan <- string(timestamp)
	return 0, nil
}

var lastTimestamp string

func checkOrder(redirChan chan string) bool {
	timestamp := <-redirChan
	if timestamp <= lastTimestamp {
		fmt.Println(timestamp + " " + lastTimestamp)
		return false
	}
	lastTimestamp = timestamp
	return true
}

func TestCheckOrder(t *testing.T) {
	defer func() {
		lastTimestamp = "0"
	}()
	lastTimestamp = "0"
	rChan := make(chan string)
	go func() {
		rChan <- "123"
	}()
	order := checkOrder(rChan)
	if !order {
		t.Error("Check order doesn't correctly check order.")
	}
	go func() {
		rChan <- "124"
	}()
	order = checkOrder(rChan)
	if !order {
		t.Error("Check order doesn't correctly check order.")
	}
}

func TestOrderPreservation(t *testing.T) {
	defer func() {
		outHandler.SetOutputWriter(os.Stdout)
	}()
	rChannel := make(chan string)
	outHandler.SetOutputWriter(&writeRedirector{redirChan: rChannel})
	go func() {
		for i := 0; i < 100; i++ {
			inOrder := checkOrder(rChannel)
			if !inOrder {
				t.Error("Order broken")
			}
		}
	}()
	bombardDeepLogger()
	time.Sleep(time.Second)
}
*/
