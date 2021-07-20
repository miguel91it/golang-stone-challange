package main

func InitAccounts() {

	accounts := Accounts{
		Account{Name: "Miguel Lima", Cpf: "398.291.098-60", Secret: "passwd", Balance: 100.51},
		Account{Name: "Pedro Sampaio", Cpf: "123.456.098-42", Secret: "pedro123", Balance: 1.5},
		Account{Name: "Ana Moreno", Cpf: "454.376.098-88", Secret: "aninhaaMoreno", Balance: 0},
		Account{Name: "Gi Silva", Cpf: "195.368.098-70", Secret: "G!silva", Balance: 1000},
	}

	for i, account := range accounts {

		account.HashSecret()

		accounts[i] = account
	}
	if err := db.SaveAccount(accounts...); err != nil {
		panic(err.Error())
	}
}
