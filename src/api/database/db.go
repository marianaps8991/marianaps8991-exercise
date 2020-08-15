package database

import (
	"config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.DB_URL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
