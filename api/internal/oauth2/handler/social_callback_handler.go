package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/input"
)

type socialCallbackHandler struct {
	socialCallbackUseCase contract.SocialCallbackUseCase
}

func (h *socialCallbackHandler) Execute(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	nonce := r.URL.Query().Get("state")

	in := input.SocialCallback{
		Provider: provider,
		Nonce:    nonce,
	}

	if err := h.socialCallbackUseCase.Execute(r.Context(), in, w, r); err != nil {
		stackerror.HttpResponse(w, "SocialCallbackHandler", err)
		return
	}
}

func NewSocialCallbackHandler(socialCallbackUseCase contract.SocialCallbackUseCase) contract.SocialCallbackHandler {
	return &socialCallbackHandler{
		socialCallbackUseCase: socialCallbackUseCase,
	}
}
