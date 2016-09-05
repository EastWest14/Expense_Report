//Config package defines the Config data structure used to set up the application.
//Also contains the methods used to load the config.
package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Config struct {
	DBUser         string
	DBPassword     string
	DBName         string
	DeepLoggerPath string
}

const (
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_NAME     = "DB_NAME"
	DL_PATH     = "DEEP_LOGGER_PATH"
)

//Generates config from a file.
func LoadConfigFromFile(filepath string) (config *Config, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, errors.New("Failed opening config: " + err.Error())
	}
	configRaw, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("Failed reading config: " + err.Error())
	}
	config, err = LoadConfigFromString(string(configRaw))
	if err != nil {
		return nil, errors.New("Failed loading config: " + err.Error())
	}
	return config, nil
}

//Generates config from a string.
func LoadConfigFromString(configString string) (config *Config, err error) {
	conf := Config{}
	var buf interface{}
	if err := json.Unmarshal([]byte(configString), &buf); err != nil {
		return nil, err
	}
	dictRepres := buf.(map[string]interface{})
	if user, ok := dictRepres[DB_USER]; !ok {
		return nil, errors.New("DB_USER not found.")
	} else {
		conf.DBUser = user.(string)
	}
	if password, ok := dictRepres[DB_PASSWORD]; !ok {
		return nil, errors.New("DB_PASSWORD not found.")
	} else {
		conf.DBPassword = password.(string)
	}
	if dName, ok := dictRepres[DB_NAME]; !ok {
		return nil, errors.New("DB_NAME not found.")
	} else {
		conf.DBName = dName.(string)
	}
	if dlPath, ok := dictRepres[DL_PATH]; !ok {
		return nil, errors.New("DEEP_LOGGER_PATH not found.")
	} else {
		conf.DeepLoggerPath = dlPath.(string)
	}

	return &conf, nil
}
