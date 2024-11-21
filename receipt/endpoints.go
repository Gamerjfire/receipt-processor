package receipt

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var givenDataMap map[string]int = make(map[string]int)

func GetData(requestIn *gin.Context) {
	entryValue, validity := givenDataMap[requestIn.Params.ByName("id")]
	if validity {
		requestIn.IndentedJSON(http.StatusOK, PointsReturn{Points: entryValue})
	} else {
		requestIn.IndentedJSON(http.StatusNotFound, "No found entry under that id")
	}
}

func StoreData(dataToStore *gin.Context) {
	var receipt TotalReceipt

	if err := dataToStore.BindJSON(&receipt); err != nil {
		fmt.Println(err)
		dataToStore.IndentedJSON(http.StatusInternalServerError, "JSON isn't in required format.")
	} else {
		newUUID := uuid.New()

		givenDataMap[newUUID.String()] = ProcessReceipt(receipt)

		dataToStore.IndentedJSON(http.StatusOK, IdReturn{ReceiptId: newUUID.String()})
	}
}
