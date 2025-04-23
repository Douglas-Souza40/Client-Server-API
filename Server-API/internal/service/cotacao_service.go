package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Douglas-Souza40/Client-Server-API/Client-Server-API/Server-API/model"
	"github.com/Douglas-Souza40/Client-Server-API/Client-Server-API/Server-API/repository"
	"log"
	"net/http"
	"time"
)

func GetCotacao(ctx context.Context) (*model.Cotacao, error) {
	URL := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	ctxAPI, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctxAPI, "GET", URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// Verifica se o erro é devido ao timeout do contexto
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("erro: operação excedeu o tempo limite ao buscar cotação")
		}
		return nil, err
	}

	var result map[string]model.Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	cotacao := result["USDBRL"]

	err = repository.SalvarCotacao(ctx, cotacao)
	if err != nil {
		log.Println("Erro ao salvar cotação:", err)
	}
	return &cotacao, nil
}
