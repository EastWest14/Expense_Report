package main

import (
	"Expense_Tracker/config"
	"Expense_Tracker/controller"
	"Expense_Tracker/dbaccessor"
	"Expense_Tracker/service"
	"errors"
	"fmt"
	dl "github.com/deeplogger"
	"io/ioutil"
	"os"
)

const (
	TEMP_PROD_CONFIG_PATH = "./conf_files/prod_config/prod_config.json"

	MAIN_INPUT_HANDLER_NAME        = "Main"
	SERVICE_INPUT_HANDLER_NAME     = "Service"
	CONTROLLER_INPUT_HANDLER_NAME  = "Controller"
	DB_ACCESSOR_INPUT_HANDLER_NAME = "DB_Accessor"
	OUTPUT_HANDLER_NAME            = "Out"
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
	fmt.Println()
	mainInpHandler.LogMessage(`Deep Logger succesfully configured.`)
	contr := setupApplication(config)
	mainInpHandler.LogMessage(`Transfering control to controller.`)
	contr.AcceptControl()
	fmt.Println()
}

func setupApplication(conf *config.Config) controller.Controller {
	dalModule := setupDBAccessModule(conf)
	mainInpHandler.LogMessage(`DB Access Module succesfully configured. Connection to DB established.`)
	serv := setupServiceModule()
	mainInpHandler.LogMessage(`Service Module succesfully configured.`)
	contr := setupControllerModule()
	contr.SetService(serv)
	contr.SetDAL(dalModule)
	mainInpHandler.LogMessage(`Controller Module succesfully configured.`)
	return contr
}

func setupDBAccessModule(conf *config.Config) dbaccessor.DBAccess {
	dbAccessModule := dbaccessor.NewDBAccessor()
	dbAccessModule.SetDBConfig(dbaccessor.NewDBConfig(conf.DBUser, conf.DBPassword, conf.DBName))
	err := dbAccessModule.Connect("postgres")
	if err != nil {
		mainInpHandler.LogMessage(`Failed to connect to DB. Exiting.`)
		panic(fmt.Sprintf("Failed to connect to DB. Error: %s", err))
	}
	err = dbAccessModule.CheckConnection()
	if err != nil {
		mainInpHandler.LogMessage(`DB connection check is negative. Exiting.`)
		panic(fmt.Sprintf("Check DB Connection failed with error: %s", err.Error()))
	}
	return dbAccessModule
}

func setupServiceModule() service.Service {
	return service.NewService()
}

func setupControllerModule() controller.Controller {
	contr := controller.NewController()
	return contr
}

var mainInpHandler dl.InputHandler

var outHandler dl.OutputHandler

func constructDeepLoggerSystem(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return errors.New("Failed opening DeepLogger config: " + err.Error())
	}
	configRaw, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.New("Failed reading DeepLogger config: " + err.Error())
	}
	inpHandlers, _, outHandlers, err := dl.ConstructLoggerFromConfig(string(configRaw))
	if err != nil {
		return err
	}
	var ok bool
	mainInpHandler, ok = inpHandlers[MAIN_INPUT_HANDLER_NAME]
	if !ok {
		return errors.New("Main input handler not found")
	}
	service.ServInpHandler, ok = inpHandlers[SERVICE_INPUT_HANDLER_NAME]
	if !ok {
		return errors.New("Service handler not found")
	}
	dbaccessor.DBAccessorInpHandler, ok = inpHandlers[DB_ACCESSOR_INPUT_HANDLER_NAME]
	if !ok {
		return errors.New("DB Accessor input handler not found")
	}
	controller.ContrInpHandler, ok = inpHandlers[CONTROLLER_INPUT_HANDLER_NAME]
	if !ok {
		return errors.New("Controller handler not found")
	}
	outHandler, ok = outHandlers[OUTPUT_HANDLER_NAME]
	if !ok {
		return errors.New("Output handler not found")
	}
	outHandler.SetOutputWriter(os.Stdout)
	return nil
}
