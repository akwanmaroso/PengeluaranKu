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
		Name:        "General",
		Description: "All about your activity",
		Color:       "#46e33d",
	},
}
