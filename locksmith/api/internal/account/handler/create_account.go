package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/types/input"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/go-chi/chi/v5"
)

type createAccountHandler struct {
	createAccountUseCase contract.CreateAccountUseCase
}

// Execute implements contract.CreateAccountHandler.
func (h *createAccountHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")

	if projectID == "" {
		stackerror.HttpResponse(w, "CreateAccountHandler", errors.New("project_id is required"))
		return
	}

	var in input.Account

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		stackerror.HttpResponse(w, "CreateAccountHandler", err)
		return
	}

	in.ProjectID = projectID

	if err := in.ValidateHttp(w); err != nil {
		return
	}

	out, err := h.createAccountUseCase.Execute(r.Context(), in)
	if err != nil {
		stackerror.HttpResponse(w, "CreateAccountHandler", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(out)
}

func NewCreateAccountHandler(createAccountUseCase contract.CreateAccountUseCase) contract.CreateAccountHandler {
	return &createAccountHandler{
		createAccountUseCase: createAccountUseCase,
	}
}
