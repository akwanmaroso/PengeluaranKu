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

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := mysql.NewRepositoryTransactionsMysql(db)
	func(transactionsRepository repository.TransactionsRepository) {
		transactions, err := transactionsRepository.FindAll()
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

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := mysql.NewRepositoryTransactionsMysql(db)
	func(transactionRepository repository.TransactionsRepository) {
		transaction, err := transactionRepository.Save(transaction)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.URL.Path, transaction.ID))
		responses.JSON(w, http.StatusCreated, transaction)
	}(repo)

}
