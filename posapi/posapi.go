package posapi

/*
#include <stdio.h>
#include <errno.h>
#include <stdlib.h>
*/
import (
	"C"
)

import (
	"fmt"
	"sync"
	"unsafe"

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

func Open() (*PosAPI, error) {

	lib, err := dl.Open("./libPosAPI.so", dl.RTLD_LAZY)
	if err != nil {
		fmt.Println("err", err.Error())
		return nil, err
	}

	var checkAPIC func() *C.char
	lib.Sym("checkApi", &checkAPIC)

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

	return &PosAPI{
		Lib:             lib,
		checkAPIC:       checkAPIC,
		getInformationC: getInformationC,
		callFunctionC:   callFunctionC,
		putC:            putC,
		returnBillC:     returnBillC,
		sendDataC:       sendDataC,
	}, nil
}

func (p *PosAPI) Close() {
	p.Lib.Close()
}
func (p *PosAPI) CheckAPI() string {
	cResult := p.checkAPIC()
	defer C.free(unsafe.Pointer(cResult))
	result := C.GoString(cResult)
	return result
}

func (p *PosAPI) GetInformation() string {
	cResult := p.getInformationC()
	defer C.free(unsafe.Pointer(cResult))
	result := C.GoString(cResult)
	return result
}

func (p *PosAPI) CallFunction(function string, params string) string {
	cParams := C.CString(params)
	defer C.free(unsafe.Pointer(cParams))

	cFunction := C.CString(function)
	defer C.free(unsafe.Pointer(cFunction))

	cResult := p.callFunctionC(cFunction, cParams)
	defer C.free(unsafe.Pointer(&cResult))

	result := C.GoString(cResult)

	return result
}

func (p *PosAPI) Put(params string) string {
	cParams := C.CString(params)
	defer C.free(unsafe.Pointer(cParams))

	cResult := p.putC(cParams)
	defer C.free(unsafe.Pointer(cParams))

	result := C.GoString(cResult)
	return result
}

func (p *PosAPI) ReturnBill(params string) string {
	cParams := C.CString(params)
	defer C.free(unsafe.Pointer(cParams))

	cResult := p.returnBillC(cParams)
	defer C.free(unsafe.Pointer(cParams))

	result := C.GoString(cResult)
	return result
}

func (p *PosAPI) SendData() string {
	cResult := p.sendDataC()
	defer C.free(unsafe.Pointer(cResult))
	result := C.GoString(cResult)
	return result
}
