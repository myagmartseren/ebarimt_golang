package main

import "C"

import (
	"fmt"

	"github.com/sbinet/go-python"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	python.Initialize()
	defer python.Finalize()

	posapi := python.PyImport_ImportModule("posapi")
	if posapi == nil {
		panic("Error importing module")
	}
	posapi_obj := posapi.GetAttrString("PosApi")
	if posapi_obj == nil {
		panic("could not retrieve 'posapi.PosApi'")
	}
	out := posapi_obj.CallMethodObjArgs("checkApi")
	if out == nil {
		panic("could not dump checkApi")
	}
	fmt.Printf(" = %q\n",
		python.PyString_AsString(out),
	)

	fmt.Printf("cPickle.loads(%q) =\n",
		python.PyString_AsString(out),
	)
}
