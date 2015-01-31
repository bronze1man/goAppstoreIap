package goAppstoreIap

import (
	"fmt"
	"testing"
)

func TestAppStorePayVerify(ot *testing.T) {
	requestVerifyApi = requestVerifyApiMock("IosProductId", "TransactionId")
	resp, err := VerifyReceipt("")
	if err != nil {
		panic(err)
	}
	if resp.ProductId != "IosProductId" {
		panic(fmt.Errorf(`resp.ProductId[%s]!="IosProductId"`, resp.ProductId))
	}
	if resp.TransactionId != "TransactionId" {
		panic(fmt.Errorf(`resp.TransactionId[%s]!="TransactionId"`, resp.TransactionId))
	}
}

func requestVerifyApiMock(AppStoreProductId string, AppStoreTransactionId string) func(url string, reqData iosIapVerifyReceiptRequest) (jsonResp *iosIapVerifyReceiptResponse, err error) {
	return func(url string, req iosIapVerifyReceiptRequest) (jsonResp *iosIapVerifyReceiptResponse, err error) {
		return &iosIapVerifyReceiptResponse{
			Status: 0,
			Receipt: iosIapVerifyReceiptResponseInner{
				InApp: []iosIapReceiptTransaction{
					{
						ProductId:     AppStoreProductId,
						TransactionId: AppStoreTransactionId,
					},
				},
			},
		}, nil
	}
}
