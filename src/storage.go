package main

import (
	"fmt"
	"time"
)

type Storage interface {
	SaveAccount(newAccounts ...Account) error
	SaveTransfer(newTransfers ...Transfer) error
	FindAccount(id int) Account
	FindTransfers(accountId int) (Transfers, error)
	FindAccounts() (Accounts, error)
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
	}

	return nil
}

func (s *StorageInMemory) SaveTransfer(newTransfers ...Transfer) error {
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

	return s.accounts
}

func (s *StorageInMemory) FindTransfers(accountId int) (Transfers, error) {
	return s.transfers[accountId], nil
}

func NewStorage() Storage {

	return new(StorageInMemory)

}
