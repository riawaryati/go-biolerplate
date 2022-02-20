package logistic

type JNEGetRatesParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type JNEGetRatesResponse struct {
	Price  []JNEGetRatesResponseDetail `json:"price"`
	Error  string                      `json:"error"`
	Status bool                        `json:"status"`
}

type JNEGetRatesResponseDetail struct {
	Service     string `json:"service_display"`
	ServiceCode string `json:"service_code"`
	Currency    string `json:"currency"`
	Price       string `json:"price"`
}
