package handler

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth2/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth2/types/input"
)

type registerHandler struct {
	registerUseCase contract.RegisterUseCase
}

func (h *registerHandler) Execute(w http.ResponseWriter, r *http.Request) {
	var req input.Register
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.registerUseCase.Execute(r.Context(), req)
	if err != nil {
		stackerror.HttpResponse(w, "RegisterHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func NewRegisterHandler(registerUseCase contract.RegisterUseCase) contract.RegisterHandler {
	return &registerHandler{registerUseCase: registerUseCase}
}
