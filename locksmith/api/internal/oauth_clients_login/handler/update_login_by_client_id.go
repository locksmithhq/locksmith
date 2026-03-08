package handler

import (
	"encoding/json"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/types/input"
	"github.com/go-chi/chi/v5"
)

type updateLoginByClientIDHandler struct {
	updateLoginByClientIDUseCase contract.UpdateLoginByClientIDUseCase
}

// Execute implements contract.UpdateLoginByClientIDHandler.
func (h *updateLoginByClientIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "id")

	var in input.Login
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		stackerror.HttpResponse(w, "UpdateLoginByClientIDHandler", err)
		return
	}

	if err := h.updateLoginByClientIDUseCase.Execute(r.Context(), clientID, in); err != nil {
		stackerror.HttpResponse(w, "UpdateLoginByClientIDHandler", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NewUpdateLoginByClientIDHandler(updateLoginByClientIDUseCase contract.UpdateLoginByClientIDUseCase) contract.UpdateLoginByClientIDHandler {
	return &updateLoginByClientIDHandler{
		updateLoginByClientIDUseCase: updateLoginByClientIDUseCase,
	}
}
