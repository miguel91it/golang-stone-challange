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

func (a *Account) UpdateBalance(ammount float64) float64 {
	// atualiza o balance cumulativamente
	a.Balance = a.Balance + ammount

	return a.Balance
}

func (a *Account) checkBalanceForDebit(ammountToDebit float64) error {

	currentBalance := a.Balance

	if ammountToDebit > currentBalance {

		return fmt.Errorf(" current account balance '%f' is less than the ammount to debit %f", currentBalance, ammountToDebit)
	}

	return nil
}

func (a *Account) HashSecret() {

	h := sha256.New()

	h.Write([]byte(a.Secret))

	secret_hash := h.Sum(nil)

	a.Secret = hex.EncodeToString(secret_hash)
}
