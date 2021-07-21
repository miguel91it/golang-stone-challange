package main

import (
	"encoding/json"
	"fmt"
)

type Login struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

func NewLoginFromJson(jsonDecoder *json.Decoder) (*Login, error) {

	// cria o objeto de login com os dados vindos do json
	var login Login

	if err := jsonDecoder.Decode(&login); err != nil {

		return &Login{}, fmt.Errorf("error to decode json received to Login object: %s", err.Error())
	}

	login.Secret = HashSecret(login.Secret)

	return &login, nil
}

func (l *Login) Authenticate() error {
	// vai verificar se cpf e secret batem com o armazenado no banco

	accountByCpf := db.FindAccountByCpf(l.Cpf)

	if l.Secret != accountByCpf.Secret {

		return fmt.Errorf("either CPF or Secret is not correct")
	}

	return nil
}
