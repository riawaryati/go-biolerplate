package order

import (
	"time"
)

type Order struct {
	OrderID      int64     `json:"orderId" gorm:"primaryKey;autoIncrement" db:"order_id"`
	CustomerName string    `json:"customerName" db:"customer_name"`
	OrderedAt    time.Time `json:"orderedAt" db:"ordered_at"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID;references:OrderID;"`
}

type OrderRequest struct {
	OrderID      int64  `json:"orderId"`
	CustomerName string `json:"customerName"`
	OrderedAt    string `json:"orderedAt"`
	Items        []Item `json:"items"`
}
