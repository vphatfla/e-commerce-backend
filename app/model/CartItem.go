package model

// CartItem struct represents the CartItem table
type CartItem struct {
	CartItemID int     `json:"cartItemID"`
	CartID     int     `json:"cartID"`
	ProductID  int     `json:"productID"`
	Quantity   int     `json:"quantity"`
	SubTotal   float64 `json:"subTotal"`
}
