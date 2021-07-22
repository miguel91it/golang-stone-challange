# Stone Challange

Projeto desenvolvido pelo candidato João Miguel, email miguel91_it@hotmail.com, telefone 11 97002399.

A seguir descreverei como o projeto foi estruturado e como fazer para executá-lo.

## Estrutura do Projeto

A API de trsnferência entre contas possui as seguintes entidades (modelos):

    * Account
    * Transfer
    * Login
    * Token 
    * StorageInMemory

### Entidade Account

Esta entidade é responsável por agrupar os dados referentes à uma conta.

    * Id         (int)     -> id da conta
    * Name       (string)  -> Nome do titular da conta
    * Cpf        (string)  -> Cpf do titular da conta
    * Secret     (string)  -> Senha da conta
    * Balance    (float64) -> Saldo atual da conta
    * Created_at (time)    -> Data e hora que a conta foi criada


### Entidade Transfer

Esta entidade é responsável por agrupar os dados referentes a uma transferência entre contas.

    * Id                     (string) -> UUID da transferência
    * Account_origin_id      (int)    -> conta de origem da transferência
    * Account_destination_id (int)    -> conta de destino da transferência
    * Ammount                (float)  -> valor da transferência
    * Created_at             (time)   -> Data e hora que a conta foi criada

Toda vez que uma nova transferência tentar ser realizada, o saldo da conta de origem será avaliado e, se houver saldo maior ou igual ao que se deseja transferir, então será permitido. Do contrário não será permitido e um erro será retornado ao usuário da API.

Quando for possível realizar a transferência entre contas, o sistema irá registrar duas componentes de Transferência no banco de dados: débito e crédito.

Registrará na lista de transferências da conta de origem uma nova transferência com valor negativo, significando débito. E, de forma simétrica, registrará na lista de transferẽncias da conta de destino uma nova transferência com o mesmo valor, porém positivo, isto é, um crédito.

Assim, dessa forma, é possível rastrear o caminho qual o atual salde de uma determinada conta percorreu, entre débitos e créditos.

### Entidade Login

Esta entidade é responsável por agrupar os dados referentes a um login.

    * Cpf    (string) -> cpf da conta
    * secret (string) -> senha da conta

Toda vez que um usuário se logar na API, o cpf e a senha deverão ser fornecidos. O sistema então autenticará o usuário conferindo se para aquele CPF fornecido, a senha fornecida bate com a senha armazenada na conta do usuário. Se sim, então ele será autenticado e autorizado.


### Entidade Token

Esta entidade é responsável por agrupar os dados referentes a um token.

    * Token           (string) -> string do token gerado
    * Cpf             (string) -> cpf da conta
    * AccountOriginId (int)    -> id da conta

Toda vez que um usuário se logar na API, um novo token será gerado, a string do token será inserida no slice de tokens do banco de dados em memória e a string do token será retornada ao usuário para realizar futuras requisições.

### Entidade StorageInMemory

A entidade StorageInMemory chama-se assim porque optei por realizar um banco de dados em memória. A desvantagem dessa abordagem é que ao encerrar o sistema da API, todos os dados serão perdidos, porém, como não é um sistema que necessita manter dados persisitidos para além de uma avaliação, então não haverá problema nisso.

A entidade StorageInMemory armazena 3 estruturas de dados necessárias para a API toda funcionar:

    * accounts  ([]Accounts)         -> um slice de Accounts
    * transfers (map[int][]Transfer) -> um map cujas chaves são os id's de cada conta cadastrada e o valor é um slice de Transfer (array de Transfer)
    * tokens    ([]tokens)           -> um slice de Tokens

Além disso a entidade StorageInMemory implementa uma interface *Storage* com funções típicas de banco de dados como SaveAccount, SaveTransfer, FindToken, etc.

A estrutura accounts armazena todas as contas criadas.

A estrutura transfers armazenará sempre 2 componentes (débito/crédito) para cada transferência realizada.

toda vez que uma nova conta for criada na API, uma chave nova será criada no map de transfer sendo o id dessa nova conta a chave a ser usada no map e um slice vazio será criado como valor dessa chave.

Por exemplo:

    A estrutura transfers possui o seguinte estado no momento:

        { }

    Alguem criou uma nova conta com id 1 e, depois, com id 2. Então duas novas chaves serão criadas em transfers:

        {
            "1": [],
            "2": []
        }

    E assim, sucessivamente.

    No momento em que uma nova transferencia for realizada, por exemplo, da conta 2 para a conta 1, duas componentes de trasnfers serão criadas:

        {
            "1': [
                {
                    "Id": 087c5544-2504-434c-8a78-dd6346879547,
                    "Account_origin_id": 1,
                    "Account_destination_id": 2,
                    "Ammount": 256.67,
                    "Created_at": "2021-07-20T18:36:50.728821618-03:00"
                }
            ],
            "2': [
                {
                    "Id": 087c5544-2504-434c-8a78-dd6346879547,
                    "Account_origin_id": 1,
                    "Account_destination_id": 2,
                    "Ammount": -256.67,
                    "Created_at": "2021-07-20T18:36:50.728821618-03:00"
                }
            ],
        }

    Repare que na conta 2 o valor é negativo enquanto que na conta 1 é positivo: débito e crédito.
        

## Como rodar o projeto

O Projeto esta containerizado em uma imagem GO e, para rodá-lo é preciso seguir duas etapas:

*  Realizar o build da imagem docker contendo o projeto. Para isso, após realizar o clone do repositório em sua máquina, entre na raiz do projeto:

> cd golang-stone-challenge

* Uma vez dentro da raiz do projeto, será preciso realizar o build da imagem docker do projeto. Rode o comando make a seguir:

> make docker-build

* Após a conclusão do build da imagem, basta rodar o container em modo background. Para isso rode o comando make a seguir:

> make docker-run

* Se quiser rodar em foreground, para ver alguns logs do servidor:

> docker-run-attached

* Por fim, para remover o container:

> make docker-remove

Perceba que o Makefile que está na raiz do projeto encapsula os detalhes de build e run do container, para facilitar o processo.

Caso você encontre problemas ocm o uso do comando make, então siga as seguintes etapas. Mas, antes, certifique-se de estar na raiz do projeto clonado.

* Realize o build da imagem:

> docker build -t stone-challenge .

* Rode o container, mantendo-se anexado a ele para visualizar logs na saída padrão:

> docker run -p 8000:8000 stone-challenge

* Ou, rode o container desanexado para que ele rode em background:

> docker run -d -p 16453:16453 stone-challenge


## Como Testar a API

Para testar a API sugiro usar uma ferramenta de linha de comando como `Curl` ou `Wget`, ou uma ferramenta visual como `Postman`. Ou até mesmo algum tipo de automação programada apra bater nos endpoints.

Os endpoints para usar a API são:

> GET: localhost:16453/accounts

```
    Header do Request:

        Authorization: <acces_token>
    
    Exemplo de Header

        Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9....HAiOjE2MjY5MjEzOTd9.KBBJk4gE0HVIeKwPLrfWu2Lu05JcmHZcs7LSG7GRNzY

    Exemplo de Body da Resposta:

        [
            {
                "id": 1,
                "name": "Miguel Lima",
                "cpf": "398.291.098-60",
                "secret": "0d6be69b264717f2dd33652e212b173104b4a647b7c11ae72e9885f11cd312fb",
                "balance": 100.51,
                "created_at": "2021-07-21T23:33:32.556628925-03:00"
            },
            {
                "id": 2,
                "name": "Pedro Sampaio",
                "cpf": "123.456.098-42",
                "secret": "2702cb34ee041711b9df0c67a8d5c9de02110c80e3fc966ba8341456dbc9ef2b",
                "balance": 1.5,
                "created_at": "2021-07-21T23:33:32.556629455-03:00"
            }
        ]
```

> GET: localhost:16453/accounts/{account_id}/balance

```
    Exemplo de Header

        Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9....HAiOjE2MjY5MjEzOTd9.KBBJk4gE0HVIeKwPLrfWu2Lu05JcmHZcs7LSG7GRNzY

    Exemplo de Body da Resposta se a conta existir:

        {
            "Balance": 1000.21
        }

    Exemplo de Body da Resposta se a conta NÃO existir:

        "Account not found"
```

> POST: localhost:16453/accounts

```
    Exemplo de Header

        Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9....HAiOjE2MjY5MjEzOTd9.KBBJk4gE0HVIeKwPLrfWu2Lu05JcmHZcs7LSG7GRNzY

    Exemplo de Body da Requisição:

        {
            "name":"Priscila Maia",
            "cpf":"076.636.543-60",
            "secret":"jhdbdbfhbh@@@$dbdb",
            "balance":1000.21
        }

    Exemplo de Body da Resposta se a conta for criada com sucesso:
        
        "New account created succesfully"
    
    Exemplos de Body da Resposta se a conta NÃO for criada:

        "Error to create the new account: account already exists with this cpf: 076.636.543-60"

        "Error to create new Account: error to decode json received to Account object: invalid character 'n' looking for beginning of object key string"
```

> GET: localhost:16453/transfers

```
    Exemplo de Header

        Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9....HAiOjE2MjY5MjEzOTd9.KBBJk4gE0HVIeKwPLrfWu2Lu05JcmHZcs7LSG7GRNzY

    Exemplo de Body da Resposta (Neste exemplo a conta logada é a conta de Id = 4):

        [
            {
                "id": "be82c878-6ed9-4a3c-af30-77d477dba569",
                "account_origin_id": 4,
                "account_destination_id": 2,
                "ammount": -0.75,
                "created_at": "2021-07-21T23:46:54.088515204-03:00"
            }
        ]
```

> POST: localhost:16453/transfers

```
    Exemplo de Header

        Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9....HAiOjE2MjY5MjEzOTd9.KBBJk4gE0HVIeKwPLrfWu2Lu05JcmHZcs7LSG7GRNzY

    Exemplo de Body da Requisição:

        {
            "account_destination_id": 2,
            "ammount": 0.75
        }

    Exemplo de Body da Resposta se a transferência entre contas for realizada com sucesso:
        
        "Transfer performed succesfully"
    
    Exemplos de Body da Resposta se a transferência entre contas NÃO for realizada:

        "Error to validate the Transfer data: Account destination does not exists"

        "Error to perform the Transfer: error checking the balance for debit. current account balance '999.250000' is less than the ammount to debit 1000.000000"
```

> POST: localhost:16453/login


```
    Exemplo de Body da Requisição:

        {
            "cpf":"076.636.543-60",
            "secret":"jhdbdbfhbh@@@$dbdb"
        }

    Exemplo de Body da Resposta se o login for realizada com sucesso:
        
        "Bearer Token Created: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY20...iMjMyMzN9.BSpW_12Vz8BF0QZO00W5uJoDotz-xstsQ3eyvZlF2b0"
    
    Exemplos de Body da Resposta se o login NÃO for realizado:

        "Not Authenticated: either CPF or Secret is not correct"
```