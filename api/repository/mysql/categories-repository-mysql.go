package mysql

import (
	"errors"
	"fmt"
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

func (r *repositoryCategoriesMysql) FindAll(cid uint64) ([]models.Category, error) {
	var err error
	var categories []models.Category
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.Category{}).Where("creator_id = ?", cid).Limit(100).Find(&categories).Error
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

func (r *repositoryCategoriesMysql) Delete(cid uint64) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Model(&models.Category{}).Where("id = ?", cid).Take(&models.Category{}).Delete(&models.Category{})
		fmt.Println(rs)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			if gorm.IsRecordNotFoundError(rs.Error) {
				return 0, errors.New("category not found")
			}
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}
