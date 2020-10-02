package mysql

import (
	"github.com/akwanmaroso/PengeluaranKu/api/helpers/channels"
	"github.com/akwanmaroso/PengeluaranKu/api/models"
	"github.com/jinzhu/gorm"
)

type repositoryTransactionsMysql struct {
	db *gorm.DB
}

func NewRepositoryTransactionsMysql(db *gorm.DB) *repositoryTransactionsMysql {
	return &repositoryTransactionsMysql{db}
}

func (r *repositoryTransactionsMysql) Save(transaction models.Transaction) (models.Transaction, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.Transaction{}).Create(&transaction).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return transaction, nil
	}
	return models.Transaction{}, nil
}
