package handler

import (
	"encoding/json"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/types/input"
	"github.com/go-chi/chi/v5"
)

type updateSignupByClientIDHandler struct {
	updateSignupByClientIDUseCase contract.UpdateSignupByClientIDUseCase
}

// Execute implements contract.UpdateSignupByClientIDHandler.
func (h *updateSignupByClientIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "id")

	var in input.Signup
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		stackerror.HttpResponse(w, "UpdateSignupByClientIDHandler", err)
		return
	}

	if err := h.updateSignupByClientIDUseCase.Execute(r.Context(), clientID, in); err != nil {
		stackerror.HttpResponse(w, "UpdateSignupByClientIDHandler", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewUpdateSignupByClientIDHandler(updateSignupByClientIDUseCase contract.UpdateSignupByClientIDUseCase) contract.UpdateSignupByClientIDHandler {
	return &updateSignupByClientIDHandler{
		updateSignupByClientIDUseCase: updateSignupByClientIDUseCase,
	}
}
