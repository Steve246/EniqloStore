package dto

import "time"

type ProductInfo struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductQuantityChange struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
