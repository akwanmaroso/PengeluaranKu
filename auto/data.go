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
		Description: "anyting about fee",
		Color:       "#c92432",
	},
	{
		Name:        "General",
		Description: "anyting",
		Color:       "#24c98a",
	},
	{
		Name:        "Study",
		Description: "eg. Pay spp",
		Color:       "#b024c9",
	},
}
