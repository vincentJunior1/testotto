package httpEntity

type OmzetResponse struct {
	MerchantName    string      `json:"merchant_name"`
	Omzet           int         `json:"omzet"`
	TransactionDate interface{} `json:"transaction_date"`
}
