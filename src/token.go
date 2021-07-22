package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// define o Model Token
type Token struct {
	Token           string
	Cpf             string
	AccountOriginId int
}

// define uma variavel exportável Tokens como sendo um slice de Token
type Tokens []Token

/*
	NewToken cria um objeto do tipo Token a partir dos parametros fornecidos

	entrada:
		- cpf string
		- accountOriginId int

	saida:
		- ponteiro para um Token
		- error
*/
func NewToken(cpf string, accountOriginId int) (*Token, error) {

	// instancia um novo objeto do tipo Token
	token := &Token{
		Cpf:             cpf,
		AccountOriginId: accountOriginId,
	}

	// chama o método para gerar a string do tokend e acesso
	err := token.GenerateTokenString()

	// se der erro
	if err != nil {

		return &Token{}, fmt.Errorf("%s", err.Error())
	}

	// se nao der erro, retorna o endreço do objeto token
	return token, nil
}

/*
	GenerateTokenString cria um token de acesso e grava-o no objeto Token

	entrada:

	saida:
		- error
*/
func (t *Token) GenerateTokenString() error {

	/*
		Processo de geração de token de acesso retirado do seguinte site:

		É um processo do qual eu não domino, confesso. Não tive tempo de validar se ele está correto.
		Porém em muitos sites o processo é semelhante e por isso assumi, para modo de desenvolvimento e testes, como sendo correto.
		Para produção eu teria que tirar um tempo apra estudá-lo e validá-lo.
	*/

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

/*
	GetAccountOriginIdFromToken recupera o atributo accountOriginId a aprtir de uma string de token de acesso

	entrada:
		- token string

	saida:
		- int
*/
func GetAccountOriginIdFromToken(tokenStr string) int {

	/*
		O mesmo caso da geração do token. Vide comentários acima.
	*/

	claims := jwt.MapClaims{}

	jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("cpf"), nil
	})

	accountOriginId := claims["account_origin_id"].(float64)

	return int(accountOriginId)
}

/*
	AuthorizeToken verifica se um token passado está autorizado a acessar o sistema

	entrada:
		- tokenStr string

	saida:
		- error
*/
func AuthorizeToken(tokenStr string) error {

	// recupera a lista de tokens
	tokensInDb := db.FindTokens()

	// percorre a lsita de tokens e...
	for _, tokenInDb := range tokensInDb {

		// verifica se o token recebido na requisição está cadastrado na lista de tokens
		if tokenStr == tokenInDb.Token {

			// se estiver, então retorna um erro nil sinalizando que esta tudo ok
			return nil
		}
	}

	// se não, retorna um erro e o acesso à API será negado
	return fmt.Errorf("token is not authorized")
}
