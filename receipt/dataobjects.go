package receipt

//Go doesn't have default Date/Time so we will have to parse them later
type TotalReceipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Entry `json:"items"`
	Total        float64 `json:"total"`
}

type Entry struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price"`
}

//Example data TODO Remove
var ExampleReceipt = TotalReceipt{Retailer: "Target", PurchaseDate: "2022-01-01", PurchaseTime: "13:01",
	Items: []Entry{{ShortDescription: "Mountain Dew 12PK", Price: 6.49}, {ShortDescription: "Emils Cheese Pizza", Price: 12.25},
		{ShortDescription: "Knorr Creamy Chicken", Price: 1.26}, {ShortDescription: "Doritos Nacho Cheese", Price: 3.35},
		{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: 12.00}}, Total: 35.35,
}
