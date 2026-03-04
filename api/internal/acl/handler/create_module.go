package handler

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/types/input"
)

type createModuleHandler struct {
	createModuleUseCase contract.CreateModuleUseCase
}

// Execute implements contract.CreateRoleHandler.
func (h *createModuleHandler) Execute(w http.ResponseWriter, r *http.Request) {
	var in input.Module
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if in.ValidateHttp(w) != nil {
		return
	}

	if err := h.createModuleUseCase.Execute(r.Context(), in); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewCreateModuleHandler(createModuleUseCase contract.CreateModuleUseCase) contract.CreateModuleHandler {
	return &createModuleHandler{createModuleUseCase: createModuleUseCase}
}
