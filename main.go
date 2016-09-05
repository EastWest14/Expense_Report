package main

import (
	"Expense_Tracker/config"
	"errors"
	dl "github.com/Deep-Logger"
	"github.com/Deep-Logger/event"
	dlhandlers "github.com/Deep-Logger/handlers"
	"io/ioutil"
	"os"
)

const (
	TEMP_PROD_CONFIG_PATH = "./conf_files/prod_config/prod_config.json"

	MAIN_INPUT_HANDLER_NAME = "Main"
	OUTPUT_HANDLER_NAME     = "Out"
)

func main() {
	config, err := config.LoadConfigFromFile(TEMP_PROD_CONFIG_PATH)
	if err != nil {
		panic(err.Error())
	}
	err = constructDeepLoggerSystem(config.DeepLoggerPath)
	if err != nil {
		panic(err.Error())
	}
	mainInpHandler.LogEvent(event.New("Hello"))
}

var mainInpHandler dlhandlers.InputHandler
var outHandler dlhandlers.OutputHandler

//temp
func constructDeepLoggerSystem(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return errors.New("Failed opening DeepLogger config: " + err.Error())
	}
	configRaw, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.New("Failed reading DeepLogger config: " + err.Error())
	}
	inpHandlers, _, outHandlers := dl.ConstructLoggerFromConfig(string(configRaw))
	var ok bool
	mainInpHandler, ok = inpHandlers[MAIN_INPUT_HANDLER_NAME]
	if !ok {
		return errors.New("Main input handler not found")
	}
	outHandler, ok = outHandlers[OUTPUT_HANDLER_NAME]
	if !ok {
		return errors.New("Output handler not found")
	}
	outHandler.SetOutputWriter(os.Stdout)
	return nil
}
