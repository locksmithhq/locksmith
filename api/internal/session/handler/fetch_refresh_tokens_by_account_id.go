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

type fetchRefreshTokensByAccountIDHandler struct {
	fetchRefreshTokensByAccountIDUseCase contract.FetchRefreshTokensByAccountIDUseCase
}

func (h *fetchRefreshTokensByAccountIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")
	accountID := chi.URLParam(r, "account_id")
	if projectID == "" || accountID == "" {
		stackerror.HttpResponse(w, "FetchRefreshTokensByAccountIDHandler", errors.New("project_id and account_id are required"))
		return
	}

	query := r.URL.Query()
	sessionID := query.Get("session_id")

	page, _ := strconv.Atoi(query.Get("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit < 1 || limit > 100 {
		limit = 50
	}

	tokens, err := h.fetchRefreshTokensByAccountIDUseCase.Execute(r.Context(), projectID, accountID, sessionID, page, limit)
	if err != nil {
		stackerror.HttpResponse(w, "FetchRefreshTokensByAccountIDHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)
}

func NewFetchRefreshTokensByAccountIDHandler(fetchRefreshTokensByAccountIDUseCase contract.FetchRefreshTokensByAccountIDUseCase) contract.FetchRefreshTokensByAccountIDHandler {
	return &fetchRefreshTokensByAccountIDHandler{fetchRefreshTokensByAccountIDUseCase: fetchRefreshTokensByAccountIDUseCase}
}
