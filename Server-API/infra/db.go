package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() (*sql.DB, error) {
	var err error
	once.Do(func() {
		db, err = sql.Open("sqlite3", "cotacao.db")
		if err == nil {
			_, err = db.Exec("CREATE TABLE IF NOT EXISTS cotacao (id INTEGER PRIMARY KEY AUTOINCREMENT, bid TEXT)")
		}
	})
	fmt.Println("Conectando ao banco de dados")
	return db, err
}
