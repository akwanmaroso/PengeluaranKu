package api

import (
	"fmt"
	"log"
	"net/http"
)

func Run() {
	fmt.Printf("\n\tListening [::]:%d\n", 9000)
	listen(9000)
}

func listen(port int) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
