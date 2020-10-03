package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akwanmaroso/PengeluaranKu/api/database"
	"github.com/akwanmaroso/PengeluaranKu/api/helpers/responses"
	"github.com/akwanmaroso/PengeluaranKu/api/models"
	"github.com/akwanmaroso/PengeluaranKu/api/repository"
	"github.com/akwanmaroso/PengeluaranKu/api/repository/mysql"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	category := models.Category{}
	err = json.Unmarshal(body, &category)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := mysql.NewRepositoryCategoriesMysql(db)
	func(categoriesRepository repository.CategoriesRepository) {
		category, err := categoriesRepository.Save(category)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.URL.Path, category.ID))
		responses.JSON(w, http.StatusCreated, category)
	}(repo)
}
