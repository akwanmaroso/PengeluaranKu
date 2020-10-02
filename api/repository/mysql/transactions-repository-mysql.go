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

func (r *repositoryTransactionsMysql) FindAll() ([]models.Transaction, error) {
	var err error
	var transactions []models.Transaction
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Transaction{}).Limit(100).Find(&transactions).Error
		if err != nil {
			ch <- false
			return
		}
		if len(transactions) > 0 {
			for i, _ := range transactions {
				err = r.db.Debug().Model(&models.Transaction{}).Where("id = ?", transactions[i].CreatorID).Find(&transactions[i].Creator).Error
				if err != nil {
					ch <- false
					return
				}
			}
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return transactions, nil
	}
	return nil, err
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
	return models.Transaction{}, err
}
