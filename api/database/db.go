package database

import (
	"github.com/akwanmaroso/PengeluaranKu/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Connect to database
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DbDriver, config.DbURL)
	if err != nil {
		return nil, err
	}
	return db, err
}
