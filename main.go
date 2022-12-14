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
	fmt.Println(obj.Check_Callable())
	out1 := obj.CallFunction()
	if out1 == nil {
		log.Fatalf("error calling obj\n")
	}
	fmt.Println("out1", out1)
	checkApi := out1.GetAttrString("checkApi")
	if checkApi == nil {
		panic("could bind checkapi")
	}

	fmt.Println(checkApi)
	out := checkApi.CallFunction()
	if out == nil {
		log.Fatalf("error calling checkApi()\n")
	}

	fmt.Printf("posApi.checkApi() = %q\n",
		python.PyString_AsString(out),
	)

	// The Python function takes no params but when using the C api
	// we're required to send (empty) *args and **kwargs anyways.
	// helloFunc.CallObject(python.PyDict_New())

	// test := helloFunc.CallMethod("checkApi")
	// fmt.Print(test)
}
