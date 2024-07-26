package model

type Product struct {
	Name        string  `json:"name"`
	ProductType string  `json:"product_type"`
	Picture     string  `json:"picture"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
