package main

import "C"

import (
	"fmt"

	"github.com/sbinet/go-python"
)

func main() {
	python.Initialize()
	defer python.Finalize()

	fooModule := python.PyImport_ImportModule("posapi")
	if fooModule == nil {
		panic("Error importing module")
	}

	helloFunc := fooModule.GetAttrString("PosApi")
	if helloFunc == nil {
		panic("Error importing function")
	}

	// The Python function takes no params but when using the C api
	// we're required to send (empty) *args and **kwargs anyways.
	helloFunc.CallObject(python.PyDict_New())

	test := helloFunc.CallMethod("checkApi")
	fmt.Print(test)
}
