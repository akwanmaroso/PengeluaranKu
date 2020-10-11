package mysql

import (
	"errors"

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

func (r *repositoryTransactionsMysql) FindAll(uid uint64) ([]models.Transaction, error) {
	var err error
	var transactions []models.Transaction
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Transaction{}).Where("creator_id = ?", uid).Limit(100).Find(&transactions).Error
		if err != nil {
			ch <- false
			return
		}
		if len(transactions) > 0 {
			for i := range transactions {
				err = r.db.Debug().Model(&models.Transaction{}).Where("id = ?", transactions[i].CreatorID).Find(&transactions[i].Creator).Error
				if err != nil {
					ch <- false
					return
				}

				err = r.db.Debug().Model(&models.Transaction{}).Where("id = ?", transactions[i].CategoryID).Find(&transactions[i].Category).Error
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

func (r *repositoryTransactionsMysql) Delete(tid uint64) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Transaction{}).Where("id = ?", tid).Take(&models.Transaction{}).Delete(&models.Transaction{})
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			if gorm.IsRecordNotFoundError(rs.Error) {
				return 0, errors.New("transactions not found")
			}
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}
