package mysql

import (
	"github.com/akwanmaroso/PengeluaranKu/api/helpers/channels"
	"github.com/akwanmaroso/PengeluaranKu/api/models"
	"github.com/jinzhu/gorm"
)

type repositoryCategoriesMysql struct {
	db *gorm.DB
}

func NewRepositoryCategoriesMysql(db *gorm.DB) *repositoryCategoriesMysql {
	return &repositoryCategoriesMysql{db}
}

func (r *repositoryCategoriesMysql) Save(category models.Category) (models.Category, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.Category{}).Create(&category).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return category, nil
	}
	return models.Category{}, err
}
