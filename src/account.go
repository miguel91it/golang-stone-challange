package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

// define o Model Account
type Account struct {
	Id         int       `json:"id"`
	Name       string    `json:"name,omitempty"`
	Cpf        string    `json:"cpf,omitempty"`
	Secret     string    `json:"secret,omitempty"`
	Balance    float64   `json:"balance"`
	Created_at time.Time `json:"created_at,omitempty"`
}

// define uma variavel exportável Accounts como sendo um slice de Account
type Accounts []Account

/*
	NewAccountFromJson cria um objeto do tipo Account a partir de um json contendo os mesmos campos do tipo Account

	entrada:
		- json.Decoder

	saida:
		- ponteiro para uma Account
		- error
*/
func NewAccountFromJson(jsonDecoder *json.Decoder) (*Account, error) {

	// define uma variavel do tipo Account
	var account Account

	// tenta decodificar o JSON chegando e armazená-lo na variavel Account
	if err := jsonDecoder.Decode(&account); err != nil {

		return &Account{}, fmt.Errorf("error to decode json received to Account object: %s", err.Error())
	}

	// realiza o parse do campo Secret da conta para um Hash SHA256 para armazenar a senha com mais segurança
	account.HashSecret()

	return &account, nil
}

/*
	UpdateBalance atualiza o Balance de uma Account com o valor (ammount) fornecido e retorna o Balance atualizado

	entrada:
		- ammount float64

	saida:
		- Balance float64
*/
func (a *Account) UpdateBalance(ammount float64) float64 {
	// atualiza o balance cumulativamente
	a.Balance = a.Balance + ammount

	return a.Balance
}

/*
	checkBalanceForDebit verifica se o Balance atual da conta é suficiente para realizar um débito (uma retirada) de um determinado valor (ammountToDebit).
	Se for suficiente retorna nil, se for insuficiente retornará um erro.

	entrada:
		- ammountToDebit float64

	saida:
		- error
*/
func (a *Account) checkBalanceForDebit(ammountToDebit float64) error {

	// recupera o atual Balance da conta
	currentBalance := a.Balance

	// se o valor a ser debitado for maior que o saldo atual da conta, então gera um erro
	if ammountToDebit > currentBalance {

		return fmt.Errorf(" current account balance '%f' is less than the ammount to debit %f", currentBalance, ammountToDebit)
	}

	return nil
}

/*
	TODO: deletar esa função aqui porque agora ela esta em utils
*/
func (a *Account) HashSecret() {

	h := sha256.New()

	h.Write([]byte(a.Secret))

	secret_hash := h.Sum(nil)

	a.Secret = hex.EncodeToString(secret_hash)
}
