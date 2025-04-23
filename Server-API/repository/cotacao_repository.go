package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Douglas-Souza40/Client-Server-API/Client-Server-API/Server-API/infra"
	"github.com/Douglas-Souza40/Client-Server-API/Client-Server-API/Server-API/model"
	"time"
)

func SalvarCotacao(ctx context.Context, c model.Cotacao) error {
	ctxDB, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	db, err := infra.GetDB()
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctxDB, "INSERT INTO cotacao (bid) VALUES (?)", c.Bid)
	if err != nil {
		// Verifica se o erro é devido ao timeout do contexto
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("erro: operação excedeu o tempo limite")
		}
		return err
	}
	fmt.Println("Dados salvos com sucesso")
	return nil
}
