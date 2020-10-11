package repository

import "github.com/akwanmaroso/PengeluaranKu/api/models"

type TransactionsRepository interface {
	Save(models.Transaction) (models.Transaction, error)
	FindAll(uid uint64) ([]models.Transaction, error)
	Delete(tid uint64) (int64, error)
}
