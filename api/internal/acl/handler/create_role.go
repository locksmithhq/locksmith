package handler

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/types/input"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
)

type createRoleHandler struct {
	createRoleUseCase contract.CreateRoleUseCase
}

// Execute implements contract.CreateRoleHandler.
func (h *createRoleHandler) Execute(w http.ResponseWriter, r *http.Request) {
	var in input.Role
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		stackerror.HttpResponse(w, "CreateRole", err)
		return
	}

	if in.ValidateHttp(w) != nil {
		return
	}

	if err := h.createRoleUseCase.Execute(r.Context(), in); err != nil {
		stackerror.HttpResponse(w, "CreateRole", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewCreateRoleHandler(createRoleUseCase contract.CreateRoleUseCase) contract.CreateRoleHandler {
	return &createRoleHandler{createRoleUseCase: createRoleUseCase}
}
