package dbaccessor

import (
	"database/sql"
	"errors"
	"fmt"
	dl "github.com/deeplogger"
	_ "github.com/lib/pq"
)

var DBAccessorInpHandler dl.InputHandler = dl.NewBlankInputHandler()

func NewDBAccessor() *AccessModule {
	DBAccessorInpHandler.LogMessage(`Creating new DB Access Module.`)
	return &AccessModule{}
}

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
}

func NewDBConfig(dbuser, dbpassword, dbname string) *DBConfig {
	return &DBConfig{DBUser: dbuser, DBPassword: dbpassword, DBName: dbname}
}

type AccessModule struct {
	DbConfig *DBConfig
	db       *sql.DB
}

func (am *AccessModule) SetDBConfig(dConfig *DBConfig) {
	DBAccessorInpHandler.LogMessage(`Setting DB config of Access Module.`)
	am.DbConfig = dConfig
}

//Connect to the DB using config.
func (am *AccessModule) Connect(driver string) error {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		am.DbConfig.DBUser, am.DbConfig.DBPassword, am.DbConfig.DBName)
	var err error
	am.db, err = sql.Open(driver, dbinfo)
	if err != nil {
		DBAccessorInpHandler.LogMessage(`Failed to connect to DB.`)
		return err
	}
	DBAccessorInpHandler.LogMessage(`Succesfully connected to DB.`)
	return nil
}

//Pings the DB to check connection.
func (am *AccessModule) CheckConnection() error {
	if am.db == nil {
		DBAccessorInpHandler.LogMessage(`DB is nil.`)
		return errors.New("DB is nil")
	}
	err := am.db.Ping()
	if err != nil {
		DBAccessorInpHandler.LogMessage(`DB ping fail.`)
		return errors.New("Ping error" + err.Error())
	}
	DBAccessorInpHandler.LogMessage(`DB ping succesfull.`)
	return nil
}
