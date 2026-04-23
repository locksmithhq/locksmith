package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/types/input"
)

type upsertSocialProviderHandler struct {
	upsertSocialProviderUseCase contract.UpsertSocialProviderUseCase
}

func (h *upsertSocialProviderHandler) Execute(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "client_id")
	provider := chi.URLParam(r, "provider")

	var in input.SocialProvider
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		stackerror.HttpResponse(w, "UpsertSocialProviderHandler", err)
		return
	}

	result, err := h.upsertSocialProviderUseCase.Execute(r.Context(), clientID, provider, in)
	if err != nil {
		stackerror.HttpResponse(w, "UpsertSocialProviderHandler", err)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func NewUpsertSocialProviderHandler(
	upsertSocialProviderUseCase contract.UpsertSocialProviderUseCase,
) contract.UpsertSocialProviderHandler {
	return &upsertSocialProviderHandler{
		upsertSocialProviderUseCase: upsertSocialProviderUseCase,
	}
}
