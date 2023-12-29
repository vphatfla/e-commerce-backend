package model

// OrderItem struct represents the OrderItem table
type OrderItem struct {
	OrderItemID int     `json:"orderItemID"`
	OrderID     int     `json:"orderID"`
	ProductID   int     `json:"productID"`
	Quantity    int     `json:"quantity"`
	SubTotal    float64 `json:"subTotal"`
}
