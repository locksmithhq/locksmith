package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/project/contract"
	"github.com/locksmithhq/locksmith/api/internal/project/types/input"
	"github.com/go-chi/chi/v5"
)

type updateProjectHandler struct {
	updateProjectUseCase contract.UpdateProjectUseCase
}

func (h *updateProjectHandler) Execute(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		stackerror.HttpResponse(w, "UpdateProjectHandler", errors.New("id is required"))
		return
	}

	var in input.Project
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		stackerror.HttpResponse(w, "CreateProjectHandler", err)
		return
	}

	if err := in.ValidateHttp(w); err != nil {
		return
	}

	project, err := h.updateProjectUseCase.Execute(r.Context(), id, in)
	if err != nil {
		stackerror.HttpResponse(w, "UpdateProjectHandler", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(project)
}

func NewUpdateProjectHandler(updateProjectUseCase contract.UpdateProjectUseCase) contract.UpdateProjectHandler {
	return &updateProjectHandler{updateProjectUseCase: updateProjectUseCase}
}
