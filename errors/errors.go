package errors

import (
	"log"
	"runtime"
)

var logFatalf = log.Fatalf

func IfAnErrorOccursCallsLogFatal(err error, message string) {
	if err != nil {
		pc, _, _, _ := runtime.Caller(1)
		nameOfTheFunction := runtime.FuncForPC(pc).Name()

		logFatalf("\n### \nERROR: %s \n%s %v \n###", nameOfTheFunction, message, err)
	}
}
