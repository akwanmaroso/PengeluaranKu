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

func (r *repositoryCategoriesMysql) FindAll() ([]models.Category, error) {
	var err error
	var categories []models.Category
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.Category{}).Limit(100).Find(&categories).Error
		if err != nil {
			ch <- false
			return
		}
		if len(categories) > 0 {
			for i := range categories {
				err = r.db.Debug().Model(&models.User{}).Where("id = ?", categories[i].CreatorID).Find(&categories[i].Creator).Error
				if err != nil {
					ch <- false
					return
				}
			}
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return categories, nil
	}
	return nil, err
}
