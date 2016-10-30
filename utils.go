package mlib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReadPostStr(r *http.Request) *string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	r.Body.Close()

	ret := buf.String()
	return &ret
}

func ParseJson(jsonStr *string, ptr interface{}) error {
	err := json.Unmarshal([]byte(*jsonStr), ptr)
	if err != nil {
		Error("Unmarshal error %s", err)
		return err
	}

	return nil
}

func MarhalJson(v interface{}) (*string, error) {
	dataBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	ret := string(dataBytes)
	return &ret, nil
}

func ReadConfig(cfgPath *string, cfg interface{}) error {

	dat, err := ioutil.ReadFile(*cfgPath)
	if err != nil {
		Error("Error %s", err)
		return err
	}

	err = json.Unmarshal(dat, cfg)
	if err != nil {
		Error("Unmarshal error %s", err)
		return err
	}
	return nil
}
