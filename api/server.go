package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akwanmaroso/PengeluaranKu/api/router"
	"github.com/akwanmaroso/PengeluaranKu/auto"
	"github.com/akwanmaroso/PengeluaranKu/config"
	"github.com/rs/cors"
)

// Run server
func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tListening [::]:%d\n", 9000)
	listen(9000)
}

func listen(port int) {
	r := router.New()
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
