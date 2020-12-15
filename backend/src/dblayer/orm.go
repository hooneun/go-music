package dblayer

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

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

// SignInUser user sign in
func (db *DBORM) SignInUser(email, password string) (customer models.Customer, err error) {
	result := db.Table("customers").Where(&models.Customer{
		Email: email,
	})
	err = result.First(&customer).Error
	if err != nil {
		return customer, err
	}
	if !checkPassword(customer.Password, password) {
		return customer, errors.New("Invalid password")
	}

	customer.Password = ""
	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}
	return customer, result.Find(&customer).Error
}

// SignOutUserByID ~
func (db *DBORM) SignOutUserByID(id int) error {
	customer := models.Customer{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	return db.Table("customers").Where(&customer).Update("loggedin", 0).Error
}

// GetCustomerOrdersByID !
func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, err error) {
	return orders, db.Table("orders").Select("*").Joins("JOIN customers ON customers.id = customer_id").Joins("JOIN products ON products.id = product_id").Where("customer_id = ?", id).Scan(&orders).Error
}

// hashPassword !
func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}
	sBytes := []byte(*s)
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	*s = string(hashedBytes[:])
	return nil
}

// checkPassword !
func checkPassword(existingHash, incomingPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPassword)) == nil
}
