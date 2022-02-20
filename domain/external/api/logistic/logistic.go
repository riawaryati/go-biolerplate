package logistic

type GetRatesData struct {
	LogisticID        int64          `json:"logistic_id"`
	LogisticServiceID int64          `json:"logistic_service_id"`
	ServiceCode       string         `json:"service_code"`
	ProductContains   string         `json:"product_contains"`
	TotalWeight       float64        `json:"total_weight"`  // in gram
	VolumeHeight      float64        `json:"volume_height"` // in cm
	VolumeWidth       float64        `json:"volume_width"`  // in cm
	VolumeLength      float64        `json:"volume_length"` // in cm
	Origin            GetRatesDetail `json:"origin"`
	Destination       GetRatesDetail `json:"destination"`
}

type GetRatesDetail struct {
	FullAddress string `json:"full_address"`
	DistrictID  int64  `json:"district_id"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Code        string `json:"code"`
	ZipCode     string `json:"zipcode"`
}

type RequestPickupData struct {
	LogisticID        int64                   `json:"logistic_id"`
	LogisticServiceID int64                   `json:"logistic_service_id"`
	ServiceCode       string                  `json:"service_code"`
	ProductContains   string                  `json:"product_contains"`
	TotalWeight       float64                 `json:"total_weight"`  // in gram
	VolumeHeight      float64                 `json:"volume_height"` // in cm
	VolumeWidth       float64                 `json:"volume_width"`  // in cm
	VolumeLength      float64                 `json:"volume_length"` // in cm
	Origin            RequestPickupDetail     `json:"origin"`
	Destination       RequestPickupDetail     `json:"destination"`
	Borzo             *BorzoRequestPickupData `json:"borzo"`
}

type RequestPickupDetail struct {
	FullAddress string `json:"full_address"`
	DistrictID  int64  `json:"district_id"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Code        string `json:"code"`
	ZipCode     string `json:"zipcode"`
	Note        string `json:"note"`
	PICName     string `json:"pic_name"`
	PICPhone    string `json:"pic_phone"`
}

type BorzoRequestPickupData struct {
	IsBoxRequired bool   `json:"is_box_required"`
	PaymentMethod string `json:"payment_method"`
	VehicleTypeID int    `json:"vehicle_type_id"`
}

type RequestPickupResponse struct {
	BookingCode string                          `json:"booking_code"`
	Airwaybill  string                          `json:"airwaybill"`
	RequestTime string                          `json:"request_time"`
	Borzo       *BorzoRequestPickupResponseData `json:"borzo"`
}

type BorzoRequestPickupResponseData struct {
	PickupTrackingURL  string   `json:"pickup_tracking_url"`
	SendingTrackingURl []string `json:"sending_tracking_url"`
}
