package server

import (
	"fmt"
	"log"
	"net/http"
)

func RegisterHandler(uc usecases) {
	// Init handler struct
	product := InitHandlerProduct(uc.product)

	// HTTP Handler
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/checking", checkingHandler)
	http.HandleFunc("/product", product.handleGetProduct)
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, CleoApp!!! Github: https://github.com/kben19/CleoApp")
}

func checkingHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("debug_log") == "true" {
		log.Print("Panic: checking")
	}
	_, err := w.Write([]byte("Checking Test...!!"))
	if err != nil {
		log.Print(err)
	}
}
