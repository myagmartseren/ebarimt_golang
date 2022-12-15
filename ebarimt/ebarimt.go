package ebarimt

import (
	"encoding/json"
	"errors"

	"github.com/sbinet/go-python"
)

type PosApiModule struct {
	PosApi         *python.PyObject
	checkApi       *python.PyObject
	getInformation *python.PyObject
	callFunction   *python.PyObject
	put            *python.PyObject
	returnBill     *python.PyObject
	sendData       *python.PyObject
}

func Init(path ...string) (*PosApiModule, error) {
	err := python.Initialize()
	if err != nil {
		return nil, err
	}
	posApiModule := python.PyImport_ImportModule("posapi")
	if posApiModule == nil {
		return nil, errors.New("Error importing module")
	}
	posApiClass := posApiModule.GetAttrString("PosApi")
	if posApiClass == nil {
		return nil, errors.New("Error binding the class of PostApi")
	}
	posApiObj := posApiClass.CallFunction()
	if posApiObj == nil {
		return nil, errors.New("Error creating the object of PosApi")
	}
	checkApi := posApiObj.GetAttrString("checkApi")
	if checkApi == nil {
		return nil, errors.New("Error binding function of CheckApi")
	}

	getInformation := posApiObj.GetAttrString("getInformation")
	if getInformation == nil {
		return nil, errors.New("Error binding function of CheckApi")
	}

	callFunction := posApiObj.GetAttrString("callFunction")
	if callFunction == nil {
		return nil, errors.New("Error binding function of CheckApi")
	}
	put := posApiObj.GetAttrString("put")
	if put == nil {
		return nil, errors.New("Error binding function of CheckApi")
	}
	returnBill := posApiObj.GetAttrString("returnBill")
	if returnBill == nil {
		return nil, errors.New("Error binding function of CheckApi")
	}
	sendData := posApiObj.GetAttrString("sendData")
	if sendData == nil {
		return nil, errors.New("Error binding function of CheckApi")
	}

	return &PosApiModule{
		PosApi:         posApiObj,
		checkApi:       checkApi,
		getInformation: getInformation,
		callFunction:   callFunction,
		put:            put,
		returnBill:     returnBill,
		sendData:       sendData,
	}, nil
}

func (p *PosApiModule) CheckApi() (*APIOutput, error) {
	outputP := p.checkApi.CallFunction()
	if outputP == nil {
		return nil, errors.New("Error calling funtion checkApi")
	}
	var output APIOutput
	if err := json.Unmarshal([]byte(python.PyString_AsString(outputP)), &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (p *PosApiModule) GetInformation() (*InformationOutput, error) {
	outputP := p.getInformation.CallFunction()
	if outputP == nil {
		return nil, errors.New("Error calling funtion checkApi")
	}

	var output InformationOutput
	if err := json.Unmarshal([]byte(python.PyString_AsString(outputP)), &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (p *PosApiModule) CallFunction(functionName, params string) (string, error) {
	outputP := p.callFunction.CallFunction(functionName, params)
	if outputP == nil {
		return "", errors.New("Error calling put function")
	}
	return python.PyString_AsString(outputP), nil
}

func (p *PosApiModule) Put(input *PutInput) (*PutOutput, error) {
	outputP := p.put.CallFunction(input)
	if outputP == nil {
		return nil, errors.New("Error calling put function")
	}
	var output PutOutput
	if err := json.Unmarshal([]byte(python.PyString_AsString(outputP)), &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (p *PosApiModule) ReturnBill(input *BillInput) (*PutOutput, error) {
	outputP := p.returnBill.CallFunction(input)
	if outputP == nil {
		return nil, errors.New("Error calling put function")
	}
	var output PutOutput
	if err := json.Unmarshal([]byte(python.PyString_AsString(outputP)), &output); err != nil {
		return nil, err
	}
	return &output, nil
}

func (p *PosApiModule) SendData() (*DataOutput, error) {
	outputP := p.sendData.CallFunction()
	if outputP == nil {
		return nil, errors.New("Error calling sendData function")
	}

	var output DataOutput
	if err := json.Unmarshal([]byte(python.PyString_AsString(outputP)), &output); err != nil {
		return nil, err
	}
	return &output, nil
}
