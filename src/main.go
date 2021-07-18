package main

import (
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

var db Storage

func init() {

	// inicia o repositorioinmemory
	db = NewStorage()

	accounts := Accounts{
		Account{Name: "miguel", Cpf: "398.291.098-60", Secret: "passwd", Balance: 1.5},
		Account{Name: "pedro", Cpf: "123.456.098-60", Secret: "passwd", Balance: 1.5},
	}

	if err := db.SaveAccount(accounts...); err != nil {
		panic(err.Error())
	}

}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/accounts", GetAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}/balance", GetAccountBalance).Methods("GET")
	router.HandleFunc("/accounts", CreateAccount).Methods("POST")

	router.HandleFunc("/transfer", GetTransfers).Methods("GET")
	router.HandleFunc("/transfer", MakeTransfer).Methods("POST")

	// router.HandleFunc("/login", Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
