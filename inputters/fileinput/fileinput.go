package fileinput

import (
	"errors"
	"github.com/gAssert"
	"strings"
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
	splitByNewline := strings.Split(rawInput, "\n")
	//break up by newline
	for _, stringNoNewline := range splitByNewline {
		lines := strings.SplitAfter(stringNoNewline, ";")
		//break up by semicolon

		for _, untrimmedLine := range lines {
			trimmedLine := strings.Trim(untrimmedLine, " \t")

			//A line with just whitespaces should pass
			if len(trimmedLine) == 0 {
				continue
			}

			//check if last character is a semicolon
			if trimmedLine == strings.TrimRight(trimmedLine, ";") {
				fi.emptyQueue()
				return errors.New("Last character is not a semicolon")
			}

			trimmedLineNoSemicolon := strings.TrimRight(trimmedLine, ";")

			//trim further to eliminate inner whitespaces
			finalLine := strings.Trim(trimmedLineNoSemicolon, " \t")

			//check if everything got trimmed
			if len(finalLine) == 0 {
				continue
			}

			fi.enqueueLine(finalLine)
		}
	}
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

func (fi *FileInputter) emptyQueue() {
	fi.linesQueue = []string{}
	fi.linesQueueInIndex = 0
	fi.linesQueueOutIndex = 0
}

func (fi *FileInputter) queueIsEmpty() bool {
	return len(fi.linesQueue) == 0
}
