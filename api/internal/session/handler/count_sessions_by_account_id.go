package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
)

type countSessionsByAccountIDHandler struct {
	countSessionsByAccountIDUseCase contract.CountSessionsByAccountIDUseCase
}

func (h *countSessionsByAccountIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")
	accountID := chi.URLParam(r, "account_id")
	if projectID == "" || accountID == "" {
		stackerror.HttpResponse(w, "CountSessionsByAccountIDHandler", errors.New("project_id and account_id are required"))
		return
	}

	total, err := h.countSessionsByAccountIDUseCase.Execute(r.Context(), projectID, accountID)
	if err != nil {
		stackerror.HttpResponse(w, "CountSessionsByAccountIDHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(total)
}

func NewCountSessionsByAccountIDHandler(countSessionsByAccountIDUseCase contract.CountSessionsByAccountIDUseCase) contract.CountSessionsByAccountIDHandler {
	return &countSessionsByAccountIDHandler{countSessionsByAccountIDUseCase: countSessionsByAccountIDUseCase}
}
