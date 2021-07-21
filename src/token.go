package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	Token           string
	Cpf             string
	AccountOriginId int
}

type Tokens []Token

func NewToken(cpf string, accountOriginId int) (*Token, error) {

	token := &Token{
		Cpf:             cpf,
		AccountOriginId: accountOriginId,
	}

	err := token.GenerateTokenString()

	if err != nil {

		return &Token{}, fmt.Errorf("%s", err.Error())
	}

	return token, nil
}

func (t *Token) GenerateTokenString() error {

	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true

	atClaims["cpf"] = t.Cpf

	atClaims["account_origin_id"] = t.AccountOriginId

	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	tokenStr, err := at.SignedString([]byte(t.Cpf))

	if err != nil {

		return fmt.Errorf("error to generate the access token: %s", err.Error())
	}

	t.Token = tokenStr

	return nil
}

func GetAccountOriginIdFromToken(tokenStr string) int {

	claims := jwt.MapClaims{}

	jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("cpf"), nil
	})

	// fmt.Printf("\nClaims: %+v\n", claims)

	accountOriginId := claims["account_origin_id"].(float64)

	return int(accountOriginId)
}

func AuthorizeToken(tokenStr string) error {

	tokensInDb := db.FindTokens()

	for _, tokenInDb := range tokensInDb {

		if tokenStr == tokenInDb.Token {

			return nil
		}
	}

	return fmt.Errorf("token is not authorized")
}
