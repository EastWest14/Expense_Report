//Deeplogger is a package for logging, debugging and automated testing of concurrent systems.
package deeplogger

import (
	"encoding/json"
	"errors"
	"github.com/deeplogger/dispatcher"
	"io/ioutil"
	"os"
)

//ConstructLoggerFromConfig returns input handlers, dispatcher and output handlers that can be used to construct the deep logger system.
func ConstructLoggerFromConfig(config string) (inputHandlers map[string]InputHandler, disp *dispatcher.Dispatcher, outputHandlers map[string]OutputHandler, err error) {
	disp = dispatcher.New("")
	inputHandlers = map[string]InputHandler{}
	outputHandlers = map[string]OutputHandler{}
	var dat map[string]interface{}
	err = json.Unmarshal([]byte(config), &dat)
	if err != nil {
		return nil, nil, nil, err
	}
	if dat == nil {
		return inputHandlers, disp, outputHandlers, errors.New("Unmarshalled data is nil")
	}

	dName, ok := dat["dispatcher_name"].(string)
	if !ok {
		return nil, nil, nil, errors.New("dispatcher_name not present")
	}
	disp.SetName(dName)

	isOn, ok := dat["isOn"].(bool)
	if !ok {
		return nil, nil, nil, errors.New("isOn not present")
	}
	if isOn {
		disp.TurnOn()
	} else {
		disp.TurnOff()
	}

	inNames, ok := dat["inputHandlers"].([]interface{})
	if !ok {
		return nil, nil, nil, errors.New("inputHandlers not present")
	}
	for _, inName := range inNames {
		//modifying dispatcher
		stringName := inName.(string)
		disp.AddInputHandler(stringName, true) //TODO: is on?

		//creating handlers
		handl := NewInputHandler(stringName)
		handl.SetDispatcher(disp)
		if _, present := inputHandlers[stringName]; present {
			panic("Attempt to ad duplicate input handlers.")
		}
		inputHandlers[stringName] = handl
	}

	outNames, ok := dat["outputHandlers"].([]interface{})
	if !ok {
		return nil, nil, nil, errors.New("outputHandlers not present.")
	}
	for _, outName := range outNames {
		stringName := outName.(string)
		//Creating handlers
		handl := NewOutputHandler(disp, stringName)
		if _, present := outputHandlers[stringName]; present {
			panic("Attempt to ad duplicate output handlers.")
		}
		outputHandlers[stringName] = handl
	}

	dispatchRulesData, ok := dat["dispatchRules"].([]interface{})
	if !ok {
		return nil, nil, nil, errors.New("dispatcherRules not found.")
	}
	for _, dispRule := range dispatchRulesData {
		dRule := dispRule.(map[string]interface{})
		input, ok := dRule["input"].(string)
		if !ok {
			return nil, nil, nil, errors.New("Invalid dispatcher rules.")
		}
		output, ok := dRule["output"].(string)
		if !ok {
			return nil, nil, nil, errors.New("Invalid dispatcher rules.")
		}
		disp.AddRule(dispatcher.NewRule(dispatcher.NewMatchCondition(input), output))
	}

	return inputHandlers, disp, outputHandlers, nil
}

//TODO: unparse input handlers should be a separate function with tests.
//TODO: unparse output handlers should be a separate function with tests.
//TODO: unparse dispatcher rules should be a separate function with tests.

//Construct Lgger from filepath returns joined components of the Deep Logger system. Implicitly calls ConstructLoggerFromConfig.
func ConstructLoggerFromFilepath(filepath string) (inputHandlers map[string]InputHandler, disp *dispatcher.Dispatcher, outputHandlers map[string]OutputHandler, err error) {
	content, err := loadFileToString(filepath)
	if err != nil {
		return nil, nil, nil, err
	}
	return ConstructLoggerFromConfig(content)
}

func loadFileToString(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", errors.New("Failed opening file: " + err.Error())
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", errors.New("Failed reading file: " + err.Error())
	}
	return string(content), nil
}

//CountWriter is a mock object that implements io.Writer. Used to count number of calls to Write.
type CountWriter struct {
	V int
}

//Write increments internal counter by one.
func (iw *CountWriter) Write(input []byte) (n int, err error) {
	iw.V++
	return 0, nil
}
