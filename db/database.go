package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const dsn = "host=localhost port=5432 user=postgres password=1234 dbname=prices sslmode=disable"

func OpenDb() *gorm.DB {

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
