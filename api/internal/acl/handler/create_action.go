package handler

import (
	"encoding/json"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/acl/types/input"
)

type createActionHandler struct {
	createActionUseCase contract.CreateActionUseCase
}

// Execute implements contract.CreateRoleHandler.
func (h *createActionHandler) Execute(w http.ResponseWriter, r *http.Request) {
	var in input.Action
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if in.ValidateHttp(w) != nil {
		return
	}

	if err := h.createActionUseCase.Execute(r.Context(), in); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewCreateActionHandler(createActionUseCase contract.CreateActionUseCase) contract.CreateActionHandler {
	return &createActionHandler{createActionUseCase: createActionUseCase}
}
