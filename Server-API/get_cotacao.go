package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "cotacao.db")
	if err != nil {
		log.Fatalf("Erro ao abrir o banco de dados: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, bid FROM cotacao")
	if err != nil {
		log.Fatalf("Erro ao consultar a tabela: %v", err)
	}
	defer rows.Close()

	fmt.Println("Registros na tabela cotacao:")
	for rows.Next() {
		var id int
		var bid string

		if err := rows.Scan(&id, &bid); err != nil {
			log.Fatalf("Erro ao ler os dados: %v", err)
		}

		fmt.Printf("ID: %d, Bid: %s\n", id, bid)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Erro durante a iteração: %v", err)
	}
}
