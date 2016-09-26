package config

import (
	"fmt"
	"testing"
)

func TestLoadConfigFromString(t *testing.T) {
	const (
		FAKE_USER     = "user"
		FAKE_PASSWORD = "fake_password"
		FAKE_DB_NAME  = "name"
		FAKE_DL_PATH  = "dl_path"

		configTemplate = `{"DB_USER": "%s", "DB_PASSWORD": "%s", "DB_NAME": "%s", "DEEP_LOGGER_PATH": "%s"}`
		brokenConfig1  = ""
		brokenConfig2  = `{"DB_USER": "a"}`
	)
	configString := fmt.Sprintf(configTemplate, FAKE_USER, FAKE_PASSWORD, FAKE_DB_NAME, FAKE_DL_PATH)
	conf, err := LoadConfigFromString(configString)
	if err != nil {
		t.Errorf("Error loading config: %s", err.Error())
	}
	if conf.DBUser != FAKE_USER {
		t.Errorf("Config user loaded incorrectly. Expected \"%s\", got: \"%s\"", FAKE_USER, conf.DBUser)
	}
	if conf.DBPassword != FAKE_PASSWORD {
		t.Errorf("Config password loaded incorrectly. Expected \"%s\", got: \"%s\"", FAKE_PASSWORD, conf.DBPassword)
	}
	if conf.DBName != FAKE_DB_NAME {
		t.Errorf("Config db name loaded incorrectly. Expected \"%s\", got: \"%s\"", FAKE_DB_NAME, conf.DBName)
	}
	if conf.DeepLoggerPath != FAKE_DL_PATH {
		t.Errorf("Config deep logger path loaded incorrectly. Expected \"%s\", got: \"%s\"", FAKE_DL_PATH, conf.DeepLoggerPath)
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
