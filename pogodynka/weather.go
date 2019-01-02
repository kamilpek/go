package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func weather() (*Dane, error) {
	resp, err := http.Get(imgwURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("coś poszło nie tak: %s", resp.Status)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var result Dane
	xml.Unmarshal(bytes, &result)
	return &result, nil
}
