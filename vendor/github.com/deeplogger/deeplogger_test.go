package deeplogger

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const config = `{"dispatcher_name": "DISP1",
	"isOn": true,
	"inputHandlers": ["AAA", "BBB"],
	"outputHandlers": ["ZZZ", "YYY"],
	"dispatchRules": [
		{"input":"AAA", "output": "ZZZ"},
		{"input":"BBB", "output": "YYY"}
	]
}`

func TestConstructLoggerFromConfig(t *testing.T) {
	//Failed cases
	failConfigs := []string{"",
		"---",
		`{"a": "b"}`,
		`{"dispatcher_name": "D"}`,
		`{"dispatcher_name": "D", "isOn": false}`,
		`{"dispatcher_name": "D", "isOn": false, "inputHandlers": ["i"]}`,
		`{"dispatcher_name": "D", "isOn": false, "inputHandlers": ["i"], "outputHandlers": ["o"]}`,
		`{"dispatcher_name": "D", "isOn": false, "inputHandlers": ["i"], "outputHandlers": ["o"], "dispatchRules": [{}, {}]}`,
	}
	for i, fConf := range failConfigs {
		inpH, disp, outH, err := ConstructLoggerFromConfig(fConf)
		if err == nil {
			t.Errorf("In failed case %d expected error, but got no error.", i)
		}
		if inpH != nil || disp != nil || outH != nil {
			t.Errorf("ConstructLoggerFromConfig didn't return all nils for an error case.")
		}
	}

	//Working case
	inpHandl, disp, outHandl, err := ConstructLoggerFromConfig(config)
	if err != nil {
		t.Error("Failed to load config. Error: " + err.Error())
	}
	if disp.Name() != "DISP1" {
		t.Error("Dispatcher name read incorrectly.")
	}
	if disp.IsOn() != true {
		t.Error("Dispatcher should be on.")
	}
	if exists, _ := disp.HasInputHandler("AAA"); !exists {
		t.Error("Input handler not created in dispatcher.")
	}
	if exists, _ := disp.HasInputHandler("BBB"); !exists {
		t.Error("Input handler not created in dispatcher.")
	}
	if exists := disp.HasOutputHandler("YYY"); !exists {
		t.Error("Output handler not created in dispatcher.")
	}
	if exists := disp.HasOutputHandler("ZZZ"); !exists {
		t.Error("Output handler not created in dispatcher.")
	}

	if _, exists := inpHandl["AAA"]; !exists {
		t.Error("Input handler not created.")
	}
	if _, exists := inpHandl["BBB"]; !exists {
		t.Error("Input handler not created.")
	}
	if _, exists := inpHandl["ZZZ"]; exists {
		t.Error("Input handler created for wrong key.")
	}

	if _, exists := outHandl["ZZZ"]; !exists {
		t.Error("Output handler not created.")
	}
	if _, exists := outHandl["YYY"]; !exists {
		t.Error("Output handler not created.")
	}
	if _, exists := outHandl["nope"]; exists {
		t.Error("Output handler test false positive.")
	}
}

func TestConstructLoggerFromFilepath(t *testing.T) {
	const MISS_PATH = "./noexistent"
	_, err := os.Open(MISS_PATH)
	if err == nil {
		t.Errorf("File at path %s exists, but shouldn't. Please remove it.", MISS_PATH)
		return
	}
	inpH, disp, outH, err := ConstructLoggerFromFilepath(MISS_PATH)
	if err == nil {
		t.Errorf("Expected error trying to load config from noexistent file. Got no error.")
	}
	if inpH != nil || disp != nil || outH != nil {
		t.Errorf("ConstructLoggerFromConfig didn't return all nils for an error case.")
	}
}

func TestLoadFileToString(t *testing.T) {
	const MISS_PATH = "./noexistent"
	_, err := os.Open(MISS_PATH)
	if err == nil {
		t.Errorf("File at path %s exists, but shouldn't. Please remove it.", MISS_PATH)
		return
	}
	content, err := loadFileToString(MISS_PATH)
	if err == nil {
		t.Errorf("loadFileToString doesn't return an error while trying to open a nonexistent file at path: %s", MISS_PATH)
	}
	if content != "" {
		t.Errorf("Calling loadFileToString on a non-existent filepath should return empty string. Instead got: %s", content)
	}

	fContent := []byte("You broke my heart, Fredo.")
	tDir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Errorf("Failed creating temp directory: %s", err.Error())
	}
	defer os.RemoveAll(tDir)

	tFile := filepath.Join(tDir, "test_file_to_load")
	err = ioutil.WriteFile(tFile, fContent, 0200)
	if err != nil {
		t.Errorf("Failed to write temporary file: %s", err.Error())
	}
	content, err = loadFileToString(tFile)
	if err == nil {
		t.Error("Expected permission error loading file to string. Got no error.")
	}
	if content != "" {
		t.Errorf("Expected file content to be returned as: %s because of permission error, got: %s", "", content)
	}

	tFile2 := filepath.Join(tDir, "test_file_to_load2")
	err = ioutil.WriteFile(tFile2, fContent, 0666)
	if err != nil {
		t.Errorf("Failed to write temporary file: %s", err.Error())
	}
	content, err = loadFileToString(tFile2)
	if err != nil {
		t.Errorf("Failed loading file to string: %s", err.Error())
	}
	if content != string(fContent) {
		t.Errorf("Expected file content loaded to be: %s, got: %s", string(fContent), content)
	}
}
