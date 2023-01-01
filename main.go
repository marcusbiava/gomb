package main

import (
	"errors"
	"fmt"
	errors2 "github.com/marcusbiava/gomb/errors"
)

func main() {
	fmt.Println("Hello, World!")
	f1()
}

func f1() {
	f2()
}

func f2() {
	err := errors.New("error")
	errors2.IfAnErrorOccursCallsLogFatal(err, "error here")

}
