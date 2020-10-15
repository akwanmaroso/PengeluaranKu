package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/akwanmaroso/PengeluaranKu/api/auth"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"

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

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	category.CreatorID = uid

	defer db.Close()

	repo := mysql.NewRepositoryCategoriesMysql(db)
	func(categoriesRepository repository.CategoriesRepository) {
		category, err = categoriesRepository.Save(category)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.URL.Path, category.ID))
		responses.JSON(w, http.StatusCreated, category)
	}(repo)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	cid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	defer db.Close()

	repo := mysql.NewRepositoryCategoriesMysql(db)
	func(categoriesRepository repository.CategoriesRepository) {
		categories, err := categoriesRepository.FindAll(uint64(cid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, categories)
	}(repo)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
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
		_, err := categoriesRepository.Delete(cid)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Entity", fmt.Sprintf("%d", cid))
		responses.JSON(w, http.StatusNoContent, "")
	}(repo)
}
