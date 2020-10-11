package controllers

import (
	"fmt"
	"net/http"

	"github.com/akwanmaroso/PengeluaranKu/api/database"
)

// Index controller
func Index(w http.ResponseWriter, r *http.Request) {
	_, err := database.Connect()
	if err != nil {
		fmt.Printf("error connect to db: %s", err)
	}

	_, _ = w.Write([]byte("hello"))
}
