package entity

type CartItem struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
