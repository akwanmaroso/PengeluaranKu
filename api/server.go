package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akwanmaroso/PengeluaranKu/api/router"
)

func Run() {
	fmt.Printf("\n\tListening [::]:%d\n", 9000)
	listen(9000)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
