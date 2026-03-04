package handler

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
)

type fetchModulesHandler struct {
	fetchModulesUseCase contract.FetchModulesUseCase
}

func (h *fetchModulesHandler) Execute(w http.ResponseWriter, r *http.Request) {
	modules, err := h.fetchModulesUseCase.Execute(r.Context())
	if err != nil {
		stackerror.HttpResponse(w, "FetchModules", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(modules)
}

func NewFetchModulesHandler(fetchModulesUseCase contract.FetchModulesUseCase) contract.FetchModulesHandler {
	return &fetchModulesHandler{fetchModulesUseCase: fetchModulesUseCase}
}
