package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/contract"
	"github.com/go-chi/chi/v5"
)

type getClientByIDAndProjectIDHandler struct {
	getClientByIDAndProjectIDUseCase contract.GetClientByIDAndProjectIDUseCase
}

// Execute implements contract.GetClientByIDAndProjectIDHandler.
func (h *getClientByIDAndProjectIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")
	if projectID == "" {
		stackerror.HttpResponse(w, "FetchClientsByProjectIDHandler", errors.New("project_id is required"))
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		stackerror.HttpResponse(w, "FetchClientsByProjectIDHandler", errors.New("id is required"))
		return
	}

	client, err := h.getClientByIDAndProjectIDUseCase.Execute(r.Context(), id, projectID)
	if err != nil {
		stackerror.HttpResponse(w, "FetchClientsByProjectIDHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

func NewGetClientByIDAndProjectIDHandler(
	getClientByIDAndProjectIDUseCase contract.GetClientByIDAndProjectIDUseCase,
) contract.GetClientByIDAndProjectIDHandler {
	return &getClientByIDAndProjectIDHandler{
		getClientByIDAndProjectIDUseCase: getClientByIDAndProjectIDUseCase,
	}
}
