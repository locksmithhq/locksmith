package handler

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/locksmithhq/locksmith/api/internal/acl/contract"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
)

type fetchAclHandler struct {
	fetchAclUseCase contract.FetchAclUseCase
}

// FetchAll implements contract.FetchAclHandler.
func (h *fetchAclHandler) Execute(w http.ResponseWriter, r *http.Request) {
	filter, err := paginate.BindQueryParamsToStruct(r.URL.Query())
	if err != nil {
		stackerror.HttpResponse(w, "FetchAll", err)
		return
	}

	roles, err := h.fetchAclUseCase.Execute(r.Context(), *filter)
	if err != nil {
		stackerror.HttpResponse(w, "FetchAll", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roles)
}

func NewFetchAclHandler(fetchAclUseCase contract.FetchAclUseCase) contract.FetchAclHandler {
	return &fetchAclHandler{fetchAclUseCase: fetchAclUseCase}
}
