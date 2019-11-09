package config

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB pointer to connection
var DB *gorm.DB

// InitDB initializes connection to the DB
func InitDB() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=moneyrate_dev")
	if err != nil {
		// panic("I can't establish connection to the DB")
		panic(err.Error())
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetConnMaxLifetime(time.Hour)
}
