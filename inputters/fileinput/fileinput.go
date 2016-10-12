package fileinput

import (
	"github.com/gAssert"
)

//**************** File Inputter Setup ****************

type FileInputter struct {
	linesQueue         []string
	linesQueueInIndex  int
	linesQueueOutIndex int
}

func NewFileInputter() *FileInputter {
	return &FileInputter{linesQueue: []string{}, linesQueueInIndex: 0, linesQueueOutIndex: 0}
}

//**************** Loading Raw Data ****************

func (fi *FileInputter) LoadFile(filepath string) (loadingError error) {
	return nil
}

func (fi *FileInputter) loadString(rawInput string) (parsingError error) {
	return nil
}

//**************** Managing Line Queue ****************

func (fi *FileInputter) dequeueLine() (line string, found bool) {
	gAssert.Assert(fi.linesQueueInIndex >= fi.linesQueueOutIndex, "FileInputter is in an inconsistent state - In index is < outer index")
	if fi.linesQueueInIndex == fi.linesQueueOutIndex {
		return "", false
	}
	line = fi.linesQueue[fi.linesQueueOutIndex]

	fi.linesQueueOutIndex++
	return line, true
}

func (fi *FileInputter) enqueueLine(line string) {
	gAssert.Assert(fi != nil, "FileInputter is not initialized properly - lines Queue is nil")
	fi.linesQueue = append(fi.linesQueue, line)
	fi.linesQueueInIndex++
}
