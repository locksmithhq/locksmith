package handler

import (
	"encoding/json"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/project/contract"
	"github.com/locksmithhq/locksmith/api/internal/project/types/input"
)

type createProjectHandler struct {
	createProjectUseCase contract.CreateProjectUseCase
}

func (h *createProjectHandler) Execute(w http.ResponseWriter, r *http.Request) {
	var in input.Project
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		stackerror.HttpResponse(w, "CreateProjectHandler", err)
		return
	}

	if err := in.ValidateHttp(w); err != nil {
		return
	}

	project, err := h.createProjectUseCase.Execute(r.Context(), in)
	if err != nil {
		stackerror.HttpResponse(w, "CreateProjectHandler", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

func NewCreateProjectHandler(createProjectUseCase contract.CreateProjectUseCase) contract.CreateProjectHandler {
	return &createProjectHandler{createProjectUseCase: createProjectUseCase}
}
