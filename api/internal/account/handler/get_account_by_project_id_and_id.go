package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/go-chi/chi/v5"
)

type getAccountByProjectIDAndIDHandler struct {
	getAccountByProjectIDAndIDUseCase contract.GetAccountByProjectIDAndIDUseCase
}

func (h *getAccountByProjectIDAndIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")

	if projectID == "" {
		stackerror.HttpResponse(w, "GetAccountByProjectIDAndIDHandler", errors.New("project_id is required"))
		return
	}

	id := chi.URLParam(r, "id")

	if id == "" {
		stackerror.HttpResponse(w, "GetAccountByProjectIDAndIDHandler", errors.New("id is required"))
		return
	}

	account, err := h.getAccountByProjectIDAndIDUseCase.Execute(r.Context(), projectID, id)
	if err != nil {
		stackerror.HttpResponse(w, "GetAccountByProjectIDAndIDHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

func NewGetAccountByProjectIDAndIDHandler(
	getAccountByProjectIDAndIDUseCase contract.GetAccountByProjectIDAndIDUseCase,
) contract.GetAccountByProjectIDAndIDHandler {
	return &getAccountByProjectIDAndIDHandler{
		getAccountByProjectIDAndIDUseCase: getAccountByProjectIDAndIDUseCase,
	}
}
