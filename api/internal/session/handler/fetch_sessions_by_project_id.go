package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
	"github.com/go-chi/chi/v5"
)

type fetchSessionsByProjectIDHandler struct {
	fetchSessionsByProjectIDUseCase contract.FetchSessionsByProjectIDUseCase
}

func (h *fetchSessionsByProjectIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")
	if projectID == "" {
		stackerror.HttpResponse(w, "FetchSessionsByProjectIDHandler", errors.New("project_id is required"))
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

	search := query.Get("search")

	sessions, err := h.fetchSessionsByProjectIDUseCase.Execute(r.Context(), projectID, page, limit, search)
	if err != nil {
		stackerror.HttpResponse(w, "FetchSessionsByProjectIDHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sessions)
}

func NewFetchSessionsByProjectIDHandler(fetchSessionsByProjectIDUseCase contract.FetchSessionsByProjectIDUseCase) contract.FetchSessionsByProjectIDHandler {
	return &fetchSessionsByProjectIDHandler{fetchSessionsByProjectIDUseCase: fetchSessionsByProjectIDUseCase}
}
