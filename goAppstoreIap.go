package goAppstoreIap

import (
	"fmt"
)

type iosIapVerifyReceiptRequest struct {
	ReceiptData string `json:"receipt-data"`
}
type iosIapVerifyReceiptResponse struct {
	Status  int     `json:"status"`
	Receipt Receipt `json:"receipt"`
}
type Receipt struct {
	InApp []ReceiptTransaction `json:"in_app"` // it may have two transaction.

	BundleId                   string `json:"bundle_id"`
	AppVersion                 string `json:"application_version"`
	OriginalApplicationVersion string `json:"original_application_version"`
}
type ReceiptTransaction struct {
	ProductId     string `json:"product_id"`
	TransactionId string `json:"transaction_id"`

	Quantity              string `json:"quantity"`
	OriginalTransactionId string `json:"original_transaction_id"`
	//暂时 无法读取时间 格式大概类似于:  2014-04-16 18:26:18 Etc/GMT
	PurchaseDate               string `json:"purchase_date"`
	OriginalPurchaseDate       string `json:"original_purchase_date"`
	SubscriptionExpirationDate string `json:"expires_date"`
	CancellationDate           string `json:"cancellation_date"`
	AppItemId                  string `json:"app_item_id"`
	ExternalVersionIdentifier  string `json:"version_external_identifier"`
	WebOrderLineItemId         string `json:"web_order_line_item_id"`
}

type VerifyReceiptResponse struct {
	TransactionList []ReceiptTransaction //不一定只有一个交易
}

//订单正确性验证
func VerifyReceipt(Receipt string) (response *Receipt, err error) {
	//问苹果验证Receipt的有效性
	reqData := iosIapVerifyReceiptRequest{
		ReceiptData: Receipt,
	}
	requestVerify := func(url string) (jsonResp *iosIapVerifyReceiptResponse, isTest bool, err error) {
		jsonResp, err = requestVerifyApi(url, reqData)
		if err != nil {
			return
		}
		if jsonResp.Status == 21007 {
			return nil, true, nil
		}
		return
	}
	//try normal first
	JsonResp, isTest, err := requestVerify("https://buy.itunes.apple.com/verifyReceipt")
	if err != nil {
		return nil, err
	}
	if isTest {
		// try sand box again
		JsonResp, isTest, err = requestVerify("https://sandbox.itunes.apple.com/verifyReceipt")
		if err != nil {
			return nil, err
		}
		if isTest {
			return nil, fmt.Errorf("[appStorePayVerify] testurl return should in testmod?")
		}
	}
	if JsonResp.Status != 0 {
		return nil, fmt.Errorf("[appStorePayVerify] JsonResp.Status[%d]!=0", JsonResp.Status)
	}
	return &JsonResp.Receipt, nil
}
