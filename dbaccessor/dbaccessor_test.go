package dbaccessor

import (
	_ "gopkg.in/DATA-DOG/go-sqlmock.v1"
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

/*
func TestCheckConnection(t *testing.T) {
	am := NewDBAccessModule()
	am.SetDBConfig()
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatal("Failed to connect to sqlmock")
	}
	err = checkConnection(db)
	if err != nil {
		t.Errorf("Check DB connection returned error: %s, expected success", err.Error())
	}
	db = nil
	err = checkConnection(db)
	if err == nil {
		t.Errorf("Checking DB connection succeeded, expected failure")
	}
}
*/
