package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"bro-code/golang-tutorials/api"
)

func Start() {
	go startExternalHTTPServer()
}

func startExternalHTTPServer() {
	r := mux.NewRouter()

	addTransactionsAPIs(r)

	fmt.Println("Starting External HTTP Server at 8070")
	log.Fatal(http.ListenAndServe(":8070", r))

}

func addTransactionsAPIs(r *mux.Router) {
	r.HandleFunc("/price", api.GetPrice).
		Methods("GET")

}