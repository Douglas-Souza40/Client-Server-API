# Client-Server-API

## Descrição
Este projeto é uma aplicação Go que consome uma API externa para obter a cotação do dólar (USD-BRL), salva os dados em um banco de dados SQLite e permite consultar os registros salvos.

## Pré-requisitos
Antes de executar a aplicação, certifique-se de ter os seguintes itens instalados:
- [Go](https://golang.org/dl/) (versão 1.24.2 ou superior)
- [SQLite3](https://www.sqlite.org/download.html)

## Configuração do Ambiente
1. Clone o repositório:
   ```bash
   git clone https://github.com/Douglas-Souza40/Client-Server-API.git
   cd Client-Server-API

2. Instale as dependências do projeto:
   ```
   go mod tidy

## Executando a Aplicação
1. **Executar o Server**  
   Execute o servidor que expõe a API:  
   ```bash
   go run ./Server-API/main.go
   ```
2. **Executar o Client da API**  
   ```bash
   go run ./Client-API/main.go
   ```
3. **Executar a Query de Consulta**  
   Execute o programa que consulta e exibe os registros salvos no banco de dados:
    ```bash
    go run ./Server-API/get_cotacao.go
    ```