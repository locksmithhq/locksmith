package handler

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/project/contract"
)

type fetchProjectsHandler struct {
	fetchProjectsUseCase contract.FetchProjectsUseCase
}

// Execute implements contract.FetchProjectsHandler.
func (h *fetchProjectsHandler) Execute(w http.ResponseWriter, r *http.Request) {
	params, err := paginate.BindQueryParamsToStruct(r.URL.Query())
	if err != nil {
		stackerror.HttpResponse(w, "FetchProjectsHandler", err)
		return
	}

	projects, err := h.fetchProjectsUseCase.Execute(r.Context(), *params)
	if err != nil {
		stackerror.HttpResponse(w, "FetchProjectsHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func NewFetchProjectsHandler(fetchProjectsUseCase contract.FetchProjectsUseCase) contract.FetchProjectsHandler {
	return &fetchProjectsHandler{fetchProjectsUseCase: fetchProjectsUseCase}
}
