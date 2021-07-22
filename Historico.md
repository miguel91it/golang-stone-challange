# Historico de desenvolvimento do desafio

## Primeiros passos

Li e reli o readme da stone e entendi perfeitamente o que queriam. Em termos de complexidade de negócio, o projeto é tranquilo. O bixo pegou mesmo com a questão de desenvolver testes unitários, já que nunca tive a oportunidade ainda de desenvolver software com testes. Assim, fui atrás do tutorial de aprendizado em go usando testes e iniciei a pasta _**tutorial_tests**_ a fim de ver acontecer o desenvolvimento de um sisteminha usando TDD. 

Foi bem contraintuitivo dado que tenho toda minha carreira de desenvolvimento orientado à lógica direta do software, e não a testes. Consegui entender alguns aspectos do TDD porém não consegui usar essa visão de desenvolvimento como ponto de partida. Assim, resolvi com o pouco tempo disponível, voltar ao método de desenvolvimento tradicional e só então, após os primeiros resultados do desenvolvimento da API, retomar os testes se houver tempo.

Ah, vocês poderão observar nos commits que eu usei o _**git flow**_ e comecei com uma branch chamada _**feature/primeiras-estruturas**_ cujo objetivo era sair com o projeto do zero. Muito provavelmente mudará bastante ainda, acho.

Tentarei ao máximo realizar commits granulares para evidenciar o workflow e minha linha de raciocínio.

## Git flow

Foi usado o git flow para evidenciar as features que foram sendo trabalhadas conforme o proejto evoluia.

Inicialmente foi dificil dividir a **feature/primeiras-estruturas** em features menores porque foi preciso eu ver um pouco de tudo funcionando para saber se o rumo estava certo. Então essa primeira feature agrupou o desenvolvimento de bastante funcionalidade.

Porém, da segunda feature em diante comecei a trabalhar débitos técnicos mapeados ao término da primeira feature e aí foi possível quebrar em features menores.



## Debito Tecnico

* testes

* validação e correção de race conditions

* como eu faria para corrigir algun race conditions obvios