package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
)

type resolveCustomDomainHandler struct {
	resolveCustomDomainUseCase contract.ResolveCustomDomainUseCase
}

// Execute implements contract.ResolveCustomDomainHandler.
func (h *resolveCustomDomainHandler) Execute(w http.ResponseWriter, r *http.Request) {
	hostname := r.URL.Query().Get("hostname")
	if hostname == "" {
		stackerror.HttpResponse(w, "ResolveCustomDomainHandler", errors.New("hostname is required"))
		return
	}

	clientID, err := h.resolveCustomDomainUseCase.Execute(r.Context(), hostname)
	if err != nil {
		stackerror.HttpResponse(w, "ResolveCustomDomainHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"client_id": clientID})
}

func NewResolveCustomDomainHandler(
	resolveCustomDomainUseCase contract.ResolveCustomDomainUseCase,
) contract.ResolveCustomDomainHandler {
	return &resolveCustomDomainHandler{
		resolveCustomDomainUseCase: resolveCustomDomainUseCase,
	}
}
