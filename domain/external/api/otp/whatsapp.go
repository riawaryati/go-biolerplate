package otp

type SendWhatsapp struct {
	PartnerID         int    `json:"partner_id"`
	OTP               string `json:"otp"`
	Expired           string `json:"expired"`
	DestinationNumber string `json:"destination_number"`
}

type MessageBirdWhatsappAPI struct {
	DestinationNumber string                        `json:"to"`
	Origin            string                        `json:"from"`
	Type              string                        `json:"type"`
	Content           MessageBirdWhatsappOTPContent `json:"content"`
}

type MessageBirdWhatsappOTPContent struct {
	HSM MessageBirdWhatsappOTPHSM `json:"hsm"`
}

type MessageBirdWhatsappOTPHSM struct {
	Namespace    string                         `json:"namespace"`
	Language     MessageBirdWhatsappOTPLanguage `json:"language"`
	TemplateName string                         `json:"templateName"`
	Params       []MessageBirdWhatsappOTPParams `json:"params"`
}

type MessageBirdWhatsappOTPLanguage struct {
	Code string `json:"code"`
}

type MessageBirdWhatsappOTPParams struct {
	Default string `json:"default"`
}
