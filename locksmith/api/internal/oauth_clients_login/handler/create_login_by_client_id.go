package handler

import (
	"encoding/json"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/types/input"
	"github.com/go-chi/chi/v5"
)

type createLoginByClientIDHandler struct {
	createLoginByClientIDUseCase contract.CreateLoginByClientIDUseCase
}

// Execute implements contract.CreateLoginByClientIDHandler.
func (h *createLoginByClientIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "id")

	var in input.Login
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		stackerror.HttpResponse(w, "CreateLoginByClientIDHandler", err)
		return
	}

	if err := h.createLoginByClientIDUseCase.Execute(r.Context(), clientID, in); err != nil {
		stackerror.HttpResponse(w, "CreateLoginByClientIDHandler", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewCreateLoginByClientIDHandler(createLoginByClientIDUseCase contract.CreateLoginByClientIDUseCase) contract.CreateLoginByClientIDHandler {
	return &createLoginByClientIDHandler{
		createLoginByClientIDUseCase: createLoginByClientIDUseCase,
	}
}
