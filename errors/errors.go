package errors

import (
	"log"
	"runtime"
)

var logFatalf = log.Fatalf
var logPrintf = log.Printf

func IfAnErrorOccursCallsLogFatal(err error, message string) {
	if err != nil {
		pc, _, _, _ := runtime.Caller(1)
		nameOfTheFunction := runtime.FuncForPC(pc).Name()

		logFatalf("\n### \nERROR: %s \n%s %v \n###", nameOfTheFunction, message, err)
	}
}

func IfAnErrorOccursCallsLogPrint(err error, message string) {
	if err != nil {
		pc, _, _, _ := runtime.Caller(1)
		nameOfTheFunction := runtime.FuncForPC(pc).Name()

		logPrintf("\n### \nERROR: %s \n%s %v \n###", nameOfTheFunction, message, err)
	}
}
