package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
)

type countRefreshTokensByAccountIDHandler struct {
	countRefreshTokensByAccountIDUseCase contract.CountRefreshTokensByAccountIDUseCase
}

func (h *countRefreshTokensByAccountIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")
	accountID := chi.URLParam(r, "account_id")
	if projectID == "" || accountID == "" {
		stackerror.HttpResponse(w, "CountRefreshTokensByAccountIDHandler", errors.New("project_id and account_id are required"))
		return
	}

	sessionID := r.URL.Query().Get("session_id")

	count, err := h.countRefreshTokensByAccountIDUseCase.Execute(r.Context(), projectID, accountID, sessionID)
	if err != nil {
		stackerror.HttpResponse(w, "CountRefreshTokensByAccountIDHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(count)
}

func NewCountRefreshTokensByAccountIDHandler(countRefreshTokensByAccountIDUseCase contract.CountRefreshTokensByAccountIDUseCase) contract.CountRefreshTokensByAccountIDHandler {
	return &countRefreshTokensByAccountIDHandler{countRefreshTokensByAccountIDUseCase: countRefreshTokensByAccountIDUseCase}
}
