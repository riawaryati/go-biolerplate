package logistic

type JETAuthorizationResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpireIn     string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type JETGetTrackingResponse struct {
	Airwaybill    string                   `json:"awbNumber"`
	ReferenceCode string                   `json:"referenceCode"`
	TotalWeight   float32                  `json:"totalWeight"`
	TotalPackage  int                      `json:"totalPart"`
	LastStatus    string                   `json:"status"`
	Packages      []JETPackageTrackingData `json:"connotes"`
}

type JETPackageTrackingData struct {
	Code        string               `json:"code"`
	Description string               `json:"description"`
	Weight      float32              `json:"weight"`
	Histories   []JETTrackingHistory `json:"tracks"`
}

type JETTrackingHistory struct {
	Date     string `json:"date"`
	Branch   string `json:"branch"`
	Location string `json:"location"`
	Status   string `json:"status"`
}

type JETCreateExternalWaybill struct {
	ID                  string               `json:"JetId"`
	Airwaybill          string               `json:"ReferenceCode"`
	PaymentType         string               `json:"CodType"`
	ShipmentType        string               `json:"ProductCode"`
	Description         string               `json:"Description"`
	SpecialInstrunction string               `json:"SpecialInstrunction"`
	UseInsurance        bool                 `json:"IsInsured"`
	OriginData          JETDetailBookingData `json:"Shipper"`
	DestinationData     JETDetailBookingData `json:"Receiver"`
	ItemValue           int64                `json:"ItemValue"`
	ItemLength          float32              `json:"Length"`
	ItemWidth           float32              `json:"Width"`
	ItemHeight          float32              `json:"Height"`
	ItemWeight          float32              `json:"Weight"`
}

type JETDetailBookingData struct {
	Name    string `json:"Name"`
	Address string `json:"Address"`
	Phone   string `json:"Phone"`
	Email   string `json:"Email"`
}

type JETRequestPickupBody struct {
	ShipperName     string  `json:"ShipperName"`
	ShipperPhone    string  `json:"ShipperPhone"`
	ShipperAddress  string  `json:"ShipperAddress"`
	VehicleCode     string  `json:"VehicleCode"`
	Longitude       float64 `json:"Longitude"`
	Latitude        float64 `json:"Latitude"`
	PickupItemCount int     `json:"PickupItemCount"`
	ReferenceNo     string  `json:"ReferenceNo"`
}

type JETRequestPickupResponse struct {
	BookingCode string `json:"code"`
}

type JETGetRatesResponse struct {
	Origin      string               `json:"origin"`
	Destination string               `json:"destination"`
	Services    []JETGetRatesService `json:"services"`
}

type JETGetRatesService struct {
	Code     string  `json:"code"`
	Name     string  `json:"name"`
	TotalFee float64 `json:"totalFee"`
}

type JETGetRatesRequest struct {
	Origin             string             `json:"origin"`
	OriginZipCode      string             `json:"OriginZipCode"`
	Destination        string             `json:"destination"`
	DestinationZipCode string             `json:"DestinationZipCode"`
	IsInsured          bool               `json:"isInsured"`
	ItemValue          float64            `json:"itemValue"`
	Items              []JETGetRatesItems `json:"items"`
}

type JETGetRatesItems struct {
	Weight        float64 `json:"weight"`        // In Kg
	Height        float64 `json:"height"`        // In cm
	Width         float64 `json:"width"`         // In cm
	Length        float64 `json:"length"`        // In cm
	PackagingCode string  `json:"packagingCode"` // In cm
	PackagingQty  int     `json:"packagingQty"`  // In cm
}
