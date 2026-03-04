package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/booscaaa/locksmith/api/internal/account/contract"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/go-chi/chi/v5"
)

type countAccountsByProjectIDHandler struct {
	countAccountsByProjectIDUseCase contract.CountAccountsByProjectIDUseCase
}

// Execute implements contract.CountAccountsByProjectIDHandler.
func (h *countAccountsByProjectIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")

	if projectID == "" {
		stackerror.HttpResponse(w, "CountAccountsByProjectIDHandler", errors.New("project_id is required"))
		return
	}

	params, err := paginate.BindQueryParamsToStruct(r.URL.Query())
	if err != nil {
		stackerror.HttpResponse(w, "CountAccountsByProjectIDHandler", err)
		return
	}

	total, err := h.countAccountsByProjectIDUseCase.Execute(r.Context(), projectID, *params)
	if err != nil {
		stackerror.HttpResponse(w, "CountAccountsByProjectIDHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int64{"total": total})
}

func NewCountAccountsByProjectIDHandler(countAccountsByProjectIDUseCase contract.CountAccountsByProjectIDUseCase) contract.CountAccountsByProjectIDHandler {
	return &countAccountsByProjectIDHandler{countAccountsByProjectIDUseCase: countAccountsByProjectIDUseCase}
}
