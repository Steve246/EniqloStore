package dto

import "time"

type ProductInfo struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
