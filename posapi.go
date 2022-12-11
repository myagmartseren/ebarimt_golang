package ebarimt

import (
	"bytes"
	"encoding/json"
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

func (p *PosAPI) checkApi() (*APIOutput, error) {
	resp, err := http.Get(p.URL + "/checkApi")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var output APIOutput
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, err
	}

	return &output, nil
}

func (p *PosAPI) GetInformation() (*InformationOutput, error) {
	resp, err := http.Get(p.URL + "/getInformation")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var output InformationOutput
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, err
	}

	return &output, nil
}

func (p *PosAPI) callFunction(function, params string) (interface{}, error) {
	data := map[string]string{
		"functionName": function,
		"data":         params,
	}
	json_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(p.URL+"/callFunction", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	var output map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&output)
	return output, nil
}

func (p *PosAPI) Put(input *PutInput) (*PutOutput, error) {
	json_data, err := json.Marshal(map[string]PutInput{"data": *input})

	resp, err := http.Post(p.URL+"/put", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ouput PutOutput
	if err := json.NewDecoder(resp.Body).Decode(&ouput); err != nil {
		return nil, err
	}

	return &ouput, nil
}
func (p *PosAPI) ReturnBill(input *BillInput) (*BillOutput, error) {
	json_data, err := json.Marshal(map[string]BillInput{"data": *input})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(p.URL+"/returnBill", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var output BillOutput
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, err
	}

	return &output, nil
}

func (p *PosAPI) SendData() (*DataOutput, error) {
	resp, err := http.Get(p.URL + "/sendData")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ouput DataOutput
	if err := json.NewDecoder(resp.Body).Decode(&ouput); err != nil {
		return nil, err
	}
	return &ouput, nil
}
