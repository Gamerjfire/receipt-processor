package main

import (
	"fmt"
	"receiptprocessor/receipt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/receipts/:id/points", receipt.GetData)
	router.POST("/receipts/process", receipt.StoreData)

	router.Run("localhost:8080")

	var receiptValue int = receipt.ProcessReceipt(receipt.ExampleReceipt)
	fmt.Println(receiptValue)
}
