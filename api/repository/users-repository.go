package repository

import "github.com/akwanmaroso/PengeluaranKu/api/models"

type UserRepository interface {
	Save(models.User) (models.User, error)
}
