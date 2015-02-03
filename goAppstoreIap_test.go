package goAppstoreIap

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAppStorePayVerify(ot *testing.T) {
	requestVerifyApi = requestVerifyApiMock("IosProductId", "TransactionId")
	resp, err := VerifyReceipt("")
	if err != nil {
		panic(err)
	}
	if resp.InApp[0].ProductId != "IosProductId" {
		panic(fmt.Errorf(`resp.ProductId[%s]!="IosProductId"`, resp.InApp[0].ProductId))
	}
	if resp.InApp[0].TransactionId != "TransactionId" {
		panic(fmt.Errorf(`resp.TransactionId[%s]!="TransactionId"`, resp.InApp[0].TransactionId))
	}

}

func TestAppStorePayVerify2(ot *testing.T) {
	inData := `{"status":0, "environment":"Sandbox",
"receipt":{"receipt_type":"ProductionSandbox", "adam_id":0, "app_item_id":0, "bundle_id":"com.example.xxx", "application_version":"1", "download_id":0, "version_external_identifier":0, "request_date":"2015-02-03 17:46:14 Etc/GMT", "request_date_ms":"1422985574541", "request_date_pst":"2015-02-03 09:46:14 America/Los_Angeles", "original_purchase_date":"2013-08-01 07:00:00 Etc/GMT", "original_purchase_date_ms":"1375340400000", "original_purchase_date_pst":"2013-08-01 00:00:00 America/Los_Angeles", "original_application_version":"1.0",
"in_app":[
{"quantity":"1", "product_id":"com.example.xxx.TierA", "transaction_id":"1000000141724470", "original_transaction_id":"1000000141764470", "purchase_date":"2015-02-03 17:45:50 Etc/GMT", "purchase_date_ms":"1422985550000", "purchase_date_pst":"2015-02-03 09:45:50 America/Los_Angeles", "original_purchase_date":"2015-02-03 17:41:08 Etc/GMT", "original_purchase_date_ms":"1422985268000", "original_purchase_date_pst":"2015-02-03 09:41:08 America/Los_Angeles", "is_trial_period":"false"},
{"quantity":"1", "product_id":"com.example.xxx.Tier15", "transaction_id":"1000000141264491", "original_transaction_id":"1000000141764491", "purchase_date":"2015-02-03 17:45:50 Etc/GMT", "purchase_date_ms":"1422985550000", "purchase_date_pst":"2015-02-03 09:45:50 America/Los_Angeles", "original_purchase_date":"2015-02-03 17:41:36 Etc/GMT", "original_purchase_date_ms":"1422985296000", "original_purchase_date_pst":"2015-02-03 09:41:36 America/Los_Angeles", "is_trial_period":"false"}]}}`
	requestVerifyApi = requestVerifyApiMock2([]byte(inData))
	resp, err := VerifyReceipt("")
	if err != nil {
		panic(err)
	}
	if resp.InApp[0].ProductId != "com.example.xxx.TierA" {
		panic(fmt.Errorf(`resp.InApp[0].ProductId != "com.example.xxx.TierA"`, resp.InApp[0].ProductId))
	}
}

func requestVerifyApiMock(AppStoreProductId string, AppStoreTransactionId string) func(url string, reqData iosIapVerifyReceiptRequest) (jsonResp *iosIapVerifyReceiptResponse, err error) {
	return func(url string, req iosIapVerifyReceiptRequest) (jsonResp *iosIapVerifyReceiptResponse, err error) {
		return &iosIapVerifyReceiptResponse{
			Status: 0,
			Receipt: Receipt{
				InApp: []ReceiptTransaction{
					{
						ProductId:     AppStoreProductId,
						TransactionId: AppStoreTransactionId,
					},
				},
			},
		}, nil
	}
}

func requestVerifyApiMock2(outJson []byte) func(url string, reqData iosIapVerifyReceiptRequest) (jsonResp *iosIapVerifyReceiptResponse, err error) {
	return func(url string, req iosIapVerifyReceiptRequest) (jsonResp *iosIapVerifyReceiptResponse, err error) {
		jsonResp = &iosIapVerifyReceiptResponse{}
		err = json.Unmarshal(outJson, jsonResp)
		if err != nil {
			return
		}
		return
	}
}
