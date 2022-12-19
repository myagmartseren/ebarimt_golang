package ebarimt

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
	Lib            *dl.DL
	checkAPI       func() *C.char
	getInformation func() *C.char
	callFunction   func(function *C.char, params *C.char) *C.char
	put            func(params *C.char) *C.char
	returnBill     func(params *C.char) *C.char
	sendData       func() *C.char
	mux            sync.Mutex
}

func Open(path ...string) (*PosAPI, error) {
	so_path := "/usr/lib/libPosAPI.so"
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
		Lib:            lib,
		checkAPI:       checkAPIC,
		getInformation: getInformationC,
		callFunction:   callFunctionC,
		put:            putC,
		returnBill:     returnBillC,
		sendData:       sendDataC,
	}, nil
}

func (p *PosAPI) Close() {
	p.Lib.Close()
}
func (p *PosAPI) CheckAPI() (*APIOutput, error) {
	cResult := p.checkAPI()
	defer C.free(unsafe.Pointer(cResult))

	var response APIOutput
	if err := json.Unmarshal([]byte(C.GoString(cResult)), &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (p *PosAPI) GetInformation() (*InformationOutput, error) {
	var response InformationOutput
	cResult := p.getInformation()
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

	cResult := p.callFunction(cFunction, cParams)
	defer C.free(unsafe.Pointer(cResult))

	result := C.GoString(cResult)

	return result
}

func (p *PosAPI) Put(params string) (*PutOutput, error) {
	var response PutOutput
	cParams := C.CString(params)
	defer C.free(unsafe.Pointer(cParams))

	cResult := p.put(cParams)
	defer C.free(unsafe.Pointer(cParams))
	if err := json.Unmarshal([]byte(C.GoString(cResult)), &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *PosAPI) ReturnBill(params string) (*BillOutput, error) {
	cParams := C.CString(params)
	defer C.free(unsafe.Pointer(cParams))

	cResult := p.returnBill(cParams)
	defer C.free(unsafe.Pointer(cParams))
	var response BillOutput

	if err := json.Unmarshal([]byte(C.GoString(cResult)), &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *PosAPI) SendData() (*DataOutput, error) {
	var response DataOutput
	cResult := p.sendData()
	defer C.free(unsafe.Pointer(cResult))

	if err := json.Unmarshal([]byte(C.GoString(cResult)), &response); err != nil {
		return nil, err
	}

	return &response, nil
}
