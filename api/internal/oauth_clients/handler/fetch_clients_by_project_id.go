package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/contract"
	"github.com/go-chi/chi/v5"
)

type fetchClientsByProjectIDHandler struct {
	fetchClientsByProjectIDUseCase contract.FetchClientsByProjectIDUseCase
}

// Execute implements contract.FetchClientsByProjectIDHandler.
func (h *fetchClientsByProjectIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")
	if projectID == "" {
		stackerror.HttpResponse(w, "FetchClientsByProjectIDHandler", errors.New("project_id is required"))
		return
	}

	params, err := paginate.BindQueryParamsToStruct(r.URL.Query())
	if err != nil {
		stackerror.HttpResponse(w, "FetchClientsByProjectIDHandler", err)
		return
	}

	clients, err := h.fetchClientsByProjectIDUseCase.Execute(r.Context(), projectID, *params)
	if err != nil {
		stackerror.HttpResponse(w, "FetchClientsByProjectIDHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}

func NewFetchClientsByProjectIDHandler(fetchClientsByProjectIDUseCase contract.FetchClientsByProjectIDUseCase) contract.FetchClientsByProjectIDHandler {
	return &fetchClientsByProjectIDHandler{
		fetchClientsByProjectIDUseCase: fetchClientsByProjectIDUseCase,
	}
}
