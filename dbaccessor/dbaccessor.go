package dbaccessor

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Deep-Logger/event"
	dlhandlers "github.com/Deep-Logger/handlers"
	_ "github.com/lib/pq"
)

var DBAccessorInpHandler dlhandlers.InputHandler

type DBAccessModule interface {
	SetDBConfig(*DBConfig)
	Connect() error
	CheckConnection() error
}

func NewDBAccessModule() DBAccessModule {
	if DBAccessorInpHandler == nil {
		panic("DBAccessorInpHandler is nil.")
	}
	DBAccessorInpHandler.LogEvent(event.New(`Creating new DB Access Module.`))
	return &accessModule{}
}

type DBConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
}

func NewDBConfig(dbuser, dbpassword, dbname string) *DBConfig {
	return &DBConfig{DBUser: dbuser, DBPassword: dbpassword, DBName: dbname}
}

type accessModule struct {
	DbConfig *DBConfig
	db       *sql.DB
}

func (am *accessModule) SetDBConfig(dConfig *DBConfig) {
	DBAccessorInpHandler.LogEvent(event.New(`Setting DB config of Access Module.`))
	am.DbConfig = dConfig
}

//Connect to the DB using config.
func (am *accessModule) Connect() error {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		am.DbConfig.DBUser, am.DbConfig.DBPassword, am.DbConfig.DBName)
	var err error
	am.db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		DBAccessorInpHandler.LogEvent(event.New(`Failed to connect to DB.`))
		return err
	}
	DBAccessorInpHandler.LogEvent(event.New(`Succesfully connected to DB.`))
	return nil
}

//Pings the DB to check connection.
func (am *accessModule) CheckConnection() error {
	if am.db == nil {
		DBAccessorInpHandler.LogEvent(event.New(`DB is nil.`))
		return errors.New("DB is nil")
	}
	err := am.db.Ping()
	if err != nil {
		DBAccessorInpHandler.LogEvent(event.New(`DB ping fail.`))
		return errors.New("Ping error" + err.Error())
	}
	DBAccessorInpHandler.LogEvent(event.New(`DB ping succesfull.`))
	return nil
}
