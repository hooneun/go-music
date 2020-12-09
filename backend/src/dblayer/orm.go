package dblayer

import (
	"github.com/hooneun/go-music/backend/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBORM struct
type DBORM struct {
	*gorm.DB
}

// NewORM return DBORM struct
func NewORM(dns string) (*DBORM, error) {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	return &DBORM{
		DB: db,
	}, err
}

// GetAllProducts All products
func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

// GetPromos promotion products
func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

// GetCustomerByName !
func (db *DBORM) GetCustomerByName(name string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{
		Name: name,
	}).Find(&customer).Error
}

// GetCustomerByID !
func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error
}

// GetProduct !
func (db *DBORM) GetProduct(id int) (product models.Product, err error) {
	return product, db.First(&product, id).Error
}

// AddUser Create user
func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	hashPassword(&customer.Password)
	customer.LoggedIn = true

	return customer, db.Create(&customer).Error
}
