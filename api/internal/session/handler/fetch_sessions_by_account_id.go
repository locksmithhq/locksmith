package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
)

type fetchSessionsByAccountIDHandler struct {
	fetchSessionsByAccountIDUseCase contract.FetchSessionsByAccountIDUseCase
}

func (h *fetchSessionsByAccountIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")
	accountID := chi.URLParam(r, "account_id")
	if projectID == "" || accountID == "" {
		stackerror.HttpResponse(w, "FetchSessionsByAccountIDHandler", errors.New("project_id and account_id are required"))
		return
	}

	query := r.URL.Query()

	page, _ := strconv.Atoi(query.Get("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit < 1 || limit > 100 {
		limit = 20
	}

	sessions, err := h.fetchSessionsByAccountIDUseCase.Execute(r.Context(), projectID, accountID, page, limit)
	if err != nil {
		stackerror.HttpResponse(w, "FetchSessionsByAccountIDHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sessions)
}

func NewFetchSessionsByAccountIDHandler(fetchSessionsByAccountIDUseCase contract.FetchSessionsByAccountIDUseCase) contract.FetchSessionsByAccountIDHandler {
	return &fetchSessionsByAccountIDHandler{fetchSessionsByAccountIDUseCase: fetchSessionsByAccountIDUseCase}
}
