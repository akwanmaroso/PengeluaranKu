package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/akwanmaroso/PengeluaranKu/api/auth"
	"github.com/akwanmaroso/PengeluaranKu/api/helpers/responses"
	"github.com/akwanmaroso/PengeluaranKu/api/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	token, err := auth.SigIn(user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}
