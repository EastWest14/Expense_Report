package dbaccessor

import (
	_ "testing"
)

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
