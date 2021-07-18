package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type Account struct {
	Id         int       `json:"id"`
	Name       string    `json:"name,omitempty"`
	Cpf        string    `json:"cpf,omitempty"`
	Secret     string    `json:"secret,omitempty"`
	Balance    float64   `json:"balance"`
	Created_at time.Time `json:"created_at,omitempty"`
}

type Accounts []Account

func NewAccountFromJson(jsonDecoder *json.Decoder) (*Account, error) {

	var account Account

	if err := jsonDecoder.Decode(&account); err != nil {

		return &Account{}, fmt.Errorf("error to decode json received to Account object: %s", err.Error())
	}

	account.HashSecret()

	return &account, nil
}

func (a *Account) UpdateBalance(ammount float64) (float64, error) {
	// func para atualizar o balance cumulativamente
	a.Balance = a.Balance + ammount

	return a.Balance, nil
}

func (a *Account) CheckBalanceForDebit(ammountToDebit float64) (bool, error) {

	currentBalance := a.Balance

	if ammountToDebit > currentBalance {

		return false, fmt.Errorf("this operation cannot be performed. Current Account Balance '%f' is less than the Ammount to debit %f", currentBalance, ammountToDebit)
	}

	return true, nil
}

func (a *Account) HashSecret() {

	h := sha256.New()

	h.Write([]byte(a.Secret))

	secret_hash := h.Sum(nil)

	a.Secret = hex.EncodeToString(secret_hash)
}
