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

func (fi *FileInputter) compareQueueAndStrings(elements []string) (equal bool) {
	return false
}

func TestCompareQueueAndStrings(t *testing.T) {

}
