package main

import (
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

var db Storage

func init() {

	// inicia o banco de dados em memória
	db = NewStorage()

	// inicializa alguns dados de contas no banco de dados
	InitAccounts()

}

func main() {

	// inicializa um router [de rotas]
	router := mux.NewRouter()

	// rotas para contas (Accounts)
	router.HandleFunc("/accounts", GetAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}/balance", GetAccountBalance).Methods("GET")
	router.HandleFunc("/accounts", CreateAccount).Methods("POST")

	// rotas para transferencias entre contas (Transfers)
	router.HandleFunc("/transfers", GetTransfers).Methods("GET")
	router.HandleFunc("/transfers", MakeTransfer).Methods("POST")

	// rota para login
	router.HandleFunc("/login", LoginUser).Methods("POST")

	// sobe o servidor da API na porta fornecida
	log.Fatal(http.ListenAndServe(":16453", router))
}
