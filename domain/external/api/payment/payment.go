package payment

import "time"

type GetVirtualAccount struct {
	Invoice               string    `json:"invoice"`
	BankCode              string    `json:"bank_code"`
	OriginName            string    `json:"origin_name"`
	VirtualAccountNumber  string    `json:"virtual_account_number"`
	IsStrictAmountPayment bool      `json:"is_strict_amount_payment"`
	TotalPayment          float64   `json:"total_payment"`
	ExpirationDate        time.Time `json:"expiration_date"` //ISO 8601
	IsSingleUse           bool      `json:"is_single_use"`   // just for 1 use VA
	Description           string    `json:"description"`
}

type GetVirtualAccountResponse struct {
	PartnerID          string  `json:"partner_id"`
	Invoice            string  `json:"invoice"`
	BankCode           string  `json:"bank_code"`
	MerchantCode       string  `json:"merchant_code"`
	Name               string  `json:"name"`
	VirtualAccount     string  `json:"virtual_account"`
	TotalPayment       float64 `json:"total_payment"`
	PartnerReferenceID string  `json:"partner_reference_id"`
	Status             string  `json:"status"`
}

type GetQRISRequest struct {
	Invoice      string  `json:"invoice"`
	TotalPayment float64 `json:"total_payment"`
}

type GetQRISResponse struct {
	PartnerReferenceID string  `json:"partner_reference_id"`
	QRString           string  `json:"qr_string"`
	TotalPayment       float64 `json:"total_payment"`
	Status             string  `json:"status"`
	Metadata           string  `json:"metadata"`
	ErrorCode          string  `json:"error_code,omitempty"`
	Message            string  `json:"message,omitempty"`
}
