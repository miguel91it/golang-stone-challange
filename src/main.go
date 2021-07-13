package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	mux "github.com/gorilla/mux"
)

type Account struct {
	Id         int       `json:"id"`
	Name       string    `json:"name,omitempty"`
	Cpf        string    `json:"cpf,omitempty"`
	Secret     string    `json:"secret,omitempty"`
	Balance    float64   `json:"balance,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
}

type Accounts []Account

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/accounts", GetAccounts).Methods("GET")

	router.HandleFunc("/accounts/{id}/balance", GetAccountBalance).Methods("GET")

	router.HandleFunc("/accounts", CreateAccount).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {

	accounts := Accounts{
		Account{Id: 1, Name: "miguel", Cpf: "398.291.098-60", Secret: "passwd", Balance: 1.5, Created_at: time.Now()},
	}

	fmt.Print(accounts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		panic(err)
	}
}

func GetAccountBalance(w http.ResponseWriter, r *http.Request) {

}

func CreateAccount(w http.ResponseWriter, r *http.Request) {

}
