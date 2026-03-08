package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/types/input"
	"github.com/go-chi/chi/v5"
)

type updateClientHandler struct {
	useCase contract.UpdateClientUseCase
}

// Execute implements [contract.UpdateClientHandler].
func (h *updateClientHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")
	clientID := chi.URLParam(r, "id")

	if projectID == "" || clientID == "" {
		stackerror.HttpResponse(w, "UpdateClientHandler", errors.New("project_id or client_id is empty"))
		return
	}

	var in input.Client
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		stackerror.HttpResponse(w, "UpdateClientHandler", errors.New("invalid request body"))
		return
	}

	out, err := h.useCase.Execute(r.Context(), projectID, clientID, in)
	if err != nil {
		stackerror.HttpResponse(w, "UpdateClientHandler", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(out)
}

func NewUpdateClientHandler(useCase contract.UpdateClientUseCase) contract.UpdateClientHandler {
	return &updateClientHandler{useCase: useCase}
}
