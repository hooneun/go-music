package dblayer

import (
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
