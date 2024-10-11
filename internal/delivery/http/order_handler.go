package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/agnaldojpereira/list-order/internal/usecase"
)

type OrderHandler struct {
	listOrdersUseCase usecase.ListOrdersUseCase
}

func NewOrderHandler(listOrdersUseCase usecase.ListOrdersUseCase) *OrderHandler {
	return &OrderHandler{listOrdersUseCase: listOrdersUseCase}
}

func (h *OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	orders, err := h.listOrdersUseCase.Execute(r.Context())
	if err != nil {
		if errors.Is(err, usecase.ErrNotFound) {
			respondWithError(w, http.StatusNotFound, "Nenhum pedido encontrado")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Erro interno do servidor")
		return
	}
		return
	respondWithJSON(w, http.StatusOK, orders)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao processar a resposta"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
