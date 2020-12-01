package models

import "time"

// Product struct
type Product struct {
	Image      string  `json:"image"`
	ImageAlt   string  `json:"image_alt"`
	Price      float64 `json:"price"`
	Promotion  float64 `json:"promotion"`
	Name       string  `json:"name"`
	Desciption string  `json:"description"`
}

// Customer struct
type Customer struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	LoggedIn bool   `json:"loggedin"`
}

// Order struct
type Order struct {
	Product      `json:"product"`
	Customer     `json:"custmer"`
	CustomerID   int       `json:"customer_id"`
	ProductID    int       `json:"product_id"`
	Price        float64   `json:"sell_price"`
	PurchaseDate time.Time `json:"purchase_date"`
}