package main

import (
	"ego-api/loan/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func registerHandlers() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/v1/loans", handlers.RequestLoan).Methods("POST")
	router.HandleFunc("/api/v1/loans/{LoanID}", handlers.UpdateLoan).Methods("PUT")
	router.HandleFunc("/api/v1/loans/{LoanID}", handlers.DeleteLoan).Methods("DELETE")
	router.HandleFunc("/api/v1/loans/purchase/{LoanID}/{LenderID}", handlers.PurchaseLoan).Methods("PUT")
	router.HandleFunc("/api/v1/loans/repay/{LoanID}/{BorrowerID}", handlers.RepayLoan).Methods("PUT")

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
