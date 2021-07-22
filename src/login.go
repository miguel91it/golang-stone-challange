package main

import (
	"encoding/json"
	"fmt"
)

// define o Model Login
type Login struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

/*
	NewLoginFromJson cria um objeto do tipo Login a partir de um json contendo os mesmos campos do tipo Login

	entrada:
		- json.Decoder

	saida:
		- ponteiro para um Login
		- error
*/
func NewLoginFromJson(jsonDecoder *json.Decoder) (*Login, error) {

	// cria o objeto de login com os dados vindos do json
	var login Login

	// tenta decodificar o json de login vindo como body da requisição
	if err := jsonDecoder.Decode(&login); err != nil {

		return &Login{}, fmt.Errorf("error to decode json received to Login object: %s", err.Error())
	}

	// realiza o aprse da senha para o formato Hash
	login.Secret = HashSecret(login.Secret)

	return &login, nil
}

/*
	Authenticate autentica que o usuario que esta tentando se logar é ele mesmo.

	É um método do model Login e não precisa de parâmetros de entrada porque usa os atributos do Model.

	entrada:

	saida:
		- error
*/
func (l *Login) Authenticate() error {

	// vai verificar se cpf e secret batem com o armazenado no banco
	accountByCpf := db.FindAccountByCpf(l.Cpf)

	if l.Secret != accountByCpf.Secret {

		return fmt.Errorf("either CPF or Secret is not correct")
	}

	return nil
}
