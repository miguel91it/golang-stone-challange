package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// define a interface Storage
type Storage interface {
	SaveAccount(newAccounts ...Account) error
	UpdateAccount(changedAccounts ...Account) error
	SaveTransfer(newTransfers ...Transfer) error
	FindAccount(id int) Account
	FindAccountByCpf(cpf string) Account
	FindTransfers(accountId int) Transfers
	FindAccounts() Accounts
	SaveToken(newTokens ...Token) error
	FindTokens() Tokens
}

// StorageInMemory é o banco de dados em memória. Ele perde todos os dados quando a aplicação é encerrada.
type StorageInMemory struct {
	// slice []Accounts
	accounts Accounts

	// map {account.id: []Transfers}
	transfers map[int]Transfers

	// slice []Tokens
	tokens Tokens
}

/*
	NewStorage cria um novo banco de dados em memoria

	entrada:

	saida:
		- Storage
*/
func NewStorage() Storage {

	// instancia um novo banco de dados em memoria e retorna seu ponteiro
	return &StorageInMemory{
		make(Accounts, 0),
		make(map[int]Transfers),
		make(Tokens, 0),
	}

}

/*
	SaveAccount método de Storage para salvar uma lista de Accounts

	entrada:
		- ...Account

	saida:
		- error
*/
func (s *StorageInMemory) SaveAccount(newAccounts ...Account) error {

	// para cada nova conta na lista de novas contas
	for _, newAccount := range newAccounts {

		// para cada nova conta cadastrada no banco de dados
		for _, accountInDb := range s.accounts {

			// se encontrou no banco uma conta com o CPF da nova conta, da erro
			if newAccount.Cpf == accountInDb.Cpf {

				return fmt.Errorf("account already exists with this cpf: %s", newAccount.Cpf)
			}

		}

		// gera o id da nova conta
		newAccount.Id = len(s.accounts) + 1

		// carimba a data de criação da conta
		newAccount.Created_at = time.Now()

		// adiciona a nova conta na lista de contas ja cadastradas
		s.accounts = append(s.accounts, newAccount)

		// cria uma nova entrada no map de transferencias com o id da nova conta e um slice vazio como valor
		s.transfers[newAccount.Id] = Transfers{}
	}

	return nil
}

/*
	UpdateAccount método de Storage para realizar update no balance de uma Account

	entrada:
		- ...Account

	saida:
		- error
*/
func (s *StorageInMemory) UpdateAccount(changedAccounts ...Account) error {

	// para cada conta modificada na lista passada como parametro
	for _, changedAccount := range changedAccounts {

		// recupera o ponteiro da conta alterada no banco de dados
		accountInDb := &s.accounts[changedAccount.Id-1]

		// atualiza o balance da conta alterada
		accountInDb.Balance = changedAccount.Balance

	}

	return nil
}

/*
	SaveTransfer método de Storage para salvar uma lista de Transfers

	entrada:
		- ...Transfer

	saida:
		- error
*/
func (s *StorageInMemory) SaveTransfer(newTransfers ...Transfer) error {

	// para cada nova transferencia passada no parametro de entrada
	for _, newTransfer := range newTransfers {

		// gera um novo id do tipo UUID
		newTransfer.Id = uuid.NewString()

		// carimba a data deregistro da transferência
		newTransfer.Created_at = time.Now()

		// registra a componente de débito na conta de origem
		s.transfers[newTransfer.Account_origin_id] = append(s.transfers[newTransfer.Account_origin_id], newTransfer)

		// registra a componente de credito na conta de destino
		s.transfers[newTransfer.Account_destination_id] = append(s.transfers[newTransfer.Account_destination_id], Transfer{Id: newTransfer.Id, Account_origin_id: newTransfer.Account_origin_id, Account_destination_id: newTransfer.Account_destination_id, Ammount: -newTransfer.Ammount, Created_at: newTransfer.Created_at})

	}

	return nil
}

/*
	FindAccountByCpf busca uma Account no banco de dados pelo CPF

	entrada:
		- cpf string

	saida:
		- Account
*/
func (s *StorageInMemory) FindAccountByCpf(cpf string) Account {

	// percorre as contas cadastradas no banco de dados
	for _, accountInDB := range s.accounts {

		// se o cpf desejado for igual ao cpf da conta cadastrada
		if cpf == accountInDB.Cpf {

			// retorna a conta cadastrada
			return accountInDB
		}
	}

	// retorna uma conta nula
	return Account{}
}

/*
	FindAccount busca uma Account no banco de dados pelo Id da Account

	entrada:
		- id int

	saida:
		- Account
*/
func (s *StorageInMemory) FindAccount(id int) Account {

	// percorre as contas cadastradas no banco de dados
	for _, accountInDB := range s.accounts {

		// se o id desejado for igual ao id da conta cadastrada
		if id == accountInDB.Id {

			// retorna a conta cadastrada
			return accountInDB
		}
	}

	// retorna uma conta nula
	return Account{}
}

/*
	FindAccounts retorna todas as Accounts em um slice

	entrada:

	saida:
		- Accounts
*/
func (s *StorageInMemory) FindAccounts() Accounts {

	// trecho para realizar um print formatado na saida padrao da lista de contas cadastradas
	formattedAccounts, err := FormatMap(s.accounts)

	if err != nil {
		fmt.Printf("%s", err.Error())
	} else {
		fmt.Printf("\nStorage Accounts: %s\n", formattedAccounts)
	}

	// retorna o slice de contas cadastradas
	return s.accounts
}

/*
	FindTransfers busca todas as transferências registradas para um dado id de conta

	entrada:
		- accountId int

	saida:
		- Transfers
*/
func (s *StorageInMemory) FindTransfers(accountId int) Transfers {

	formattedTransfers, err := FormatMap(s.transfers)

	if err != nil {
		fmt.Printf("%s", err.Error())
	} else {
		fmt.Printf("\nStorage Transfers: %s\n", formattedTransfers)
	}

	// retorna o slice de transferencias da accountid passada
	return s.transfers[accountId]
}

/*
	FindTokens retorna toda a lista de Tokens

	entrada:

	saida:
		- Tokens
*/
func (s *StorageInMemory) FindTokens() Tokens {

	formattedTokens, err := FormatMap(s.tokens)

	if err != nil {
		fmt.Printf("%s", err.Error())
	} else {
		fmt.Printf("\nStorage Tokens: %s\n", formattedTokens)
	}

	// retorna o slice de tokens
	return s.tokens
}

/*
	SaveToken método de Storage para salvar uma lista de tokens

	entrada:
		- ...Token

	saida:
		- error
*/
func (s *StorageInMemory) SaveToken(newTokens ...Token) error {

	// para cada novo token passado no parametro de entrada
	for _, newToken := range newTokens {

		// faz um append dele no slice de tokens do banco de dados em memoria
		s.tokens = append(s.tokens, newToken)
	}

	return nil
}
