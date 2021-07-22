package main

func InitAccounts() {

	// Cria um slice de contas
	accounts := Accounts{
		Account{Name: "Miguel Lima", Cpf: "398.291.098-60", Secret: "passwd", Balance: 100.51},
		Account{Name: "Pedro Sampaio", Cpf: "123.456.098-42", Secret: "pedro123", Balance: 1.5},
		Account{Name: "Ana Moreno", Cpf: "454.376.098-88", Secret: "aninhaaMoreno", Balance: 0},
		Account{Name: "Gi Silva", Cpf: "195.368.098-70", Secret: "G!silva", Balance: 1000},
	}

	// para cada conta do slice de contas
	for i, account := range accounts {

		// realiza o parse do campo secret para hash
		account.Secret = HashSecret(account.Secret)

		// atualiza a conta com a conta modifica (parsed do secret)
		accounts[i] = account
	}
	// salva no banco de dados em memoria essas contas
	if err := db.SaveAccount(accounts...); err != nil {
		panic(err.Error())
	}
}
