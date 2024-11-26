package receipt

type TotalReceipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Entry `json:"items"`
	Total        string  `json:"total"`
}

type Entry struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type PointsReturn struct {
	Points int `json:"points"`
}

type IdReturn struct {
	ReceiptId string `json:"id"`
}
