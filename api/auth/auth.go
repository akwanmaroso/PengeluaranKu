package auth

import (
	"fmt"

	"github.com/akwanmaroso/PengeluaranKu/api/database"
	"github.com/akwanmaroso/PengeluaranKu/api/helpers/channels"
	"github.com/akwanmaroso/PengeluaranKu/api/models"
	"github.com/akwanmaroso/PengeluaranKu/api/security"
	"github.com/jinzhu/gorm"
)

func SigIn(email, password string) (string, error) {
	user := models.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}

		defer db.Close()

		err = db.Model(models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		err = security.VerifyPassword(user.Password, password)
		if err != nil {
			fmt.Println("error")
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return CreateToken(user.ID)
	}
	return "", err
}
