package main

import (
	"ego-api/wallet/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func registerHandlers() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/v1/", handlers.Index).Methods("GET")
	router.HandleFunc("/api/v1/wallet/fund", handlers.FundWallet).Methods("POST")
	router.HandleFunc("/api/v1/wallet/transfer/{senderID}/{receiverID}", handlers.TransferFunds).Methods("POST")
	//router.HandleFunc()

	return router
}

func main() {

	router := registerHandlers()

	var port string
	if port = os.Getenv("SERVICE_PORT"); port == "" {
		log.Fatalln("SERVICE_PORT env variable not set")
	} else {
		log.Printf("Listening on port %s", port)
		log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
	}
}
