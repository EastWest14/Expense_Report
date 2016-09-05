package config

import (
	"fmt"
	"testing"
)

func TestLoadDbConfig(t *testing.T) {
	const (
		FAKE_USER     = "user"
		FAKE_PASSWORD = "fake_password"
		FAKE_DB_NAME  = "name"

		configTemplate = `{"DB_USER": "%s", "DB_PASSWORD": "%s", "DB_NAME": "%s"}`
	)
	configString := fmt.Sprintf(configTemplate, FAKE_USER, FAKE_PASSWORD, FAKE_DB_NAME)
	conf, err := LoadConfigFromString(configString)
	if err != nil {
		t.Errorf("Error loading config: %s", err.Error())
	}
	if conf.DBUser != "user" {
		t.Errorf("Config user loaded incorrectly. Expected \"user\", got: %s", conf.DBUser)
	}
	if conf.DBPassword != "fake_password" {
		t.Errorf("Config password loaded incorrectly. Expected \"fake_password\", got: %s", conf.DBPassword)
	}
	if conf.DBName != "name" {
		t.Errorf("Config db name loaded incorrectly. Expected \"name\", got: %s", conf.DBName)
	}
}
