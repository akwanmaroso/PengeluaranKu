package repository

import "github.com/akwanmaroso/PengeluaranKu/api/models"

type CategoriesRepository interface {
	Save(models.Category) (models.Category, error)
	FindAll(uint64) ([]models.Category, error)
	Delete(uint64) (int64, error)
}
