package handler

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
)

type fetchActionsHandler struct {
	fetchActionsUseCase contract.FetchActionsUseCase
}

func (h *fetchActionsHandler) Execute(w http.ResponseWriter, r *http.Request) {
	actions, err := h.fetchActionsUseCase.Execute(r.Context())
	if err != nil {
		stackerror.HttpResponse(w, "FetchActions", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(actions)
}

func NewFetchActionsHandler(fetchActionsUseCase contract.FetchActionsUseCase) contract.FetchActionsHandler {
	return &fetchActionsHandler{fetchActionsUseCase: fetchActionsUseCase}
}
