package payment

type XenditQRISRequest struct {
	ExternalID  string `json:"external_id"`
	Type        string `json:"type"`
	CallbackURL string `json:"callback_url"`
	Amount      int64  `json:"amount"`
}

type XenditQRISParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type XenditQRISResponse struct {
	ID          string `json:"id"`
	ExternalID  string `json:"external_id"`
	Amount      int64  `json:"amount"`
	QRString    string `json:"qr_string"`
	CallbackURL string `json:"callback_url"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	ErrorCode   string `json:"error_code,omitempty"`
	Message     string `json:"message,omitempty"`
}
