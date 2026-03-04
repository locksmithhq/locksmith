package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/acl/types/input"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/go-chi/chi/v5"
)

type createProjectAclHandler struct {
	createProjectAclUseCase contract.CreateProjectAclUseCase
}

func (h *createProjectAclHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	if projectId == "" {
		stackerror.HttpResponse(w, "CreateProjectAclHandler", errors.New("project id is required"))
		return
	}

	var in input.ProjectAcl
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		stackerror.HttpResponse(w, "CreateProjectAclHandler", err)
		return
	}

	err := h.createProjectAclUseCase.Execute(r.Context(), projectId, in)
	if err != nil {
		stackerror.HttpResponse(w, "CreateProjectAclHandler", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewCreateProjectAclHandler(createProjectAclUseCase contract.CreateProjectAclUseCase) contract.CreateProjectAclHandler {
	return &createProjectAclHandler{createProjectAclUseCase: createProjectAclUseCase}
}
