package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

/*
	FormatMap recebe um tipo map e reotnra uma string formatada para uma exibição identada na saida padrão

	entrada:
		- mapVar interface{}

	saida:
		- string, error
*/
func FormatMap(mapVar interface{}) (string, error) {

	mapFormatted, err := json.MarshalIndent(mapVar, "", "  ")

	if err != nil {

		return "", fmt.Errorf("error trying to format map variable: %s", err.Error())
	}

	return string(mapFormatted), nil

}

/*
	HashSecret realiza o parse de uma string para Hash usando o metodo SHA256

	entrada:
		- secret string

	saida:
		- string
*/
func HashSecret(secret string) string {

	h := sha256.New()

	h.Write([]byte(secret))

	secret_hash := h.Sum(nil)

	return hex.EncodeToString(secret_hash)
}
