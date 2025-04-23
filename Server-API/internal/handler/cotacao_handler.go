package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Douglas-Souza40/Client-Server-API/Client-Server-API/Server-API/internal/service"
	"log"
	"net/http"
	"time"
)

func GetCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()
	cotacao, err := service.GetCotacao(ctx)
	if err != nil {
		// Verifica se o erro é devido ao timeout do contexto
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("Erro: operação excedeu o tempo limite")
			http.Error(w, "Erro: operação excedeu o tempo limite", http.StatusGatewayTimeout)
			return
		}
		log.Println("Erro ao obter cotação:", err)
		http.Error(w, "Erro ao obter cotação", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"bid": cotacao.Bid})

}
