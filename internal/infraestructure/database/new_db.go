package database

import (
	"os"

	"github.com/Guilbritto/cash-api/internal/domain/transaction"
	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := os.Getenv("CONNECTION_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail to connecto to database")
	}

	db.AutoMigrate(&transaction.Transaction{})

	return db
}
