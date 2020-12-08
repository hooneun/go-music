package models

import (
	"time"

	"gorm.io/gorm"
)

// Customer struct
type Customer struct {
	gorm.Model
	Name     string  `gorm:"column:name"`
	Email    string  `gorm:"column:email" json:"email"`
	Password string  `json:"password"`
	LoggedIn bool    `gorm:"column:loggedin" json:"loggedin"`
	Orders   []Order `json:"orders"`
}

// TableName customer
func (Customer) TableName() string {
	return "customer"
}

// Product struct
type Product struct {
	gorm.Model
	Image      string  `json:"image"`
	ImageAlt   string  `gorm:"column:image_alt" json:"image_alt"`
	Price      float64 `json:"price"`
	Promotion  float64 `json:"promotion"`
	Name       string  `gorm:"column:name" json:"name"`
	Desciption string
}

// TableName products
func (Product) TableName() string {
	return "products"
}

// Order struct
type Order struct {
	gorm.Model
	Product
	Customer
	CustomerID   int       `gorm:"column:customer_id"`
	ProductID    int       `gorm:"column:product_id"`
	Price        float64   `gorm:"column:price" json:"sell_price"`
	PurchaseDate time.Time `gorm:"column:purchase_date" json:"purchase_date"`
}

// TableName order
func (Order) TableName() string {
	return "orders"
}
