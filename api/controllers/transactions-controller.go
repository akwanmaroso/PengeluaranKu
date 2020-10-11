package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/akwanmaroso/PengeluaranKu/api/auth"
	"github.com/akwanmaroso/PengeluaranKu/api/database"
	"github.com/akwanmaroso/PengeluaranKu/api/helpers/responses"
	"github.com/akwanmaroso/PengeluaranKu/api/models"
	"github.com/akwanmaroso/PengeluaranKu/api/repository"
	"github.com/akwanmaroso/PengeluaranKu/api/repository/mysql"
	"github.com/gorilla/mux"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	repo := mysql.NewRepositoryTransactionsMysql(db)
	func(transactionsRepository repository.TransactionsRepository) {
		transactions, err := transactionsRepository.FindAll(uint64(uid))
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, transactions)
	}(repo)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer r.Body.Close()

	transaction := models.Transaction{}
	err = json.Unmarshal(body, &transaction)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	transaction.CreatorID = uid

	transaction.Prepare()
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := mysql.NewRepositoryTransactionsMysql(db)
	func(transactionRepository repository.TransactionsRepository) {
		transaction, err = transactionRepository.Save(transaction)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.URL.Path, transaction.ID))
		responses.JSON(w, http.StatusCreated, transaction)
	}(repo)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	if (r).Method == "OPTIONS" {
		return
	}
	vars := mux.Vars(r)
	tid, err := strconv.ParseUint(vars["id"], 10, 64)
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

	repo := mysql.NewRepositoryTransactionsMysql(db)

	func(transactionRepository repository.TransactionsRepository) {
		_, err := transactionRepository.Delete(tid)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Entity", fmt.Sprintf("%d", tid))
		responses.JSON(w, http.StatusNoContent, "")
	}(repo)
}
