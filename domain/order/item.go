package order

type Item struct {
	ItemID      int64  `json:"lineItemId" gorm:"primaryKey;autoIncrement" db:"item_id"`
	ItemCode    string `json:"itemCode" db:"item_code"`
	Description string `json:"description" db:"description"`
	Quantity    int    `json:"quantity" db:"quantity"`
	OrderID     int64  `json:"orderId" db:"order_id"`
}
