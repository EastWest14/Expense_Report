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
	DBUser          string
	DBPassword      string
	DBName          string
	DeepLoggerPath  string
	AssertsAreFatal bool
}

const (
	DB_USER           = "DB_USER"
	DB_PASSWORD       = "DB_PASSWORD"
	DB_NAME           = "DB_NAME"
	DL_PATH           = "DEEP_LOGGER_PATH"
	ASSERTS_ARE_FATAL = "ASSERTS_ARE_FATAL"
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
	//Unmarshalling the JSON
	conf := Config{}
	var buf interface{}
	if err := json.Unmarshal([]byte(configString), &buf); err != nil {
		return nil, err
	}
	dictRepres := buf.(map[string]interface{})

	//Loading DB settings
	if user, ok := dictRepres[DB_USER]; !ok {
		return nil, errors.New("DB_USER setting not found.")
	} else {
		conf.DBUser = user.(string)
	}
	if password, ok := dictRepres[DB_PASSWORD]; !ok {
		return nil, errors.New("DB_PASSWORD setting not found.")
	} else {
		conf.DBPassword = password.(string)
	}
	if dName, ok := dictRepres[DB_NAME]; !ok {
		return nil, errors.New("DB_NAME setting not found.")
	} else {
		conf.DBName = dName.(string)
	}

	//Loading Deep Logger settings
	if dlPath, ok := dictRepres[DL_PATH]; !ok {
		return nil, errors.New("DEEP_LOGGER_PATH setting not found.")
	} else {
		conf.DeepLoggerPath = dlPath.(string)
	}

	//Loading assert settings
	if assertsAreFatal, ok := dictRepres[ASSERTS_ARE_FATAL]; !ok {
		return nil, errors.New("ASSERTS_ARE_FATAL setting not found.")
	} else {
		conf.AssertsAreFatal = assertsAreFatal.(bool)
	}

	return &conf, nil
}
