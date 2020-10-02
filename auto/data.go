package auto

import "github.com/akwanmaroso/PengeluaranKu/api/models"

var users = []models.User{
	{
		Name:     "JohnDoe",
		Email:    "jhondoe@email.com",
		Password: "123456789",
	},
}

var category = []models.Category{
	{
		Name:        "Fee",
		Description: "Tagihan",
		Color:       "#00aabb",
	},
}
