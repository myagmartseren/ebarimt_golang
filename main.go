package main

/*
#include <stdlib.h>
*/
import (
	"C"
)

import (
	"fmt"
	"sync"

	"github.com/peakdot/dl"
)

type PosAPI struct {
	Lib             *dl.DL
	checkAPIC       func() *C.char
	getInformationC func() *C.char
	callFunctionC   func(function *C.char, params *C.char) *C.char
	putC            func(params *C.char) *C.char
	returnBillC     func(params *C.char) *C.char
	sendDataC       func() *C.char
	mux             sync.Mutex
}

func main() {
	lib, err := dl.Open("./x64/libssl.so.1.0.0", dl.RTLD_LAZY)
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}

	var checkAPIC func() *C.char
	lib.Sym("checkAPI", &checkAPIC)

	var getInformationC func() *C.char
	lib.Sym("getInformation", &getInformationC)

	var callFunctionC func(function *C.char, params *C.char) *C.char
	lib.Sym("callFunction", &callFunctionC)

	var putC func(params *C.char) *C.char
	lib.Sym("put", &putC)

	var returnBillC func(params *C.char) *C.char
	lib.Sym("returnBill", &returnBillC)

	var sendDataC func() *C.char
	lib.Sym("sendData", &sendDataC)

	o := &PosAPI{
		Lib:       lib,
		checkAPIC: checkAPIC,
	}

	fmt.Println(o.checkAPI())
}

func (s *PosAPI) checkAPI() string {
	r := s.checkAPIC()
	fmt.Println("working", r)
	result := C.GoString(r)
	return result
}
