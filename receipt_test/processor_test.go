package receipt_test

import (
	"testing"

	"receiptprocessor/receipt"
)

// Test on provided dataset 1
func TestProcessorOnDataset1(t *testing.T) {
	ExampleReceipt := receipt.TotalReceipt{Retailer: "Target", PurchaseDate: "2022-01-01", PurchaseTime: "13:01",
		Items: []receipt.Entry{{ShortDescription: "Mountain Dew 12PK", Price: "6.49"}, {ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"}, {ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"}}, Total: "35.35",
	}

	if receipt.ProcessReceipt(ExampleReceipt) != 28 {
		t.Errorf("Value mismatch on expected known result.")
	}
}

// Test on provided dataset 2
func TestProcessorOnDataset2(t *testing.T) {
	ExampleReceipt := receipt.TotalReceipt{Retailer: "M&M Corner Market", PurchaseDate: "2022-03-20", PurchaseTime: "14:33",
		Items: []receipt.Entry{{ShortDescription: "Gatorade", Price: "2.25"}, {ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"}, {ShortDescription: "Gatorade", Price: "2.25"}}, Total: "9.00",
	}

	if receipt.ProcessReceipt(ExampleReceipt) != 109 {
		t.Log(receipt.ProcessReceipt(ExampleReceipt))
		t.Errorf("Value mismatch on expected known result.")
	}
}
