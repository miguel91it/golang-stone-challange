package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Storage interface {
	SaveAccount(newAccounts ...Account) error
	UpdateAccount(changedAccounts ...Account) error
	SaveTransfer(newTransfers ...Transfer) error
	FindAccount(id int) Account
	FindTransfers(accountId int) Transfers
	FindAccounts() Accounts
}

type StorageInMemory struct {
	accounts Accounts

	transfers map[int]Transfers
}

func (s *StorageInMemory) SaveAccount(newAccounts ...Account) error {

	for _, newAccount := range newAccounts {

		for _, accountInDb := range s.accounts {

			if newAccount.Cpf == accountInDb.Cpf {

				return fmt.Errorf("account already exists with this cpf: %s", newAccount.Cpf)
			}

		}

		newAccount.Id = len(s.accounts) + 1

		newAccount.Created_at = time.Now()

		s.accounts = append(s.accounts, newAccount)

		s.transfers[newAccount.Id] = Transfers{}
	}

	return nil
}

func (s *StorageInMemory) UpdateAccount(changedAccounts ...Account) error {

	for _, changedAccount := range changedAccounts {

		accountInDb := &s.accounts[changedAccount.Id-1]

		accountInDb.Balance = changedAccount.Balance

	}

	return nil
}

func (s *StorageInMemory) SaveTransfer(newTransfers ...Transfer) error {
	for _, newTransfer := range newTransfers {

		newTransfer.Id = uuid.NewString()

		newTransfer.Created_at = time.Now()

		// registra a componente de débito na conta de origem
		s.transfers[newTransfer.Account_origin_id] = append(s.transfers[newTransfer.Account_origin_id], newTransfer)

		// registra a componente de credito na conta de destino
		s.transfers[newTransfer.Account_destination_id] = append(s.transfers[newTransfer.Account_destination_id], Transfer{Account_origin_id: newTransfer.Account_origin_id, Account_destination_id: newTransfer.Account_destination_id, Ammount: -newTransfer.Ammount})

	}

	return nil
}

func (s *StorageInMemory) FindAccount(id int) Account {

	for _, accountInDB := range s.accounts {

		if id == accountInDB.Id {

			return accountInDB
		}
	}

	return Account{}
}

func (s *StorageInMemory) FindAccounts() Accounts {

	formattedAccounts, err := FormatMap(s.accounts)

	if err != nil {
		fmt.Printf("%s", err.Error())
	} else {
		fmt.Printf("\nStorage Accounts: %s\n", formattedAccounts)
	}

	return s.accounts
}

func (s *StorageInMemory) FindTransfers(accountId int) Transfers {

	formattedTransfers, err := FormatMap(s.transfers)

	if err != nil {
		fmt.Printf("%s", err.Error())
	} else {
		fmt.Printf("\nStorage Transfers: %s\n", formattedTransfers)
	}

	return s.transfers[accountId]
}

func NewStorage() Storage {

	// return new(StorageInMemory)

	return &StorageInMemory{
		make(Accounts, 0),
		make(map[int]Transfers),
	}

}
