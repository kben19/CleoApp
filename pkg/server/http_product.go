package server

import (
	"fmt"
	"log"
	"net/http"
)

func handleGetProduct(w http.ResponseWriter, r *http.Request) {
	log.Print("product get")
	fmt.Fprint(w, "Product Get")
}
