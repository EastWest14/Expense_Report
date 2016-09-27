package dbaccessor

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestConnect(t *testing.T) {
	dbAccessor := &accessModule{}
	dbAccessor.DbConfig = &DBConfig{DBUser: "mock_user", DBPassword: "mock_password", DBName: "mock_name"}
	err := dbAccessor.Connect("sqlmock")
	if err != nil {
		t.Error(err.Error())
	}
	dbAccessor.db.Close()

	err = dbAccessor.Connect("Invalid_driver")
	if err == nil {
		t.Error("Connection should have failed, but didn't.")
	}
}

func TestCheckConnection(t *testing.T) {
	am := &accessModule{}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatal("Failed to connect to sqlmock")
	}
	am.db = db
	err = am.CheckConnection()
	if err != nil {
		t.Errorf("Check DB connection returned error: %s, expected success", err.Error())
	}

	am.db.Close()
	err = am.CheckConnection()
	if err == nil {
		t.Error("Checking closed DB connection succeeded, expected failure")
	}

	am.db = nil
	err = am.CheckConnection()
	if err == nil {
		t.Error("Checking DB connection succeeded, expected failure")
	}
}

func TestNewDBConfig(t *testing.T) {
	const (
		USER     = "User"
		PASSWORD = "PS"
		DB_NAME  = "DB_NAME"
	)
	conf := NewDBConfig(USER, PASSWORD, DB_NAME)
	if conf.DBUser != USER {
		t.Errorf("Expecting user %s, got %s", USER, conf.DBUser)
	}
	if conf.DBPassword != PASSWORD {
		t.Errorf("Expecting password %s, got %s", PASSWORD, conf.DBPassword)
	}
	if conf.DBName != DB_NAME {
		t.Errorf("Expecting db name %s, got %s", DB_NAME, conf.DBName)
	}
}

func TestNewDBAccessor(t *testing.T) {
	accessor := NewDBAccessor()
	if accessor == nil {
		t.Error("Failed to initialize DB accessor.")
	}
}

func TestSetDBConfig(t *testing.T) {
	am := &accessModule{}
	dConf := &DBConfig{DBUser: "mock_user", DBPassword: "mock_password", DBName: "mock_name"}
	am.SetDBConfig(dConf)
	if am.DbConfig == nil {
		t.Error("Failed to set db config")
	}
}
