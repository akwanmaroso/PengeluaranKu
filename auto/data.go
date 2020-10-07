package auto

import "github.com/akwanmaroso/PengeluaranKu/api/models"

var users = []models.User{
	{
		Name:     "JohnDoe",
		Email:    "test@email.com",
		Password: "test",
	},
}

var category = []models.Category{
	{
		Name:        "Fee",
		Description: "Tagihan",
		Color:       "#00aabb",
	},
}
