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
	"encoding/json"
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

func Open(path ...string) (*PosAPI, error) {
	so_path := "libPosAPI.so"
	if len(path) > 0 {
		so_path = path[0]
	}
	lib, err := dl.Open(so_path, dl.RTLD_LAZY)
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
func (p *PosAPI) CheckAPI() (*APIOutput, error) {
	cResult := p.checkAPIC()
	defer C.free(unsafe.Pointer(cResult))

	var response APIOutput
	if err := json.Unmarshal([]byte(C.GoString(cResult)), &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (p *PosAPI) GetInformation() (*InformationOutput, error) {
	var response InformationOutput
	cResult := p.getInformationC()
	defer C.free(unsafe.Pointer(cResult))

	if err := json.Unmarshal([]byte(C.GoString(cResult)), &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (p *PosAPI) CallFunction(function string, params string) string {
	//  var response
	cParams := C.CString(params)
	defer C.free(unsafe.Pointer(cParams))

	cFunction := C.CString(function)
	defer C.free(unsafe.Pointer(cFunction))

	cResult := p.callFunctionC(cFunction, cParams)
	defer C.free(unsafe.Pointer(cResult))

	result := C.GoString(cResult)

	return result
}

func (p *PosAPI) Put(params string) (*PutOutput, error) {
	var response PutOutput
	cParams := C.CString(params)
	defer C.free(unsafe.Pointer(cParams))

	cResult := p.putC(cParams)
	defer C.free(unsafe.Pointer(cParams))
	if err := json.Unmarshal([]byte(C.GoString(cResult)), &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *PosAPI) ReturnBill(params string) (*BillOutput, error) {
	cParams := C.CString(params)
	defer C.free(unsafe.Pointer(cParams))

	cResult := p.returnBillC(cParams)
	defer C.free(unsafe.Pointer(cParams))
	var response BillOutput

	if err := json.Unmarshal([]byte(C.GoString(cResult)), &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *PosAPI) SendData() (*DataOutput, error) {
	var response DataOutput
	cResult := p.sendDataC()
	defer C.free(unsafe.Pointer(cResult))

	if err := json.Unmarshal([]byte(C.GoString(cResult)), &response); err != nil {
		return nil, err
	}

	return &response, nil
}
