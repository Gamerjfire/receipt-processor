package main

import (
	"receiptprocessor/receipt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/receipts/:id/points", receipt.GetData)
	router.POST("/receipts/process", receipt.StoreData)

	router.Run("localhost:8080")
}
