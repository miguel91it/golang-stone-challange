package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func FormatMap(mapVar interface{}) (string, error) {

	mapFormatted, err := json.MarshalIndent(mapVar, "", "  ")

	if err != nil {

		return "", fmt.Errorf("error trying to format map variable: %s", err.Error())
	}

	return string(mapFormatted), nil

}

func HashSecret(secret string) string {

	h := sha256.New()

	h.Write([]byte(secret))

	secret_hash := h.Sum(nil)

	return hex.EncodeToString(secret_hash)
}
