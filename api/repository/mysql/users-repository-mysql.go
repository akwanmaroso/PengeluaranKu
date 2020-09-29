package mysql

import (
	"github.com/akwanmaroso/PengeluaranKu/api/helpers/channels"
	"github.com/akwanmaroso/PengeluaranKu/api/models"
	"github.com/jinzhu/gorm"
)

type repositoryUsersMysql struct {
	db *gorm.DB
}

// NewRepositoryUsersMysql will create an object that represent the UsersRepository interface
func NewRepositoryUsersMysql(db *gorm.DB) *repositoryUsersMysql {
	return &repositoryUsersMysql{db: db}
}

func (r *repositoryUsersMysql) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}
