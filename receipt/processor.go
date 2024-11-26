package receipt

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ProcessReceipt(enteredReceipt TotalReceipt) int {
	dateFormat := time.DateOnly
	timeFormat := time.TimeOnly
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	points := 0
	//Check Alphanumerics - 1 point per
	slimmedName := nonAlphanumericRegex.ReplaceAllString(enteredReceipt.Retailer, "")
	points += len(slimmedName)

	//Separate the decimal out
	floatOfTotal, _ := strconv.ParseFloat(enteredReceipt.Total, 64)
	_, decimal := math.Modf(floatOfTotal)

	//Check for round dollar value - 50 points
	if decimal == .0 {
		points += 50
	}

	//Check quarter payments - 25 points
	if int(decimal*100)%25 == 0 {
		points += 25
	}

	//Check number of items - 5 points per 2 (Integer comparison, Go automatically floors the value.)
	points += (len(enteredReceipt.Items) / 2) * 5

	//Trimmed length of item description is mult 3, mult price by .2, round up, that many points.
	for _, givenEntry := range enteredReceipt.Items {
		description := strings.Trim(givenEntry.ShortDescription, " ")
		//Check if multiple of 3
		if len(description)%3 == 0 {
			//Do the multiplication and Ceiling check
			floatOfEntry, _ := strconv.ParseFloat(givenEntry.Price, 64)
			points += int(math.Ceil(floatOfEntry * .2))
		}
	}

	//Purchase date odd - 6 points
	date, dateError := time.Parse(dateFormat, enteredReceipt.PurchaseDate)
	if dateError != nil {
		fmt.Println("Error in date parsing, returning -1, " + dateError.Error())
		return -1
	}
	if date.Day()%2 == 1 {
		points += 6
	}

	//Purchase between 2 and 4 pm - 10 points
	time, timeError := time.Parse(timeFormat, enteredReceipt.PurchaseTime+":00")
	if timeError != nil {
		fmt.Println("Error in time parsing, returning -1, " + timeError.Error())
		return -1
	}
	if time.Hour() >= 14 && time.Hour() <= 16 {
		points += 10
	}

	return points
}
