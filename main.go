package main

import (
	"Expense_Tracker/config"
	"Expense_Tracker/dbaccessor"
	"errors"
	"fmt"
	dl "github.com/Deep-Logger"
	"github.com/Deep-Logger/event"
	dlhandlers "github.com/Deep-Logger/handlers"
	"io/ioutil"
	"os"
)

const (
	TEMP_PROD_CONFIG_PATH = "./conf_files/prod_config/prod_config.json"

	MAIN_INPUT_HANDLER_NAME        = "Main"
	DB_ACCESSOR_INPUT_HANDLER_NAME = "DB_Accessor"
	OUTPUT_HANDLER_NAME            = "Out"
)

var dbAccessModule dbaccessor.DBAccessModule

func main() {
	config, err := config.LoadConfigFromFile(TEMP_PROD_CONFIG_PATH)
	if err != nil {
		panic(err.Error())
	}
	err = constructDeepLoggerSystem(config.DeepLoggerPath)
	if err != nil {
		panic(err.Error())
	}
	mainInpHandler.LogEvent(event.New(`Deep Logger succesfully configured.`))
	setupApplication(config)
}

func setupApplication(conf *config.Config) {
	dbaccessor.DBAccessorInpHandler = dbAccessorInpHandler
	setupDBAccessModule(conf)
	mainInpHandler.LogEvent(event.New(`DB Access Module succesfully configured. Connection to DB established.`))
}

func setupDBAccessModule(conf *config.Config) {
	dbAccessModule = dbaccessor.NewDBAccessModule()
	dbAccessModule.SetDBConfig(dbaccessor.NewDBConfig(conf.DBUser, conf.DBPassword, conf.DBName))
	err := dbAccessModule.Connect()
	if err != nil {
		mainInpHandler.LogEvent(event.New(`Failed to connect to DB. Exiting.`))
		panic(fmt.Sprintf("Failed to connect to DB. Error: %s", err))
	}
	err = dbAccessModule.CheckConnection()
	if err != nil {
		mainInpHandler.LogEvent(event.New(`DB connection check is negative. Exiting.`))
		panic(fmt.Sprintf("Check DB Connection failed with error: %s", err.Error()))
	}
}

var mainInpHandler dlhandlers.InputHandler
var dbAccessorInpHandler dlhandlers.InputHandler

var outHandler dlhandlers.OutputHandler

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
	dbAccessorInpHandler, ok = inpHandlers[DB_ACCESSOR_INPUT_HANDLER_NAME]
	if !ok {
		return errors.New("DB Accessor input handler not found")
	}
	outHandler, ok = outHandlers[OUTPUT_HANDLER_NAME]
	if !ok {
		return errors.New("Output handler not found")
	}
	outHandler.SetOutputWriter(os.Stdout)
	return nil
}
