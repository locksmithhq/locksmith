package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/booscaaa/locksmith/api/internal/account/contract"
	"github.com/booscaaa/locksmith/api/internal/account/types/input"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/go-chi/chi/v5"
)

type updateAccountHandler struct {
	updateAccountUseCase contract.UpdateAccountUseCase
}

// Execute implements contract.UpdateAccountHandler.
func (h *updateAccountHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")
	accountID := chi.URLParam(r, "account_id")

	if projectID == "" {
		stackerror.HttpResponse(w, "UpdateAccountHandler", errors.New("project_id is required"))
		return
	}

	if accountID == "" {
		stackerror.HttpResponse(w, "UpdateAccountHandler", errors.New("account_id is required"))
		return
	}

	var in input.UpdateAccount

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		stackerror.HttpResponse(w, "UpdateAccountHandler", err)
		return
	}

	in.ProjectID = projectID
	in.ID = accountID

	if err := in.ValidateHttp(w); err != nil {
		return
	}

	out, err := h.updateAccountUseCase.Execute(r.Context(), in)
	if err != nil {
		stackerror.HttpResponse(w, "UpdateAccountHandler", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(out)
}

func NewUpdateAccountHandler(updateAccountUseCase contract.UpdateAccountUseCase) contract.UpdateAccountHandler {
	return &updateAccountHandler{
		updateAccountUseCase: updateAccountUseCase,
	}
}
