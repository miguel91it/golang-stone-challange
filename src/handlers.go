package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	mux "github.com/gorilla/mux"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	accounts, _ := db.FindAccounts()

	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		panic(err)
	}
}

func GetAccountBalance(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	accounts, _ := db.FindAccounts()

	for _, account := range accounts {

		idToFind, err := strconv.Atoi(params["id"])

		if err != nil {
			w.WriteHeader(http.StatusNotFound)

			fmt.Fprintf(w, "Error trying to find an account with the taken account_id '%s'. Error: %s", params["id"], err.Error())

			return
		}

		if account.Id == idToFind {

			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(struct{ Balance float64 }{account.Balance})

			return
		}
	}
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprintf(w, "Account not found")

}

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	var newAccount Account

	json.NewDecoder(r.Body).Decode(&newAccount)

	newAccount.HashSecret()

	if err := db.SaveAccount(Accounts{newAccount}...); err != nil {

		fmt.Fprintf(w, "Error to create the new account: %s", err.Error())

		return
	}

	fmt.Printf("New account created succesfully")

	fmt.Fprintf(w, "New account created succesfully")

}

func GetTransfers(w http.ResponseWriter, r *http.Request) {

}

func MakeTransfer(w http.ResponseWriter, r *http.Request) {

	// var transfer Transfer
	// var transfer interface{}

	// type transferRequest struct {
	// 	Cpf string `json:"cpf",omityempty`
	// 	Id  int    `json:"id",omityempty`
	// }

	// var transferReq transferRequest

	// json.NewDecoder(r.Body).Decode(&transfer)

	transfer, err := NewTransferFromJson(json.NewDecoder(r.Body))

	if err != nil {

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "Error to perform the Transfer: %s", err.Error())

		return
	}

	fmt.Printf("\nBody r: %v", transfer)

	// idDestination := transferReq.Id

	// fmt.Printf("\transfer Requeset Id destiantion: %d", idDestination)

	// if transferReq.Id == 0 && transferReq.Cpf == "" {
	// 	w.WriteHeader(http.StatusBadRequest)

	// 	fmt.Fprint(w, "it's necessary provide either the cpf or account id of the destination account to perform the transfer")

	// 	return
	// }

	// TODO: busca a conta miguel forçadamente, depois rever isso
	accountOrigin := db.FindAccount(1)

	// accountDestination := db.FindAccount(transferReq.Id)

	fmt.Printf("\nAccount origin: %v", accountOrigin)

	// fmt.Printf("\nAccount destination: %v", accountDestination)

	// if (Account{} == accountDestination) {

	// 	w.WriteHeader(http.StatusNotFound)

	// 	fmt.Fprintf(w, "Account Destination provided does not exists.")

	// 	return

	// }

}
