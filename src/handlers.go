package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	mux "github.com/gorilla/mux"
)

/*
	GetAccounts Handler retorna a lista de contas cadastradas no banco de dados
*/
func GetAccounts(w http.ResponseWriter, r *http.Request) {

	// verifica se o token é valido para ser usado
	_, err := CheckIfIsValidToken(r.Header)

	// se nao for valido, então terá erro
	if err != nil {

		// seta um header de negação de acesso
		w.WriteHeader(http.StatusForbidden)

		// retorna o erro no corpo da resposta
		fmt.Fprint(w, err.Error())

		return
	}

	// seta um header de tipo de conteudo da resposta para json
	w.Header().Set("Content-Type", "application/json")

	// seeta um header de OK para a resposta
	w.WriteHeader(http.StatusOK)

	// recupera a listagem de contas cadastradas
	accounts := db.FindAccounts()

	// tenta enviar na resposta a lista de contas
	if err := json.NewEncoder(w).Encode(accounts); err != nil {

		// se der erro
		// seta um status negativo no header
		w.WriteHeader(http.StatusNotAcceptable)

		// retorna o erro no corpo da resposta
		fmt.Fprintf(w, "error to encode accounts list to return to the API caller: %s", err.Error())

		return
	}

}

/*
	GetAccountBalance Handler retorna o balance de uma conta solicitada na URL da requisição
*/
func GetAccountBalance(w http.ResponseWriter, r *http.Request) {

	// verifica se o token é valido para ser usado
	_, err := CheckIfIsValidToken(r.Header)

	if err != nil {

		// se der erro
		// seta um status negativo no header
		w.WriteHeader(http.StatusForbidden)

		// retorna o erro no corpo da resposta
		fmt.Fprint(w, err.Error())

		return
	}

	// recupera a lista de parametros passados na requisição
	params := mux.Vars(r)

	// recupera a lista de contas cadastradas no banco de dados
	accounts := db.FindAccounts()

	// para cada conta da lista de contas
	for _, account := range accounts {

		// converte o id requisitado para inteiro
		idToFind, err := strconv.Atoi(params["id"])

		if err != nil {

			// se der erro
			// seta um status negativo no header
			w.WriteHeader(http.StatusNotFound)

			// retorna o erro no corpo da resposta
			fmt.Fprintf(w, "Error trying to find an account with the taken account_id '%s'. Error: %s", params["id"], err.Error())

			return
		}

		// se a conta solicitada existe no banco de dados
		if account.Id == idToFind {

			// seta um header de tipo de conteudo da resposta para json
			w.Header().Set("Content-Type", "application/json")

			// seeta um header de OK para a resposta
			w.WriteHeader(http.StatusOK)

			// envia na resposta a o saldo da conta
			json.NewEncoder(w).Encode(struct{ Balance float64 }{account.Balance})

			return
		}
	}

	// seta um status não encontrado
	w.WriteHeader(http.StatusNotFound)

	// retorna uma mensagem de conta não encontrada na resposta da requisição
	fmt.Fprintf(w, "Account not found")

}

/*
	CreateAccount Handler cria uma nova conta com os dados passados no body da requisição
*/
func CreateAccount(w http.ResponseWriter, r *http.Request) {

	// verifica se o token é valido para ser usado
	_, err := CheckIfIsValidToken(r.Header)

	if err != nil {

		// se der erro
		// seta um status negativo no header
		w.WriteHeader(http.StatusForbidden)

		// retorna o erro no corpo da resposta
		fmt.Fprint(w, err.Error())

		return
	}

	// cria um novo model account com os dados passados na requisição
	newAccount, err := NewAccountFromJson(json.NewDecoder(r.Body))

	if err != nil {

		// se der erro
		// seta um status negativo no header
		w.WriteHeader(http.StatusNotAcceptable)

		// retorna o erro no corpo da resposta
		fmt.Fprintf(w, "Error to create new Account: %s", err.Error())

		return
	}

	// tenta salvar o novo model no banco de dados
	if err := db.SaveAccount(Accounts{*newAccount}...); err != nil {

		// se der erro
		// seta um status negativo no header
		w.WriteHeader(http.StatusBadRequest)

		// retorna o erro no corpo da resposta
		fmt.Fprintf(w, "Error to create the new account: %s", err.Error())

		return
	}

	// seeta um header de OK para a resposta
	w.WriteHeader(http.StatusOK)

	// imprime na saida padrao do servidor o sucesso na criação da conta
	fmt.Printf("\nNew account created succesfully\n")

	// retorna na resposta da requisição o sucesso na criação da conta
	fmt.Fprintf(w, "New account created succesfully")

}

/*
	GetTransfers Handler retorna a lista de transferencias registradas na conta logada
*/
func GetTransfers(w http.ResponseWriter, r *http.Request) {

	// verifica se o token é valido para ser usado
	_, err := CheckIfIsValidToken(r.Header)

	if err != nil {

		// se der erro
		// seta um status negativo no header
		w.WriteHeader(http.StatusForbidden)

		// retorna o erro no corpo da resposta
		fmt.Fprint(w, err.Error())

		return
	}

	// seta um header de tipo de conteudo da resposta para json
	w.Header().Set("Content-Type", "application/json")

	// seeta um header de OK para a resposta
	w.WriteHeader(http.StatusOK)

	// recupera o token a partir do header da requisição
	token, _ := GetTokenFromHeader(r.Header)

	// recupera a conta que se logou a partir do token
	loggedAccount := GetAccountOriginIdFromToken(token)

	// recupera a lista de transferencias registradas na conta logada
	transfers := db.FindTransfers(loggedAccount)

	// tenta retornar na resposta da requisição a lista de transferencias da conta logada
	if err := json.NewEncoder(w).Encode(transfers); err != nil {

		// se der erro
		// seta um status negativo no header
		w.WriteHeader(http.StatusNotAcceptable)

		// retorna o erro no corpo da resposta
		fmt.Fprintf(w, "error to encode transfers list to return to the API caller: %s", err.Error())

		return
	}

}

/*
	MakeTransfer Handler registra uma nova transação de transferencia entre contas nas contas de origem e destino
*/
func MakeTransfer(w http.ResponseWriter, r *http.Request) {

	// tenta recuperar o token a partir do header da requisição
	token, err := CheckIfIsValidToken(r.Header)

	if err != nil {

		// se der erro
		// seta um status negativo no header
		w.WriteHeader(http.StatusForbidden)

		// retorna o erro no corpo da resposta
		fmt.Fprint(w, err.Error())

		return
	}

	// cria um novo model Transfer com os dados obtidos do token e do body da requisição
	transfer, err := NewTransferFromJson(json.NewDecoder(r.Body), token)

	if err != nil {

		// se der erro
		// seta um status negativo no header
		w.WriteHeader(http.StatusNotAcceptable)

		// retorna o erro no corpo da resposta
		fmt.Fprintf(w, "Error to validate the Transfer data: %s", err.Error())

		return
	}

	// realiza o processo de transferencia entre as contas
	if err := transfer.MakeTransfer(); err != nil {

		// se der erro
		// seta um status negativo no header
		w.WriteHeader(http.StatusNotAcceptable)

		// retorna o erro no corpo da resposta
		fmt.Fprintf(w, "Error to perform the Transfer: %s", err.Error())

		return
	}

	// seeta um header de OK para a resposta
	w.WriteHeader(http.StatusOK)

	// retorna na resposta da requisição o sucesso na realização da transferencia
	fmt.Fprintf(w, "Transfer performed succesfully")
}

/*
	LoginUser Handler realiza o login de um usuario
*/
func LoginUser(w http.ResponseWriter, r *http.Request) {

	// tenta criar um novo model de Login com os dados obtidos do body da requisição
	login, err := NewLoginFromJson(json.NewDecoder(r.Body))

	if err != nil {

		// se der erro
		// seta um status negativo no header
		w.WriteHeader(http.StatusNotAcceptable)

		// retorna o erro no corpo da resposta
		fmt.Fprintf(w, "Error to validate the Login data: %s", err.Error())

		return
	}

	// tenta autenticar o usuario e senha passados
	err = login.Authenticate()

	if err != nil {

		// se der erro
		// seta um status negativo no header
		w.WriteHeader(http.StatusNotAcceptable)

		// retorna o erro no corpo da resposta
		fmt.Fprintf(w, "Not Authenticated: %s", err.Error())

		return
	}

	// recupera a conta cadastrada do usuario que se logou
	accountOrigin := db.FindAccountByCpf(login.Cpf)

	// gera um novo token de acesso apra esse usuario
	token, err := NewToken(login.Cpf, accountOrigin.Id)

	if err != nil {

		// se der erro
		// seta um status negativo no header
		w.WriteHeader(http.StatusNotAcceptable)

		// retorna o erro no corpo da resposta
		fmt.Fprintf(w, "Token not created: %s", err.Error())

		return
	}

	// tenta gravar esse token no banco de dados
	if err := db.SaveToken(*token); err != nil {

		// se der erro
		// seta um status negativo no header
		w.WriteHeader(http.StatusBadRequest)

		// retorna o erro no corpo da resposta
		fmt.Fprintf(w, "Error to save the new token: %s", err.Error())

		return
	}

	// seeta um header de OK para a resposta
	w.WriteHeader(http.StatusOK)

	// retorna como resposta da requisição de login um objeto json com o token de acesso
	fmt.Fprintf(w, "Bearer Token Created: %s", token.Token)

}

/*
	GetTokenFromHeader Handler recupera o token de acesso do header da requisição
*/
func GetTokenFromHeader(header http.Header) (string, error) {

	// tenta recuperar a string do token de acesso no atributo Authorization do header da requisição
	tokenFromHeader := header.Get("Authorization")

	// se o token não foi fornecido
	if tokenFromHeader == "" {

		// retorna o erro
		return "", fmt.Errorf("no acces token was provided in the request header")
	}

	// retorna o token fornecido
	return tokenFromHeader, nil
}

/*
	CheckIfIsValidToken Handler verifica se o token recuperado do header é valido, isto é, já foi autorizado a ser usado nas requisções
*/
func CheckIfIsValidToken(header http.Header) (string, error) {

	// tenta recuperar o token de acesso a partir do header da requisição
	token, err := GetTokenFromHeader(header)

	if err != nil {

		// retorna o erro
		return "", fmt.Errorf("invalid token: %s", err.Error())
	}

	// verifica se o token fornecido está de fato autorizado a ser usado
	if err := AuthorizeToken(token); err != nil {

		// retorna o erro
		return "", fmt.Errorf("token provided is not allowed to access resources. Please, login again and send the new token received: %s", err.Error())
	}

	// fmt.Printf("\nToken provided is valid and Authorized\n")

	// retorna o token
	return token, nil
}
