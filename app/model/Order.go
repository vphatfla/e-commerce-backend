package model

import "time"

// Order struct represents the Order table
type Order struct {
	OrderID   int       `json:"orderID"`
	Amount    float64   `json:"amount"`
	UserID    int       `json:"userID"`
	OrderDate time.Time `json:"orderDate"`
}
