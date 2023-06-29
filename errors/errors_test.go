package errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestIfAnErrorOccursCallsLogFatal(t *testing.T) {
	var ok bool

	logFatalf = func(format string, v ...any) {
		pc, _, _, _ := runtime.Caller(1)
		fmt.Println(runtime.FuncForPC(pc).Name())
		ok = true
	}

	IfAnErrorOccursCallsLogFatal(errors.New("error occurs"), "errors here")

	assert.Equal(t, true, ok)
}

func TestIfAnErrorOccursCallsLogPrintln(t *testing.T) {
	var ok bool

	logPrintf = func(format string, v ...any) {
		pc, _, _, _ := runtime.Caller(1)
		fmt.Println(runtime.FuncForPC(pc).Name())
		ok = true
	}

	IfAnErrorOccursCallsLogPrint(errors.New("error occurs"), "errors here")

	assert.Equal(t, true, ok)
}

func TestIfAnErrorOccursIgnore(t *testing.T) {
	// Test case: No error occurs
	var err error = nil
	IfAnErrorOccursIgnore(err)
	// No need to perform any assertions, just ensure that the function doesn't cause any errors

	// Test case: An error occurs
	err = errors.New("an error occurred")
	IfAnErrorOccursIgnore(err)
	// No need to perform any assertions, just ensure that the function doesn't cause any errors
}
