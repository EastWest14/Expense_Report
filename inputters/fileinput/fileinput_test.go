package fileinput

import (
	"testing"
)

//**************** Test File Inputter Setup ****************

func TestNewFileInputter(t *testing.T) {
	fInputter := NewFileInputter()
	if fInputter == nil {
		t.Error("failed to initialize file inputter")
	}
	if len(fInputter.linesQueue) != 0 {
		t.Error("Inputter line queue initialized to a non-empty value")
	}
	if fInputter.linesQueueInIndex != 0 {
		t.Error("Inputter in index is not 0")
	}
	if fInputter.linesQueueOutIndex != 0 {
		t.Error("Inputter out index is not 0")
	}
}

//**************** Test Loading Raw Data ****************

func TestLoadString(t *testing.T) {
	//Case setup
	const (
		inputEmpty            = ""
		inputSpace            = " "
		inputTab              = "\t"
		inputEmptyLine        = "\n"
		inputOneLine          = "expense 56.78;"
		inputLineAndEmptyLine = "expense 56.78;\n"
		inputOneLineAndSpaces = "  expense 56.78; \t\n"
		inputTwoLines         = "expense 56.78;\nexpense 10.00"
		inputTwoLinesAndEmpty = "expense 56.78;\n\n\nexpense 10.00"

		invalidLineNoSemicolon   = "potato"
		invalidLineJustSemicolon = ";"
		invalidTwoLinesTogether  = "expense 56.78; expense 56.78"
		invalidDoubleSemicolon   = "expense 56.78;;"
	)
	cases := []struct {
		inputString   string
		expectedLines []string
		expectedError bool
	}{
		//Pass cases
		{inputString: inputEmpty, expectedLines: nil, expectedError: false},
		{inputString: inputSpace, expectedLines: nil, expectedError: false},
		{inputString: inputTab, expectedLines: nil, expectedError: false},
		{inputString: inputEmptyLine, expectedLines: nil, expectedError: false},
		{inputString: inputOneLine, expectedLines: []string{"expense 56.78"}, expectedError: false},
		{inputString: inputLineAndEmptyLine, expectedLines: []string{"expense 56.78"}, expectedError: false},
		{inputString: inputOneLineAndSpaces, expectedLines: []string{"expense 56.78"}, expectedError: false},
		{inputString: inputTwoLines, expectedLines: []string{"expense 56.78", "expense 10.00"}, expectedError: false},
		{inputString: inputTwoLinesAndEmpty, expectedLines: []string{"expense 56.78", "expense 10.00"}, expectedError: false},

		//Fail cases
		{inputString: invalidLineNoSemicolon, expectedError: true},
		{inputString: invalidLineJustSemicolon, expectedError: true},
		{inputString: invalidTwoLinesTogether, expectedError: true},
		{inputString: invalidDoubleSemicolon, expectedError: true},
	}

	//Checking line parsing results
	for i, aCase := range cases {
		fInputter := NewFileInputter()

		err := fInputter.loadString(aCase.inputString)
		errorPresent := (err != nil)
		if errorPresent != aCase.expectedError {
			t.Errorf("Error in case: %d, error value doesn't match", i)
		}
		linesCorrectlyParsed := fInputter.compareQueueAndStrings(aCase.expectedLines)
		if !linesCorrectlyParsed {
			t.Errorf("Error in case: %d, parsed lines don't match", i)
		}
	}
}

//**************** Test Line Queue ****************

func TestEnqueue(t *testing.T) {
	cases := [][]string{
		{""},
		{"line1"},
		{"line1", "line2", "line3", "line4"},
	}
	for _, aCase := range cases {
		//Enqueue lines
		fInputter := NewFileInputter()
		if len(fInputter.linesQueue) != 0 {
			t.Error("Lines queue initialized not empty")
			break
		}
		for _, aLine := range aCase {
			fInputter.enqueueLine(aLine)
		}

		//Check the index variable has correct value
		if fInputter.linesQueueInIndex != len(aCase) {
			t.Errorf("Queue in index has value %d, expected %d", fInputter.linesQueueInIndex, len(aCase))
		}
		//Check correct lines are in the queue
		if len(fInputter.linesQueue) != len(aCase) {
			t.Errorf("Number of elements found in the queue: %d, expected: %d", len(fInputter.linesQueue), len(aCase))
			continue
		}
		for i, aLine := range aCase {
			if aLine != fInputter.linesQueue[i] {
				t.Errorf("Inserted element number %d is %s, expect %s", i, fInputter.linesQueue[i], aLine)
			}
		}
	}

}

func TestDequeue(t *testing.T) {
	fInputter := NewFileInputter()
	if len(fInputter.linesQueue) != 0 {
		t.Error("Lines queue initialized not empty")
		return
	}
	//Enqueue test lines
	queuedLines := []string{"line1", "line2", "line3"}
	for _, aLine := range queuedLines {
		fInputter.linesQueue = append(fInputter.linesQueue, aLine)
	}
	fInputter.linesQueueInIndex = len(queuedLines)

	//Test correct lines are being dequeued
	i := 0
	for {
		dequeuedLine, found := fInputter.dequeueLine()
		if !found {
			break
		}

		if dequeuedLine != queuedLines[i] {
			t.Errorf("Dequeued a line: %s, expected: %s", dequeuedLine, queuedLines[i])
		}
		i++
	}
	if i != len(queuedLines) {
		t.Errorf("Number of dequed elements: %d, expected: %d", i, len(queuedLines))
	}
	if fInputter.linesQueueOutIndex != fInputter.linesQueueInIndex {
		t.Errorf("Queue out index hasn't reached in index")
	}
	if dequeuedLine, found := fInputter.dequeueLine(); dequeuedLine != "" || found {
		t.Errorf("Incorrect dequeue of an empty queue")
	}
}

//**************** Utilities ****************

func (fi *FileInputter) compareQueueAndStrings(lines []string) (equal bool) {
	numLines := len(lines)
	for i := 0; i < numLines; i++ {
		dequeuedLine, found := fi.dequeueLine()
		if !found {
			//Number of elements unequal
			return false
		}
		if dequeuedLine != lines[i] {
			//Dequed line not equal to slice line
			return false
		}
	}
	if _, found := fi.dequeueLine(); found {
		//Queue has more elements than slice
		return false
	}
	return true
}

func TestCompareQueueAndStrings(t *testing.T) {
	cases := []struct {
		linesToEnqueue   []string
		linesToCompareTo []string
		shouldBeEqual    bool
	}{
		{linesToEnqueue: nil, linesToCompareTo: nil, shouldBeEqual: true},
		{linesToEnqueue: []string{""}, linesToCompareTo: []string{""}, shouldBeEqual: true},
		{linesToEnqueue: []string{"hello"}, linesToCompareTo: []string{"hello"}, shouldBeEqual: true},
		{linesToEnqueue: []string{"hello", "world"}, linesToCompareTo: []string{"hello", "world"}, shouldBeEqual: true},
		{linesToEnqueue: nil, linesToCompareTo: []string{"a"}, shouldBeEqual: false},
		{linesToEnqueue: []string{""}, linesToCompareTo: nil, shouldBeEqual: false},
		{linesToEnqueue: []string{""}, linesToCompareTo: []string{"a"}, shouldBeEqual: false},
		{linesToEnqueue: []string{"b"}, linesToCompareTo: []string{"a"}, shouldBeEqual: false},
		{linesToEnqueue: []string{"hello", "world"}, linesToCompareTo: []string{"hello"}, shouldBeEqual: false},
		{linesToEnqueue: []string{"hello", "world", "bye"}, linesToCompareTo: []string{"hello", "world"}, shouldBeEqual: false},
		{linesToEnqueue: []string{"hello", "world", "bye"}, linesToCompareTo: nil, shouldBeEqual: false},
	}

	for i, aCase := range cases {
		fInputter := NewFileInputter()
		for _, aLine := range aCase.linesToEnqueue {
			fInputter.enqueueLine(aLine)
		}

		comparatorResult := fInputter.compareQueueAndStrings(aCase.linesToCompareTo)
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
