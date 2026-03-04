package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/project/contract"
	"github.com/go-chi/chi/v5"
)

type getProjectByIDHandler struct {
	getProjectByIDUseCase contract.GetProjectByIDUseCase
}

// Execute implements contract.GetProjectByIDHandler.
func (h *getProjectByIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "id")
	if projectID == "" {
		stackerror.HttpResponse(w, "GetProjectByIDHandler", errors.New("projectID is required"))
		return
	}

	project, err := h.getProjectByIDUseCase.Execute(r.Context(), projectID)
	if err != nil {
		stackerror.HttpResponse(w, "GetProjectByIDHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

func NewGetProjectByIDHandler(getProjectByIDUseCase contract.GetProjectByIDUseCase) contract.GetProjectByIDHandler {
	return &getProjectByIDHandler{getProjectByIDUseCase: getProjectByIDUseCase}
}
