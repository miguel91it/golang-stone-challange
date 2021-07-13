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

var accounts Accounts

func main() {

	accounts = Accounts{
		Account{Id: 1, Name: "miguel", Cpf: "398.291.098-60", Secret: "passwd", Balance: 1.5, Created_at: time.Now()},
		Account{Id: 2, Name: "pedro", Cpf: "398.291.098-60", Secret: "passwd", Balance: 1.5, Created_at: time.Now()},
	}

	router := mux.NewRouter()

	router.HandleFunc("/accounts", GetAccounts).Methods("GET")

	router.HandleFunc("/accounts/{id}/balance", GetAccountBalance).Methods("GET")

	router.HandleFunc("/accounts", CreateAccount).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {

	// fmt.Print(accounts)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		panic(err)
	}
}

func GetAccountBalance(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	fmt.Println(params)
	fmt.Println(params["id"])

	fmt.Println(r)

	for _, account := range accounts {

		fmt.Println(account)

		if string(account.Id) == params["id"] {
			json.NewEncoder(w).Encode(account)
			return
		}
	}
	json.NewEncoder(w).Encode(&Account{})

}

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	var newAccount Account

	json.NewDecoder(r.Body).Decode(&newAccount)

	accounts = append(accounts, newAccount)

}
