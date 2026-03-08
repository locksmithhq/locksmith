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

type createClientHandler struct {
	createClientUseCase contract.CreateClientUseCase
}

// Execute implements [contract.CreateClientHandler].
func (h *createClientHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "project_id")
	if projectId == "" {
		stackerror.HttpResponse(w, "CreateClientHandler", errors.New("project_id is required"))
		return
	}

	var in input.Client
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		stackerror.HttpResponse(w, "CreateClientHandler", errors.New("invalid json"))
		return
	}

	in.ProjectID = projectId

	client, err := h.createClientUseCase.Execute(r.Context(), in)
	if err != nil {
		stackerror.HttpResponse(w, "CreateClientHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

func NewCreateClientHandler(createClientUseCase contract.CreateClientUseCase) contract.CreateClientHandler {
	return &createClientHandler{
		createClientUseCase: createClientUseCase,
	}
}
