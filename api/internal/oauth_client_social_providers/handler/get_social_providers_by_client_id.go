package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/contract"
)

type getSocialProvidersByClientIDHandler struct {
	getSocialProvidersByClientIDUseCase contract.GetSocialProvidersByClientIDUseCase
}

func (h *getSocialProvidersByClientIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "client_id")

	providers, err := h.getSocialProvidersByClientIDUseCase.Execute(r.Context(), clientID)
	if err != nil {
		stackerror.HttpResponse(w, "GetSocialProvidersByClientIDHandler", err)
		return
	}

	json.NewEncoder(w).Encode(providers)
}

func NewGetSocialProvidersByClientIDHandler(
	getSocialProvidersByClientIDUseCase contract.GetSocialProvidersByClientIDUseCase,
) contract.GetSocialProvidersByClientIDHandler {
	return &getSocialProvidersByClientIDHandler{
		getSocialProvidersByClientIDUseCase: getSocialProvidersByClientIDUseCase,
	}
}
