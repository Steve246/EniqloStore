package dto

type RequestRegistBody struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	Password    string `json:"password"`
}

type RequestLoginBody struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type RequestCustomerRegistBody struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

type RequestProduct struct {
	Name        string `json:"name"`
	SKU         string `json:"sku"`
	Category    string `json:"category"`
	ImageURL    string `json:"imageUrl"`
	Notes       string `json:"notes"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Location    string `json:"location"`
	IsAvailable bool   `json:"isAvailable"`
}

type ProductQueryParams struct {
	ID          string `form:"id"`
	Name        string `form:"name"`
	Limit       int    `form:"limit,default=5"`
	Offset      int    `form:"offset,default=0"`
	IsAvailable string `form:"isavailable"`
	Category    string `form:"category"`
	SKU         string `form:"sku"`
	Price       string `form:"price"`
	InStock     string `form:"inStock"`
	CreatedAt   string `form:"createdAt"`
}

// CheckoutRequest struct represents the request payload for checkout
type CheckoutRequest struct {
	CustomerID     string    `json:"customerId"`
	ProductDetails []Product `json:"productDetails"`
	Paid           int       `json:"paid"`
	Change         int       `json:"change"`
}

// Product struct represents the product details
type Product struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
