package gAssert

import (
	"testing"
)

//*************** Example ***************

func Example() {
	someString := "Hello, world!"
	//Assert will pass
	Assert(len(someString) > 0, "String length is 0")

	//Assert will fail
	AssertSoft(2*2 == 5, "Two times two is not five")
	//Output: Two times two is not five
}

//*************** Mock Structures ***************

type mockProcessKiller struct {
	wasHit bool
}

func (pk *mockProcessKiller) mockKillProcessCommand(message string) {
	pk.wasHit = true
}

//*************** Test Setting Assert Package Behavior ***************

func TestSetAssertsFatal(t *testing.T) {
	//Re-setting the variable after the test
	startingAssertsAreFatal := assertsAreFatal
	defer func() {
		assertsAreFatal = startingAssertsAreFatal
	}()

	//Test making asserts fatal
	assertsAreFatal = false
	SetAssertsFatal(true)
	if !assertsAreFatal {
		t.Error("Asserts were suppose to be set as fatal")
	}

	//Test making asserts not fatal
	assertsAreFatal = true
	SetAssertsFatal(false)
	if assertsAreFatal {
		t.Error("Assers were suppose to be set as not fatal")
	}
}

func TestSetActionOnAssert(t *testing.T) {
	//Re-setting the variable after the test
	startingActionOnAssert := actionOnAssert
	defer func() {
		actionOnAssert = startingActionOnAssert
	}()

	//Creating a mock function
	wasHit := false
	checkHitFunction := func(message string) {
		wasHit = true
	}

	//Check if SetActionOnAssert really modifies the assert action
	actionOnAssert = func(message string) {
		return
	}
	SetActionOnAssert(checkHitFunction)
	actionOnAssert("")
	//Check correct function was called
	if !wasHit {
		t.Error("Failed to set function on assert")
	}
}

func TestNoActionOnAssert(t *testing.T) {
	//Re-setting the variable after the test
	startingActionOnAssert := actionOnAssert
	defer func() {
		actionOnAssert = startingActionOnAssert
	}()

	//Creating a mock function
	wasHit := false
	checkHitFunction := func(message string) {
		wasHit = true
	}

	//Set the assert action
	actionOnAssert = checkHitFunction
	//Un-set the assert action
	NoActionOnAssert()

	//Check if the assert action was un-set
	actionOnAssert("")
	if wasHit {
		t.Error("Assert action executed despite being disabled")
	}
}

//*************** Test Assert Statements ***************

func TestAssertSoft(t *testing.T) {
	//Re-setting the variable after the test
	startingActionOnAssert := actionOnAssert
	defer func() {
		actionOnAssert = startingActionOnAssert
	}()

	//Creating a mock function
	wasHit := false
	checkHitFunction := func(message string) {
		wasHit = true
	}

	//Test assert on true
	actionOnAssert = checkHitFunction
	AssertSoft(true, "")
	if wasHit {
		t.Error("AssertSoft fails on a true condition")
	}

	//Test assert on false
	wasHit = false
	AssertSoft(false, "")
	if !wasHit {
		t.Error("AssertSoft passes on a false condition")
	}
}

func TestAssert(t *testing.T) {
	//Re-setting the variable after the test
	startingKillProcessCommand := killProcessCommand
	defer func() {
		killProcessCommand = startingKillProcessCommand
	}()
	//Re-setting the variable after the test
	startingActionOnAssert := actionOnAssert
	defer func() {
		actionOnAssert = startingActionOnAssert
	}()
	//Re-setting the variable after the test
	startingAssertsAreFatal := assertsAreFatal
	defer func() {
		assertsAreFatal = startingAssertsAreFatal
	}()

	//Mocking the killProcess function
	mockPK := mockProcessKiller{}
	killProcessCommand = mockPK.mockKillProcessCommand
	//Creating a mock function
	var assertActionCalled bool
	mockAssertAction := func(message string) {
		assertActionCalled = true
	}
	actionOnAssert = mockAssertAction

	//Assert should act like SoftAssert or HardAssert depending on the value of assertsAreFatal
	cases := []struct {
		assertConditionShouldPass bool
		assertsAreFatal           bool
		expectActionOnAssertCall  bool
		exceptKillProcessCall     bool
	}{
		{true, true, false, false},
		{true, false, false, false},
		{false, true, true, true},
		{false, false, true, false},
	}
	for i, aCase := range cases {
		//Pre-set mock variables before every case
		assertActionCalled = false
		mockPK.wasHit = false

		assertsAreFatal = aCase.assertsAreFatal
		Assert(aCase.assertConditionShouldPass, "")
		//Check expectations
		if assertActionCalled != aCase.expectActionOnAssertCall {
			t.Errorf("Assert action called/not called incorrectly in case %d", i)
		}
		if mockPK.wasHit != aCase.exceptKillProcessCall {
			t.Errorf("Kill process called/not called incorrectly in case %d", i)
		}
	}
}

func TestAssertHard(t *testing.T) {
	//Re-setting the variable after the test
	startingKillProcessCommand := killProcessCommand
	defer func() {
		killProcessCommand = startingKillProcessCommand
	}()

	//Mocking the killProcess function
	mockPK := mockProcessKiller{wasHit: false}
	killProcessCommand = mockPK.mockKillProcessCommand
	//Test assert on true
	AssertHard(true, "")
	if mockPK.wasHit {
		t.Error("AssertHard fails on a true condition")
	}

	//Test assert on false
	mockPK = mockProcessKiller{wasHit: false}
	AssertHard(false, "")
	if !mockPK.wasHit {
		t.Error("AssertHard passes on a false condition")
	}
}

//*************** Test Process Killing ***************

func TestKillProcess(t *testing.T) {
	//Re-setting the variable after the test
	startingKillProcessCommand := killProcessCommand
	defer func() {
		killProcessCommand = startingKillProcessCommand
	}()

	//Mocking the process killProcessCommand
	mockPK := mockProcessKiller{wasHit: false}
	killProcessCommand = mockPK.mockKillProcessCommand

	killProcess("")
	if !mockPK.wasHit {
		t.Error("Kill Process not called")
	}
}
