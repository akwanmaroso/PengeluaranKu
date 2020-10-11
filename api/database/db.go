package database

import (
	"github.com/akwanmaroso/PengeluaranKu/config"

	// Register driver mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Connect to database
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDriver, config.DBURL)
	if err != nil {
		return nil, err
	}
	return db, err
}
