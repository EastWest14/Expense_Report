package main

import (
	"Expense_Tracker/config"
	"Expense_Tracker/controller"
	"Expense_Tracker/dbaccessor"
	"Expense_Tracker/service"
	"errors"
	"fmt"
	dl "github.com/deeplogger"
	"os"
)

const (
	TEMP_PROD_CONFIG_PATH = "./conf_files/prod_config/prod_config.json"
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

//**************** Deep Logger Setup ****************

const (
	MAIN_INPUT_HANDLER_NAME        = "Main"
	SERVICE_INPUT_HANDLER_NAME     = "Service"
	CONTROLLER_INPUT_HANDLER_NAME  = "Controller"
	DB_ACCESSOR_INPUT_HANDLER_NAME = "DB_Accessor"
	OUTPUT_HANDLER_NAME            = "Out"
)

//Setting up logging input point
var mainInpHandler dl.InputHandler = dl.NewBlankInputHandler()

//Setting up logging output point
var outHandler dl.OutputHandler

func constructDeepLoggerSystem(filepath string) error {
	inpHandlers, _, outHandlers, err := dl.ConstructLoggerFromFilepath(filepath)
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

//**************** Full Application Setup ****************

func setupApplication(conf *config.Config) *controller.ControllerModule {
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

//**************** Setting up Database Module ****************

type DBAccess interface {
	Connect(driver string) error
	CheckConnection() error
}

func setupDBAccessModule(conf *config.Config) DBAccess {
	dbAccessModule := dbaccessor.NewDBAccessor()
	dbAccessModule.SetDBConfig(dbaccessor.NewDBConfig(conf.DBUser, conf.DBPassword, conf.DBName))
	if err := verifyDBConnection(dbAccessModule); err != nil {
		panic(err.Error())
	}
	return dbAccessModule
}

func verifyDBConnection(dbAccessModule DBAccess) error {
	err := dbAccessModule.Connect("postgres")
	if err != nil {
		mainInpHandler.LogMessage(`Failed to connect to DB. Exiting.`)
		return errors.New("Failed to connect to DB. Error: " + err.Error())
	}
	err = dbAccessModule.CheckConnection()
	if err != nil {
		mainInpHandler.LogMessage(`DB connection check is negative. Exiting.`)
		return errors.New("Check DB Connection failed with error: " + err.Error())
	}
	return nil
}

//**************** Setting up Service Module ****************

type Service interface {
	WaitForCommand(func(c *service.Command, err error))
}

func setupServiceModule() Service {
	return service.NewService()
}

//**************** Setting up Controller Module ****************

func setupControllerModule() *controller.ControllerModule {
	return controller.NewController()
}
