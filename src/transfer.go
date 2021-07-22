package main

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

// define o Model Account
type Transfer struct {
	Id                     string    `json:"id"`
	Account_origin_id      int       `json:"account_origin_id,omitempty"`
	Account_destination_id int       `json:"account_destination_id,omitempty"`
	Ammount                float64   `json:"ammount,omitempty"`
	Created_at             time.Time `json:"created_at,omitempty"`
}

// define uma variavel exportável Transfers como sendo um slice de Transfer
type Transfers []Transfer

/*
	NewTransferFromJson cria um objeto do tipo Transfer a partir de um json contendo os mesmos campos do tipo Transfer

	entrada:
		- json.Decoder

	saida:
		- ponteiro para um Transfer
		- error
*/
func NewTransferFromJson(jsonDecoder *json.Decoder, token string) (*Transfer, error) {

	var transfer Transfer

	// tenta decodificar o json da requsiição no objeto Transfer
	if err := jsonDecoder.Decode(&transfer); err != nil {

		return &Transfer{}, fmt.Errorf("error to decode the json to the Transfer object: %s", err.Error())
	}

	// verifica se a conta de destino realmente existe, se nao retorna erro
	if !(transfer.CheckIfAccountDestinationExists()) {

		return &Transfer{}, fmt.Errorf("Account destination does not exists")
	}

	// verifica se o ammount desejado é maior que zero, afinal, quem realiza transf TIRA dinheiro da sua conta para outra, se nao retorna erro
	if !(transfer.CheckIfAmmountIsValid()) {

		return &Transfer{}, fmt.Errorf("ammount desired to transfer is invalid. Provide an ammount greater than zero")
	}

	// recupera a conta de origem, se nao retorna erro
	if err := transfer.FillAccountOriginId(token); err != nil {

		return &Transfer{}, fmt.Errorf("cannot get the account origin from token: %s", err.Error())
	}

	// como a transferência se inicia sempre por quem vai debitar, seta, então, o ammount como negativo
	transfer.Ammount = -transfer.Ammount

	return &transfer, nil
}

/*
	MakeTransfer realiza de fato as etapas de transferência entre contas

	entrada:

	saida:
		- error
*/
func (t *Transfer) MakeTransfer() error {

	// recupera toda a conta de origem
	accountOrigin := db.FindAccount(t.Account_origin_id)

	// recupera toda a conta de destino
	accountDestination := db.FindAccount(t.Account_destination_id)

	// checa se a conta a ser debitada tem o saldo ok para realizar operação
	if err := accountOrigin.checkBalanceForDebit(math.Abs(t.Ammount)); err != nil {

		return fmt.Errorf("error checking the balance for debit.%s", err.Error())
	}

	// atualiza o balance da conta origem
	accountOrigin.UpdateBalance(t.Ammount)

	// credita ammount na conta destino
	accountDestination.UpdateBalance(math.Abs(t.Ammount))

	// persiste no banco o update balance de origem
	if err := db.UpdateAccount(accountOrigin); err != nil {

		return fmt.Errorf("error trying to save the transfer in database: %s\nPlease, retry", err.Error())
	}

	// persiste no banco o update balance de destino
	if err := db.UpdateAccount(accountDestination); err != nil {

		return fmt.Errorf("error trying to save the transfer in database: %s\nPlease, retry", err.Error())
	}

	// salva na conta de origem o registro do debito da transferencia
	if err := db.SaveTransfer(*t); err != nil {

		return fmt.Errorf("error trying to save the transfer in database: %s\nPlease, retry", err.Error())
	}

	// retorna sucesso
	return nil
}

/*
	CheckIfAccountDestinationExists verifica se a conta de destino existe

	entrada:

	saida:
		- bool
*/
func (t *Transfer) CheckIfAccountDestinationExists() bool {

	return checkIfAccountExists(t.Account_destination_id)
}

/*
	checkIfAccountExists verifica se uma determinada conta existe (função não exportável)

	entrada:

	saida:
		- bool
*/
func checkIfAccountExists(accountId int) bool {

	accountFounded := db.FindAccount(accountId)

	return accountId == accountFounded.Id && accountFounded.Id != 0
}

/*
	CheckIfAmmountIsValid verifica se o valor a ser transferido é valido: maior que zero.

	entrada:

	saida:
		- bool
*/
func (t *Transfer) CheckIfAmmountIsValid() bool {

	return t.Ammount > 0
}

/*
	FillAccountOriginId preenche o account_origin_id da transferencia com a conta de origem contida no tokend e acesso

	entrada:
		- token string

	saida:
		- error
*/
func (t *Transfer) FillAccountOriginId(token string) error {

	// recuperar o account origin id a partir do tokend e acesso
	t.Account_origin_id = GetAccountOriginIdFromToken(token)

	return nil
}
