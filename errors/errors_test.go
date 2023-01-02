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
