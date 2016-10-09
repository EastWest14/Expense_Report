package config

import (
	"fmt"
	"testing"
)

func TestLoadConfigFromString(t *testing.T) {
	const (
		MOCK_USER     = "user"
		MOCK_PASSWORD = "fake_password"
		MOCK_DB_NAME  = "name"
		MOCK_DL_PATH  = "dl_path"

		configTemplate = `{"DB_USER": "%s", "DB_PASSWORD": "%s", "DB_NAME": "%s", "DEEP_LOGGER_PATH": "%s", "ASSERTS_ARE_FATAL": true}`
		brokenConfig1  = ""
		brokenConfig2  = `{"DB_USER": "a"}`
	)
	configString := fmt.Sprintf(configTemplate, MOCK_USER, MOCK_PASSWORD, MOCK_DB_NAME, MOCK_DL_PATH)
	conf, err := LoadConfigFromString(configString)
	if err != nil {
		t.Errorf("Error loading config: %s", err.Error())
	}
	if conf.DBUser != MOCK_USER {
		t.Errorf("Config user loaded incorrectly. Expected \"%s\", got: \"%s\"", MOCK_USER, conf.DBUser)
	}
	if conf.DBPassword != MOCK_PASSWORD {
		t.Errorf("Config password loaded incorrectly. Expected \"%s\", got: \"%s\"", MOCK_PASSWORD, conf.DBPassword)
	}
	if conf.DBName != MOCK_DB_NAME {
		t.Errorf("Config db name loaded incorrectly. Expected \"%s\", got: \"%s\"", MOCK_DB_NAME, conf.DBName)
	}
	if conf.DeepLoggerPath != MOCK_DL_PATH {
		t.Errorf("Config deep logger path loaded incorrectly. Expected \"%s\", got: \"%s\"", MOCK_DL_PATH, conf.DeepLoggerPath)
	}
	if !conf.AssertsAreFatal {
		t.Errorf("Config deep logger path loaded incorrectly. Expected ASSERTS_ARE_FATAL to be on")
	}

	_, err = LoadConfigFromString(brokenConfig1)
	if err == nil {
		t.Error("Broken config1 doesn't generate error")
	}
	_, err = LoadConfigFromString(brokenConfig2)
	if err == nil {
		t.Error("Broken config2 doesn't generate error")
	}
}
