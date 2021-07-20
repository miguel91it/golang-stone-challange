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


Documentar aqui como rodar o projeto

o que sao cada endpoint e o que eles retornam, com imagens