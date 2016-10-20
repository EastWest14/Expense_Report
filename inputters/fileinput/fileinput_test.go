package fileinput

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

//**************** Test File Inputter Setup ****************

func TestNewFileInputter(t *testing.T) {
	fInputter := NewFileInputter()
	if fInputter == nil {
		t.Error("failed to initialize file inputter")
	}
	if len(fInputter.segmentQueue) != 0 {
		t.Error("Inputter segment queue initialized to a non-empty value")
	}
	if fInputter.segmentQueueInIndex != 0 {
		t.Error("Inputter in index is not 0")
	}
	if fInputter.segmentQueueOutIndex != 0 {
		t.Error("Inputter out index is not 0")
	}
}

//**************** Extracting Segments ****************

func TestExtractSegments(t *testing.T) {
	fInputter := NewFileInputter()
	segment, found := fInputter.ExtractSegment()
	if found || segment != "" {
		t.Error("Finding elements in a newly initialized empty queue")
	}
	segments := []string{"I", "will", "be", "back !"}
	for _, aSegment := range segments {
		fInputter.enqueueSegment(aSegment)
	}

	for i := 0; i < len(segments); i++ {
		segment, found := fInputter.ExtractSegment()
		if !found {
			t.Errorf("Error in case %d. Expected segment to be found, not found", i)
		}
		if segment != segments[i] {
			t.Errorf("Error in case %d. Expected %s, got %s", i, segments[i], segment)
		}
	}
	segment, found = fInputter.ExtractSegment()
	if found || segment != "" {
		t.Error("Finding elements in a queue that is suppose to be empty")
	}
	segment, found = fInputter.ExtractSegment()
	if found || segment != "" {
		t.Error("Finding elements in a queue that is suppose to be empty on a second attempt")
	}

}

//**************** Test Loading Raw Data ****************

func TestLoadFile(t *testing.T) {
	const MISS_PATH = "./noexistent_file"
	_, err := os.Open(MISS_PATH)
	if err == nil {
		t.Errorf("Setup for test incorrect. File at path %s exists, but shouldn't. Please remove it.", MISS_PATH)
		return
	}

	fInputter := NewFileInputter()
	err = fInputter.LoadFile(MISS_PATH)
	if err == nil {
		t.Errorf("loadFiledoesn't return an error while trying to open a nonexistent file at path: %s", MISS_PATH)
	}
	if !fInputter.queueIsEmpty() {
		t.Errorf("Calling LoadFile on a non-existent filepath shouldn't add elements to the queue")
	}

	//
	fileContent := []byte("You broke my heart, Fredo;  Fredo runs away ;")
	tempDir, err := ioutil.TempDir("", "test_directory")
	if err != nil {
		t.Errorf("Test setup error. Failed creating temp directory: %s", err.Error())
	}
	defer os.RemoveAll(tempDir)
	fileNoPermission := filepath.Join(tempDir, "test_file_to_load")
	//LoadFile doesn't have permission to open the file
	err = ioutil.WriteFile(fileNoPermission, fileContent, 0200)
	if err != nil {
		t.Errorf("Test setup error. Failed to write temporary file: %s", err.Error())
	}
	err = fInputter.LoadFile(fileNoPermission)
	if err == nil {
		t.Error("Expected permission error loading file to string. Got no error")
	}
	if !fInputter.queueIsEmpty() {
		t.Errorf("Calling LoadFile on a file that progran can't access shouldn't add elements to the queue")
	}

	//
	fileWithEmptyContent := filepath.Join(tempDir, "test_file_to_load2")
	err = ioutil.WriteFile(fileWithEmptyContent, []byte(""), 0666)
	if err != nil {
		t.Errorf("Test setup error. Failed to write temporary file: %s", err.Error())
	}
	err = fInputter.LoadFile(fileWithEmptyContent)
	if err != nil {
		t.Errorf("Failed loading file to string: %s", err.Error())
	}
	if !fInputter.compareQueueAndStrings(nil) {
		t.Errorf("LoadFile on an empty file shouldn't add elements to the queue")
	}

	//
	fileWithValidContent := filepath.Join(tempDir, "test_file_to_load3")
	err = ioutil.WriteFile(fileWithValidContent, fileContent, 0666)
	if err != nil {
		t.Errorf("Test setup error. Failed to write temporary file: %s", err.Error())
	}
	err = fInputter.LoadFile(fileWithValidContent)
	if err != nil {
		t.Errorf("Failed loading file to string: %s", err.Error())
	}
	if !fInputter.compareQueueAndStrings([]string{"You broke my heart, Fredo", "Fredo runs away"}) {
		t.Errorf("Expected file to load [%s] and [%s], didn't get the right result", "You broke my heart, Fredo", "Fredo runs away")
	}
}

func TestLoadString(t *testing.T) {
	//Case setup
	const (
		inputEmpty             = ""
		inputSpace             = " "
		inputTab               = "\t"
		inputEmptyLine         = "\n"
		inputJustSemicolon     = ";"
		inputOneLine           = "expense 56.78;"
		inputOneLineInnerSpace = "expense 56.78  \t;"
		inputOtherText         = "any phrase;"
		inputLineAndEmptyLine  = "expense 56.78;\n"
		inputOneLineAndSpaces  = "  expense 56.78; \t\n"
		inputTwoLines          = "expense 56.78;\nexpense 10.00;"
		inputTwoLinesAndEmpty  = "expense 56.78;\n\n\nexpense 10.00;"
		inputTwoLinesTogether  = "expense 56.78; expense 56.78;"
		inputDoubleSemicolon   = "expense 30;;"

		invalidLineNoSemicolon     = "potato"
		invalidTwoLinesNoSemicolon = "expense 56.78; no semicolon"
	)
	inputVeryLongLine := []byte{}
	expectedVeryLongSlice := []string{}
	for i := 0; i < 10000; i++ {
		inputVeryLongLine = append(inputVeryLongLine, []byte("potato; ")...)
		expectedVeryLongSlice = append(expectedVeryLongSlice, "potato")
	}

	cases := []struct {
		inputString      string
		expectedSegments []string
		expectedError    bool
	}{
		//Pass cases
		{inputString: inputEmpty, expectedSegments: nil, expectedError: false},
		{inputString: inputSpace, expectedSegments: nil, expectedError: false},
		{inputString: inputTab, expectedSegments: nil, expectedError: false},
		{inputString: inputEmptyLine, expectedSegments: nil, expectedError: false},
		{inputString: inputJustSemicolon, expectedSegments: nil, expectedError: false},
		{inputString: inputOneLine, expectedSegments: []string{"expense 56.78"}, expectedError: false},
		{inputString: inputOneLineInnerSpace, expectedSegments: []string{"expense 56.78"}, expectedError: false},
		{inputString: inputOtherText, expectedSegments: []string{"any phrase"}, expectedError: false},
		{inputString: inputLineAndEmptyLine, expectedSegments: []string{"expense 56.78"}, expectedError: false},
		{inputString: inputOneLineAndSpaces, expectedSegments: []string{"expense 56.78"}, expectedError: false},
		{inputString: inputTwoLines, expectedSegments: []string{"expense 56.78", "expense 10.00"}, expectedError: false},
		{inputString: inputTwoLinesAndEmpty, expectedSegments: []string{"expense 56.78", "expense 10.00"}, expectedError: false},
		{inputString: inputTwoLinesTogether, expectedSegments: []string{"expense 56.78", "expense 56.78"}, expectedError: false},
		{inputString: inputDoubleSemicolon, expectedSegments: []string{"expense 30"}, expectedError: false},
		{inputString: string(inputVeryLongLine), expectedSegments: expectedVeryLongSlice, expectedError: false},

		//Fail cases
		{inputString: invalidLineNoSemicolon, expectedError: true},
		{inputString: invalidTwoLinesNoSemicolon, expectedError: true},
	}

	//Checking segment parsing results
	for i, aCase := range cases {
		fInputter := NewFileInputter()

		err := fInputter.loadString(aCase.inputString)
		if err != nil {
			if !aCase.expectedError {
				t.Errorf("Error in case: %d. Expected no error, got: %s", i, err.Error())
			}
		}
		if err == nil {
			if aCase.expectedError {
				t.Errorf("Error in case: %d. Expected error, but got none", i)
			}
		}
		segmentsCorrectlyParsed := fInputter.compareQueueAndStrings(aCase.expectedSegments)
		if !segmentsCorrectlyParsed {
			t.Errorf("Error in case: %d. Parsedsegments don't match. Expected: %v", i, aCase.expectedSegments)
		}
	}
}

//**************** Test Segment Queue ****************

func TestEnqueue(t *testing.T) {
	cases := [][]string{
		{""},
		{"segment1"},
		{"segment1", "segment2", "segment3", "segment4"},
	}
	for _, aCase := range cases {
		//Enqueue segments
		fInputter := NewFileInputter()
		if len(fInputter.segmentQueue) != 0 {
			t.Error("Segments queue initialized not empty")
			break
		}
		for _, aSegment := range aCase {
			fInputter.enqueueSegment(aSegment)
		}

		//Check the index variable has correct value
		if fInputter.segmentQueueInIndex != len(aCase) {
			t.Errorf("Queue in index has value %d, expected %d", fInputter.segmentQueueInIndex, len(aCase))
		}
		//Check correct segments are in the queue
		if len(fInputter.segmentQueue) != len(aCase) {
			t.Errorf("Number of elements found in the queue: %d, expected: %d", len(fInputter.segmentQueue), len(aCase))
			continue
		}
		for i, aSegment := range aCase {
			if aSegment != fInputter.segmentQueue[i] {
				t.Errorf("Inserted element number %d is %s, expect %s", i, fInputter.segmentQueue[i], aSegment)
			}
		}
	}

}

func TestDequeue(t *testing.T) {
	fInputter := NewFileInputter()
	if len(fInputter.segmentQueue) != 0 {
		t.Error("Segments queue initialized not empty")
		return
	}
	//Enqueue test segments
	queuedSegments := []string{"segment1", "segment2", "segment3"}
	for _, aSegment := range queuedSegments {
		fInputter.segmentQueue = append(fInputter.segmentQueue, aSegment)
	}
	fInputter.segmentQueueInIndex = len(queuedSegments)

	//Test correct segments are being dequeued
	i := 0
	for {
		dequeuedSegment, found := fInputter.dequeueSegment()
		if !found {
			break
		}

		if dequeuedSegment != queuedSegments[i] {
			t.Errorf("Dequeued a segment: %s, expected: %s", dequeuedSegment, queuedSegments[i])
		}
		i++
	}
	if i != len(queuedSegments) {
		t.Errorf("Number of dequed elements: %d, expected: %d", i, len(queuedSegments))
	}
	if fInputter.segmentQueueOutIndex != fInputter.segmentQueueInIndex {
		t.Errorf("Queue out index hasn't reached in index")
	}
	if dequeuedSegment, found := fInputter.dequeueSegment(); dequeuedSegment != "" || found {
		t.Errorf("Incorrect dequeue of an empty queue")
	}
}

func TestEmptyQueue(t *testing.T) {
	fInputter := NewFileInputter()
	queuedSegments := []string{"segment1", "segment2", "segment3"}
	for _, aSegment := range queuedSegments {
		fInputter.segmentQueue = append(fInputter.segmentQueue, aSegment)
	}
	fInputter.segmentQueueInIndex = len(queuedSegments)

	fInputter.emptyQueue()
	if _, found := fInputter.dequeueSegment(); found {
		t.Errorf("Didn't empty the queue")
	}

	//Test on already empty queue
	fInputter = NewFileInputter()
	fInputter.emptyQueue()
	if !fInputter.queueIsEmpty() {
		t.Errorf("Error when trying to empty an already empty queue")
	}
}

func TestIsEmpty(t *testing.T) {
	fInputter := NewFileInputter()
	queuedSegments := []string{"segment1", "segment2", "segment3"}
	for _, aSegment := range queuedSegments {
		fInputter.segmentQueue = append(fInputter.segmentQueue, aSegment)
	}
	fInputter.segmentQueueInIndex = len(queuedSegments)

	isEmptyResponse := fInputter.queueIsEmpty()
	if isEmptyResponse {
		t.Error("A non-empty queue reported as empty")
	}

	fInputter = NewFileInputter()
	isEmptyResponse = fInputter.queueIsEmpty()
	if !isEmptyResponse {
		t.Error("An empty queue reported as non-empty")
	}
}

//**************** Utilities ****************

func (fi *FileInputter) compareQueueAndStrings(segments []string) (equal bool) {
	numSegments := len(segments)
	for i := 0; i < numSegments; i++ {
		dequeuedSegment, found := fi.dequeueSegment()
		if !found {
			//Number of elements unequal
			return false
		}
		if dequeuedSegment != segments[i] {
			//Dequed segment not equal to slice segment
			return false
		}
	}
	if _, found := fi.dequeueSegment(); found {
		//Queue has more elements than slice
		return false
	}
	return true
}

func TestCompareQueueAndStrings(t *testing.T) {
	cases := []struct {
		segmentsToEnqueue   []string
		segmentsToCompareTo []string
		shouldBeEqual       bool
	}{
		{segmentsToEnqueue: nil, segmentsToCompareTo: nil, shouldBeEqual: true},
		{segmentsToEnqueue: []string{""}, segmentsToCompareTo: []string{""}, shouldBeEqual: true},
		{segmentsToEnqueue: []string{"hello"}, segmentsToCompareTo: []string{"hello"}, shouldBeEqual: true},
		{segmentsToEnqueue: []string{"hello", "world"}, segmentsToCompareTo: []string{"hello", "world"}, shouldBeEqual: true},
		{segmentsToEnqueue: nil, segmentsToCompareTo: []string{"a"}, shouldBeEqual: false},
		{segmentsToEnqueue: []string{""}, segmentsToCompareTo: nil, shouldBeEqual: false},
		{segmentsToEnqueue: []string{""}, segmentsToCompareTo: []string{"a"}, shouldBeEqual: false},
		{segmentsToEnqueue: []string{"b"}, segmentsToCompareTo: []string{"a"}, shouldBeEqual: false},
		{segmentsToEnqueue: []string{"hello", "world"}, segmentsToCompareTo: []string{"hello"}, shouldBeEqual: false},
		{segmentsToEnqueue: []string{"hello", "world", "bye"}, segmentsToCompareTo: []string{"hello", "world"}, shouldBeEqual: false},
		{segmentsToEnqueue: []string{"hello", "world", "bye"}, segmentsToCompareTo: nil, shouldBeEqual: false},
	}

	for i, aCase := range cases {
		fInputter := NewFileInputter()
		for _, aSegment := range aCase.segmentsToEnqueue {
			fInputter.enqueueSegment(aSegment)
		}

		comparatorResult := fInputter.compareQueueAndStrings(aCase.segmentsToCompareTo)
		if aCase.shouldBeEqual {
			if !comparatorResult {
				t.Errorf("CompareQueueAndStrings should return true, but returns false. Case number: %d", i)
			}
		} else {
			if comparatorResult {
				t.Errorf("CompareQueueAndStrings should return false, but returns true. Case number: %d", i)
			}
		}
	}
}
