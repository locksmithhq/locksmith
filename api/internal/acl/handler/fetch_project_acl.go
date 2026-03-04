package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/go-chi/chi/v5"
)

type fetchProjectAclHandler struct {
	fetchProjectAclUseCase contract.FetchProjectAclUseCase
}

// FetchAll implements contract.FetchAclHandler.
func (h *fetchProjectAclHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectId := chi.URLParam(r, "projectId")
	if projectId == "" {
		stackerror.HttpResponse(w, "FetchAll", errors.New("projectId is required"))
		return
	}

	roles, err := h.fetchProjectAclUseCase.Execute(r.Context(), projectId)
	if err != nil {
		stackerror.HttpResponse(w, "FetchAll", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roles)
}

func NewFetchProjectAclHandler(fetchProjectAclUseCase contract.FetchProjectAclUseCase) contract.FetchProjectAclHandler {
	return &fetchProjectAclHandler{fetchProjectAclUseCase: fetchProjectAclUseCase}
}
