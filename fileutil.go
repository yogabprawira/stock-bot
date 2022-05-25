package main

import (
	"encoding/json"
	"io/ioutil"
)

func SaveFile(symbol string, resp Response) error {
	filename := symbol + ".json"
	jsonByte, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, jsonByte, 0644)
}

func LoadFile(symbol string) (Response, error) {
	filename := symbol + ".json"
	jsonByte, err := ioutil.ReadFile(filename)
	if err != nil {
		return Response{}, err
	}
	var resp Response
	err = json.Unmarshal(jsonByte, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}
