//Package gAssert is a utility that provides configurable runtime assertions.
//Users can configure the package to run a custom action on every failed assert call.
//Depending on runtime settings, a failed assert may kill the process.
//This allows to have different behavior in testing and production environments.
//Default behavior is to kill the process on a failed call to Assert.
package gAssert

import (
	"fmt"
	"log"
)

//*************** Setting Assert Package Behavior ***************

var assertsAreFatal bool = true

//SetAssertsFatal determines whether a failed Assert will kill the process.
func SetAssertsFatal(isFatal bool) {
	assertsAreFatal = isFatal
}

var actionOnAssert func(string) = func(message string) {
	fmt.Println(message)
}

//SetActionOnAssert sets the action that will happen after an assert fails.
//Default assert action prints out the message.
func SetActionOnAssert(actionFunc func(message string)) {
	actionOnAssert = actionFunc
}

//NoActionOnAssert sets the assert action to immediately return.
//This option may improve performance.
func NoActionOnAssert() {
	actionOnAssert = func(message string) {
		return
	}
}

//*************** Assert Statements ***************

//AssertSoft triggers an assert action on a false condition, then returns.
//Function doesn't kill the process regardless of the value of SetAssertsFatal.
func AssertSoft(condition bool, message string) {
	if condition {
		return
	}
	actionOnAssert(message)
}

//Assert triggers an assert action on a false condition.
//Default assert action prints out the message.
//Afterwards, the function kills the process or returns
//depending on the value of SetAssertsFatal.
//Default behavior is to kill the process.
func Assert(condition bool, message string) {
	if condition {
		return
	}
	actionOnAssert(message)
	if assertsAreFatal {
		killProcess(message)
	}
}

//AssertHard triggers an assert action on a false condition, then kills
//the process regardless of the value of SetAssertsFatal.
func AssertHard(condition bool, message string) {
	if condition {
		return
	}
	killProcess(message)
}

//*************** Killing Process ***************

//killProcessCommand exists for testing reasons only.
var killProcessCommand func(message string) = func(message string) {
	log.Fatal(message)
}

func killProcess(message string) {
	killProcessCommand(message)
}
