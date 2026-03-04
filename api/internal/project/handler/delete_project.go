package handler

import (
	"errors"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/project/contract"
	"github.com/go-chi/chi/v5"
)

type deleteProjectHandler struct {
	deleteProjectUseCase contract.DeleteProjectUseCase
}

func (h *deleteProjectHandler) Execute(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		stackerror.HttpResponse(w, "DeleteProjectHandler", errors.New("id is required"))
		return
	}

	err := h.deleteProjectUseCase.Execute(r.Context(), id)
	if err != nil {
		stackerror.HttpResponse(w, "DeleteProjectHandler", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NewDeleteProjectHandler(deleteProjectUseCase contract.DeleteProjectUseCase) contract.DeleteProjectHandler {
	return &deleteProjectHandler{deleteProjectUseCase: deleteProjectUseCase}
}
