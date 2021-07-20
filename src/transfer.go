package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Transfer struct {
	Id                     string    `json:"id"`
	Account_origin_id      int       `json:"account_origin_id,omitempty"`
	Account_destination_id int       `json:"account_destination_id,omitempty"`
	Ammount                float64   `json:"ammount,omitempty"`
	Created_at             time.Time `json:"created_at,omitempty"`
}

type Transfers []Transfer

func NewTransferFromJson(jsonDecoder *json.Decoder) (*Transfer, error) {

	var transfer Transfer

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
	if err := transfer.FillAccountOriginId(); err != nil {

		return &Transfer{}, fmt.Errorf("cannot get the account origin from token: %s", err.Error())
	}

	transfer.Ammount = -transfer.Ammount

	return &transfer, nil
}

func (t *Transfer) MakeTransfer(destination_account string) error {

	// descobre a conta de origem extraindo do token

	// verifica se a conta de origem tem saldo para fazer a transferencia

	// debita ammount da conta origem
	// atualiza o balance da conta origem
	// credita ammount na conta destino
	// atualiza o balance da conta destino

	// retorna sucesso
	return nil
}

func (t *Transfer) CheckIfAccountDestinationExists() bool {

	return checkIfAccountExists(t.Account_destination_id)
}

func checkIfAccountExists(accountId int) bool {

	accountFounded := db.FindAccount(accountId)

	return accountId == accountFounded.Id && accountFounded.Id != 0
}

func (t *Transfer) CheckIfAmmountIsValid() bool {

	return t.Ammount > 0
}

func (t *Transfer) FillAccountOriginId() error {

	// TODO: recuperar o account origin id (inicialmente fixarei o account origin id como sendo 1, porem depois terei que refatorar para buscar essa info no token)
	t.Account_origin_id = getAccountFromToken()

	return nil
}

func getAccountFromToken() int {
	return 1
}
