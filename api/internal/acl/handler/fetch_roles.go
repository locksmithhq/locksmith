package handler

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
)

type fetchRolesHandler struct {
	fetchRolesUseCase contract.FetchRolesUseCase
}

func (h *fetchRolesHandler) Execute(w http.ResponseWriter, r *http.Request) {
	roles, err := h.fetchRolesUseCase.Execute(r.Context())
	if err != nil {
		stackerror.HttpResponse(w, "FetchRoles", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roles)
}

func NewFetchRolesHandler(fetchRolesUseCase contract.FetchRolesUseCase) contract.FetchRolesHandler {
	return &fetchRolesHandler{fetchRolesUseCase: fetchRolesUseCase}
}
