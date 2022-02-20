package logistic

type BorzoGetRatesAPI struct {
	Matter string               `json:"matter"`
	Type   string               `json:"type"`
	Points []BorzoGetRatesPoint `json:"points"`
}

type BorzoGetRatesPoint struct {
	Address   string `json:"address"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type BorzoGetRatesResponse struct {
	IsSuccessful bool                `json:"is_successful"`
	Order        BorzoGetRatesDetail `json:"order"`
}

type BorzoGetRatesDetail struct {
	DeliveryFeeAmount string `json:"delivery_fee_amount"`
}

type BorzoRequestPickupAPI struct {
	Type          string                    `json:"type"`
	Matter        string                    `json:"matter"`
	IsBoxRequired bool                      `json:"is_motobox_required"`
	PaymentMethod string                    `json:"payment_method"`
	VehicleTypeID int                       `json:"vehicle_type_id"`
	Points        []BorzoRequestPickupPoint `json:"points"`
}

type BorzoRequestPickupPoint struct {
	Address       string                          `json:"address"`
	Latitude      string                          `json:"latitude"`
	Longitude     string                          `json:"longitude"`
	Note          string                          `json:"note"`
	ContactPerson BorzoRequestPickupContactPerson `json:"contact_person"`
}

type BorzoRequestPickupContactPerson struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

type BorzoRequestPickupResponse struct {
	IsSuccessful bool                     `json:"is_successful"`
	Order        BorzoRequestPickupDetail `json:"order"`
}

type BorzoRequestPickupDetail struct {
	OrderID           int64                             `json:"order_id"`
	Ordername         string                            `json:"order_name"`
	VehicleTypeID     int                               `json:"vehicle_type_id"`
	RequestPickupTime string                            `json:"created_datetime"`
	Points            []BorzoRequestPickupPointResponse `json:"points"`
}

type BorzoRequestPickupPointResponse struct {
	PointID     int64  `json:"points_id"`
	TrackingURL string `json:"tracking_url"`
}
