# Historico de desenvolvimento do desafio

## Primeiros passos

Li e reli o readme da stone e entendi perfeitamente o que queriam. Em termos de complexidade de negócio, o projeto é tranquilo. O bixo pegou mesmo com a questão de desenvolver testes unitários, já que nunca tive a oportunidade ainda de desenvolver software com testes. Assim, fui atrás do tutorial de aprendizado em go usando testes e iniciei a pasta _**tutorial_tests**_ a fim de ver acontecer o desenvolvimento de um sisteminha usando TDD. 

Foi bem contraintuitivo dado que tenho toda minha carreira de desenvolvimento orientado à lógica direta do software, e não a testes. Consegui entender alguns aspectos do TDD porém não consegui usar essa visão de desenvolvimento como ponto de partida. Assim, resolvi com o pouco tempo disponível, voltar ao método de desenvolvimento tradicional e só então, após os primeiros resultados do desenvolvimento da API, retomar os testes se houver tempo.

Ah, vocês poderão observar nos commits que eu usei o _**git flow**_ e comecei com uma branch chamada _**feature/primeiras-estruturas**_ cujo objetivo era sair com o projeto do zero. Muito provavelmente mudará bastante ainda, acho.

Tentarei ao máximo realizar commits granulares para evidenciar o workflow e minha linha de raciocínio.

## Git flow

Foi usado o git flow para padronizar o trabalho de versionamento e evidenciar as features que foram sendo trabalhadas conforme o projeto evoluia.

Inicialmente foi dificil dividir a **feature/primeiras-estruturas** em features menores porque foi preciso eu ver um pouco de tudo funcionando para saber se o rumo estava certo. Então essa primeira feature agrupou o desenvolvimento de bastante funcionalidade.

Porém, da segunda feature em diante comecei a trabalhar débitos técnicos mapeados ao término da primeira feature e aí foi possível quebrar em features menores.



## Debitos Técnicos

* Testes

Infelizmente não tive tempo hábil apra concretizar sua reliazação. Como disse acima, é uma área difícil apra mim ainda mas garanto que será uma das primeiras temática minhas de estudo passando ou não na seleção. No meu projeto da minha empresa é algo que precisamos muito começar a fazer dado que o projeto fica maior a cada novo dia.

* Validação e correção de race conditions

O projeto tem uma fragilidade de race conditions que eu detectei mas não tive tempo de modificar.

Cada nova requisição que chega ao servidor disparará uma goroutine para processar essa requisição. O problema reside ai porque o banco de dados em memória possui 3 estruturas (accounts [slice], transfers [map] e tokens [slice]) que sofrem escrita. Como há possibilidade de requisições concorrentes a fragilidade para race conditions reside aí nessas 3 estruturas para os casos em que goroutines diferentes tentem escrever simultaneamente.

para corrigir o problema eu criaria 3 channels, um para cada estrutura, em que as solicitações de escrita nas estruturas seria inseridas no channel e uma rotina do outro lado escutaria esses channels recuperando uma solicitação de cada vez e, assim fazendo as escritas seguras nas estruturas.

porém não tive tempo hábil apra temrinar essa modificação e assumi o débito ténico.