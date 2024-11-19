package main

import (
	"fmt"

	"receiptprocessor.com/receipt"
)

func main() {
	var receiptValue int = receipt.ProcessReceipt(receipt.ExampleReceipt)
	fmt.Println(receiptValue)
}
