package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/go-chi/chi/v5"
)

type fetchAccountsByProjectIDHandler struct {
	fetchAccountsByProjectIDUseCase contract.FetchAccountsByProjectIDUseCase
}

// Execute implements contract.FetchAccountsByProjectIDHandler.
func (h *fetchAccountsByProjectIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")

	if projectID == "" {
		stackerror.HttpResponse(w, "FetchAccountsByProjectIDHandler", errors.New("project_id is required"))
		return
	}

	params, err := paginate.BindQueryParamsToStruct(r.URL.Query())
	if err != nil {
		stackerror.HttpResponse(w, "FetchAccountsByProjectIDHandler", err)
		return
	}

	accounts, err := h.fetchAccountsByProjectIDUseCase.Execute(r.Context(), projectID, *params)
	if err != nil {
		stackerror.HttpResponse(w, "FetchAccountsByProjectIDHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}

func NewFetchAccountsByProjectIDHandler(fetchAccountsByProjectIDUseCase contract.FetchAccountsByProjectIDUseCase) contract.FetchAccountsByProjectIDHandler {
	return &fetchAccountsByProjectIDHandler{fetchAccountsByProjectIDUseCase: fetchAccountsByProjectIDUseCase}
}
