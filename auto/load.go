package auto

import (
	"log"
	"os"

	"github.com/akwanmaroso/PengeluaranKu/api/database"
	"github.com/akwanmaroso/PengeluaranKu/api/models"
)

// Load is migrate data to db
func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().DropTableIfExists(&models.Category{}, &models.User{}, &models.Partner{}, &models.Transaction{}).Error
	if err != nil {
		os.Exit(1)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Category{}, &models.Partner{}, &models.Transaction{}).Error
	if err != nil {
		os.Exit(1)
	}

	defer db.Close()

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			os.Exit(1)
		}

		category[i].CreatorID = users[i].ID
		err = db.Debug().Model(&models.Transaction{}).Create(&category[i]).Error
		if err != nil {
			log.Fatal(err)
		}
	}
}
