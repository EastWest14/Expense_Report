//fileinput package processes a text file and breaks it up into
//semicolon separated string segments. The package provides the first
//level of parsing and presents the rest of the program with a cleaner
//form of input.
package fileinput

import (
	"errors"
	"github.com/gAssert"
	"io/ioutil"
	"os"
	"strings"
)

//**************** File Inputter Setup ****************

//File Inputter loads a text file, processes it and outputs the
//broken up segments.
type FileInputter struct {
	segmentQueue         []string
	segmentQueueInIndex  int
	segmentQueueOutIndex int
}

func NewFileInputter() *FileInputter {
	return &FileInputter{segmentQueue: []string{}, segmentQueueInIndex: 0, segmentQueueOutIndex: 0}
}

//**************** Loading Raw Data ****************

//TODO: LoadFile should dump the queue in case of processing error.

//LoadFiles reads a textfile and processes it into segments represented by strings.
//Each line should contain 1+ semicolon terminated components or whitespaces only.
//If a line is not terminated by a semicolon LoadFile will return an error.
func (fi *FileInputter) LoadFile(filepath string) (loadingError error) {
	file, err := os.Open(filepath)
	if err != nil {
		return errors.New("Failed opening file: " + err.Error())
	}
	rawContent, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.New("Failed reading file: " + err.Error())
	}
	return fi.loadString(string(rawContent))
}

func (fi *FileInputter) loadString(rawInput string) (parsingError error) {
	splitByNewline := strings.Split(rawInput, "\n")
	//break up by newline
	for _, stringNoNewline := range splitByNewline {
		segments := strings.SplitAfter(stringNoNewline, ";")
		//break up by semicolon

		for _, untrimmedSegment := range segments {
			trimmedSegment := strings.Trim(untrimmedSegment, " \t")

			//A segment with just whitespaces should pass
			if len(trimmedSegment) == 0 {
				continue
			}

			//check if last character is a semicolon
			if trimmedSegment == strings.TrimRight(trimmedSegment, ";") {
				fi.emptyQueue()
				return errors.New("Last character is not a semicolon")
			}

			trimmedSegmentNoSemicolon := strings.TrimRight(trimmedSegment, ";")

			//trim further to eliminate inner whitespaces
			finalSegment := strings.Trim(trimmedSegmentNoSemicolon, " \t")

			//check if everything got trimmed
			if len(finalSegment) == 0 {
				continue
			}

			fi.enqueueSegment(finalSegment)
		}
	}
	return nil
}

//**************** Extracting Segments ****************

//ExtractSegment returns the next extracted segment.
func (fi *FileInputter) ExtractSegment() (segment string, found bool) {
	return fi.dequeueSegment()
}

//**************** Managing Segment Queue ****************

func (fi *FileInputter) dequeueSegment() (segment string, found bool) {
	gAssert.Assert(fi.segmentQueueInIndex >= fi.segmentQueueOutIndex, "FileInputter is in an inconsistent state - In index is < outer index")
	if fi.segmentQueueInIndex == fi.segmentQueueOutIndex {
		return "", false
	}
	segment = fi.segmentQueue[fi.segmentQueueOutIndex]

	fi.segmentQueueOutIndex++
	return segment, true
}

func (fi *FileInputter) enqueueSegment(segment string) {
	gAssert.Assert(fi != nil, "FileInputter is not initialized properly - segment Queue is nil")
	fi.segmentQueue = append(fi.segmentQueue, segment)
	fi.segmentQueueInIndex++
}

func (fi *FileInputter) emptyQueue() {
	fi.segmentQueue = []string{}
	fi.segmentQueueInIndex = 0
	fi.segmentQueueOutIndex = 0
}

func (fi *FileInputter) queueIsEmpty() bool {
	return len(fi.segmentQueue) == 0
}
