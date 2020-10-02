package repository

import "github.com/akwanmaroso/PengeluaranKu/api/models"

type TransactionsRepository interface {
	Save(models.Transaction) (models.Transaction, error)
	FindAll() ([]models.Transaction, error)
}
