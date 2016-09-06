package dbaccessor

import (
	"github.com/Deep-Logger/event"
	dlhandlers "github.com/Deep-Logger/handlers"
)

var DBAccessorInpHandler dlhandlers.InputHandler

type DBAccessModule interface {
	SetDBConfig(*DBConfig)
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
}

func (am *accessModule) SetDBConfig(dConfig *DBConfig) {
	DBAccessorInpHandler.LogEvent(event.New(`Setting DB config of Access Module.`))
	am.DbConfig = dConfig
}
