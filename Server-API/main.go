package main

import (
	"github.com/Douglas-Souza40/Client-Server-API/Client-Server-API/Server-API/internal/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cotacao", handler.GetCotacaoHandler)
	log.Println("Servidor ouvindo na porta 8080.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
