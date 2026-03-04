package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
	"github.com/go-chi/chi/v5"
)

type countSessionsByProjectIDHandler struct {
	countSessionsByProjectIDUseCase contract.CountSessionsByProjectIDUseCase
}

func (h *countSessionsByProjectIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")
	if projectID == "" {
		stackerror.HttpResponse(w, "CountSessionsByProjectIDHandler", errors.New("project_id is required"))
		return
	}

	search := r.URL.Query().Get("search")

	total, err := h.countSessionsByProjectIDUseCase.Execute(r.Context(), projectID, search)
	if err != nil {
		stackerror.HttpResponse(w, "CountSessionsByProjectIDHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int64{"total": total})
}

func NewCountSessionsByProjectIDHandler(countSessionsByProjectIDUseCase contract.CountSessionsByProjectIDUseCase) contract.CountSessionsByProjectIDHandler {
	return &countSessionsByProjectIDHandler{countSessionsByProjectIDUseCase: countSessionsByProjectIDUseCase}
}
