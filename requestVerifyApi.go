package goAppstoreIap

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var requestVerifyApi = appStoreRequestVerifyApi

func appStoreRequestVerifyApi(url string, reqData iosIapVerifyReceiptRequest) (jsonResp *iosIapVerifyReceiptResponse, err error) {
	buf := &bytes.Buffer{}
	err = json.NewEncoder(buf).Encode(reqData)
	if err != nil {
		return
	}
	resp, err := http.Post(url, "application/json", buf)
	defer resp.Body.Close()
	jsonRespByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	jsonResp = &iosIapVerifyReceiptResponse{}
	err = json.Unmarshal(jsonRespByte, jsonResp)
	if err != nil {
		return
	}
	return
}
