package receipt_test

import (
	"encoding/json"
	"testing"

	"receiptprocessor/receipt"

	"github.com/google/go-cmp/cmp"
)

func TestJSONificaiton(t *testing.T) {
	var ExampleReceipt = receipt.TotalReceipt{Retailer: "Target", PurchaseDate: "2022-01-01", PurchaseTime: "13:01",
		Items: []receipt.Entry{{ShortDescription: "Mountain Dew 12PK", Price: "6.49"}, {ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"}, {ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"}}, Total: "35.35",
	}

	var stringVersionOfReceipt string = "{\"retailer\": \"Target\",\"purchaseDate\": \"2022-01-01\",\"purchaseTime\": \"13:01\",\"items\": [{\"shortDescription\": \"Mountain Dew 12PK\",\"price\": \"6.49\"},{\"shortDescription\": \"Emils Cheese Pizza\",\"price\": \"12.25\"},{\"shortDescription\": \"Knorr Creamy Chicken\",\"price\": \"1.26\"},{\"shortDescription\": \"Doritos Nacho Cheese\",\"price\": \"3.35\"},{\"shortDescription\": \"   Klarbrunn 12-PK 12 FL OZ  \",\"price\": \"12.00\"}],\"total\": \"35.35\"}"

	var receipt receipt.TotalReceipt

	json.Unmarshal([]byte(stringVersionOfReceipt), &receipt)

	if !cmp.Equal(receipt, ExampleReceipt) {
		t.Errorf("Receipts should be the same after unmarshalling.")
	}

}
