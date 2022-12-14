package ebarimt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type PosAPI struct {
	URL string
}

func Init(url ...string) *PosAPI {
	init_url := "http://localhost:5000"
	if len(url) > 0 {
		init_url = url[0]
	}
	return &PosAPI{
		URL: init_url,
	}

}

func (p *PosAPI) CheckApi() (*APIOutput, error) {
	r, err := http.Get(p.URL + "/checkApi")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Error Status code:%v", r.StatusCode))
	}

	var output APIOutput

	if err := json.NewDecoder(r.Body).Decode(&output); err != nil {
		return nil, err
	}

	return &output, nil
}

func (p *PosAPI) GetInformation() (*InformationOutput, error) {
	r, err := http.Get(p.URL + "/getInformation")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Error Status code:%v", r.StatusCode))
	}

	var output InformationOutput
	if err := json.NewDecoder(r.Body).Decode(&output); err != nil {
		return nil, err
	}

	return &output, nil
}

func (p *PosAPI) callFunction(function, params string) (interface{}, error) {
	data, err := json.Marshal(map[string]string{
		"functionName": function,
		"data":         params,
	})
	if err != nil {
		return nil, err
	}

	r, err := http.Post(p.URL+"/callFunction", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", nil
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Error Status code:%v", r.StatusCode))
	}

	var output map[string]interface{}

	json.NewDecoder(r.Body).Decode(&output)
	return output, nil
}

func (p *PosAPI) Put(input *PutInput) (*PutOutput, error) {
	data, err := json.Marshal(map[string]PutInput{"data": *input})

	r, err := http.Post(p.URL+"/put", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Error Status code:%v", r.StatusCode))
	}

	var ouput PutOutput
	if err := json.NewDecoder(r.Body).Decode(&ouput); err != nil {
		return nil, err
	}

	return &ouput, nil
}
func (p *PosAPI) ReturnBill(input *BillInput) (*BillOutput, error) {
	data, err := json.Marshal(map[string]BillInput{"data": *input})
	if err != nil {
		return nil, err
	}
	r, err := http.Post(p.URL+"/returnBill", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Error Status code:%v", r.StatusCode))
	}

	var output BillOutput
	if err := json.NewDecoder(r.Body).Decode(&output); err != nil {
		return nil, err
	}

	return &output, nil
}

func (p *PosAPI) SendData() (*DataOutput, error) {
	r, err := http.Get(p.URL + "/sendData")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Error Status code:%v", r.StatusCode))
	}

	var ouput DataOutput
	if err := json.NewDecoder(r.Body).Decode(&ouput); err != nil {
		return nil, err
	}
	return &ouput, nil
}
