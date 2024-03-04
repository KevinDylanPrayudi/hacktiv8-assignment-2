package database

import (
	"assigment-2/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db  *gorm.DB
	err error
)

func PostGresDB() error {
	// https://github.com/go-gorm/postgres
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=123 dbname=assignment-2 port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	db.AutoMigrate(&models.Order{}, &models.Item{})
	db.Exec("TRUNCATE items, orders;")
	db.Exec("SELECT setval(pg_get_serial_sequence('items', 'id'), COALESCE((SELECT MAX(id) + 1 FROM items), 1), false);")
	db.Exec("SELECT setval(pg_get_serial_sequence('orders', 'id'), COALESCE((SELECT MAX(id) + 1 FROM orders), 1), false);")

	if err == nil {
		fmt.Println("Postgress is running")
	}

	return err
}

func GetDB() *gorm.DB {
	return db
}
