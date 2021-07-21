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

	w.WriteHeader(http.StatusOK)

	accounts := db.FindAccounts()

	if err := json.NewEncoder(w).Encode(accounts); err != nil {

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "error to encode accounts list to return to the API caller: %s", err.Error())

		return
	}

}

func GetAccountBalance(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	accounts := db.FindAccounts()

	for _, account := range accounts {

		idToFind, err := strconv.Atoi(params["id"])

		if err != nil {
			w.WriteHeader(http.StatusNotFound)

			fmt.Fprintf(w, "Error trying to find an account with the taken account_id '%s'. Error: %s", params["id"], err.Error())

			return
		}

		if account.Id == idToFind {

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(struct{ Balance float64 }{account.Balance})

			return
		}
	}
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprintf(w, "Account not found")

}

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	newAccount, err := NewAccountFromJson(json.NewDecoder(r.Body))

	if err != nil {

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "Error to create new Account: %s", err.Error())

		return
	}

	if err := db.SaveAccount(Accounts{*newAccount}...); err != nil {

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintf(w, "Error to create the new account: %s", err.Error())

		return
	}

	w.WriteHeader(http.StatusOK)

	fmt.Printf("\nNew account created succesfully\n")

	fmt.Fprintf(w, "New account created succesfully")

}

func GetTransfers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	// TODO: mudar isso depois ara peagr o id da conta logada por mei od token
	loggedAccount := 1

	transfers := db.FindTransfers(loggedAccount)

	if err := json.NewEncoder(w).Encode(transfers); err != nil {

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "error to encode transfers list to return to the API caller: %s", err.Error())

		return
	}

}

func MakeTransfer(w http.ResponseWriter, r *http.Request) {

	transfer, err := NewTransferFromJson(json.NewDecoder(r.Body))

	if err != nil {

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "Error to validate the Transfer data: %s", err.Error())

		return
	}

	if err := transfer.MakeTransfer(); err != nil {

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "Error to perform the Transfer: %s", err.Error())

		return
	}

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Transfer performed succesfully")
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	login, err := NewLoginFromJson(json.NewDecoder(r.Body))

	if err != nil {

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "Error to validate the Login data: %s", err.Error())

		return
	}

	err = login.Authenticate()

	if err != nil {

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "Not Authenticated: %s", err.Error())

		return
	}

	accountOrigin := db.FindAccountByCpf(login.Cpf)

	token, err := NewToken(login.Cpf, accountOrigin.Id)

	if err != nil {

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "Token not created: %s", err.Error())

		return
	}

	if err := db.SaveToken(*token); err != nil {

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintf(w, "Error to save the new token: %s", err.Error())

		return
	}

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Bearer Token Created: %s", token.Token)

}

func GetTokenFromHeader(header http.Header) (string, error) {

	tokenFromHeader := header.Get("Authorization")

	if tokenFromHeader == "" {
		return "", fmt.Errorf("no acces token was provided in the request header")
	}

	return tokenFromHeader, nil
}

func CheckIfIsValidToken(header http.Header) (string, error) {

	token, err := GetTokenFromHeader(header)

	if err != nil {

		return "", fmt.Errorf("invalid token: %s", err.Error())
	}

	if err := AuthorizeToken(token); err != nil {

		return "", fmt.Errorf("token provided is not allowed to access resources. Please, login again and send the new token received: %s", err.Error())
	}

	fmt.Printf("\nToken provided is valid and Authorized\n")

	return token, nil
}
