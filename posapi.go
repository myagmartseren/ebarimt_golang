package ebarimt

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

func Init() (*PosAPI, error) {
	lib, err := dl.Open("libssl.so.1.0.0", dl.RTLD_LAZY)
	if err != nil {
		fmt.Println("err", err.Error())
		return nil, err
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

	posapi := &PosAPI{
		Lib:             lib,
		checkAPIC:       checkAPIC,
		getInformationC: getInformationC,
		callFunctionC:   callFunctionC,
		putC:            putC,
		returnBillC:     returnBillC,
		sendDataC:       sendDataC,
	}
	return posapi, nil
}

func (p *PosAPI) CheckAPI() string {
	r := p.checkAPIC()
	result := C.GoString(r)
	return result
}
func (p *PosAPI) GetInformation() string {
	return ""
}
func (p *PosAPI) CallFunction(function string, params string) string {
	return ""
}
func (p *PosAPI) Put(params string) string {
	return ""

}
func (p *PosAPI) ReturnBill(params string) string {
	return ""
}

func (p *PosAPI) SendData(params string) string {
	return ""
}
