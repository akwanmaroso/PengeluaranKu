package repository

import "github.com/akwanmaroso/PengeluaranKu/api/models"

type CategoriesRepository interface {
	Save(models.Category) (models.Category, error)
}
