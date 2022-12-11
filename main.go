package main

import "C"

import (
	"fmt"
	"log"

	"github.com/sbinet/go-python"
)

func main() {
	python.Initialize()
	defer python.Finalize()

	posApiModule := python.PyImport_ImportModule("posapi")
	if posApiModule == nil {
		panic("Error importing module")
	}

	obj := posApiModule.GetAttrString("PosApi")
	if obj == nil {
		panic("Error importing obj")
	}

	out := obj.CallFunctionObjArgs("checkApi")
	if out == nil {
		panic("could not out of obj")
	}

	fmt.Printf("posApi.checkApi() = %q\n",
		python.PyString_AsString(out),
	)

	temp := out.CallFunction()
	if out == nil {
		log.Fatalf("error calling foo()\n")
	}

	fmt.Printf("posApi.checkApi() = %q\n",
		python.PyString_AsString(temp),
	)

	checkApi := obj.GetAttrString("checkApi")
	if checkApi == nil {
		panic("checkApi is nil")
	}

	fmt.Printf("getattrstring() = %q\n",
		python.PyString_AsString(checkApi),
	)

	out = checkApi.CallMethod("O")
	if out == nil {
		panic("could not out")
	}
	// The Python function takes no params but when using the C api
	// we're required to send (empty) *args and **kwargs anyways.
	// helloFunc.CallObject(python.PyDict_New())

	// test := helloFunc.CallMethod("checkApi")
	// fmt.Print(test)
}
