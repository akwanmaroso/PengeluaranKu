package database

import (
	"fmt"

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
	fmt.Println("success connect")
	return db, err
}
