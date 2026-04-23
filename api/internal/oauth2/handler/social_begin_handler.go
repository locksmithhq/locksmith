package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/input"
)

type socialBeginHandler struct {
	socialBeginUseCase contract.SocialBeginUseCase
}

func (h *socialBeginHandler) Execute(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	in := input.SocialBegin{
		Provider:            provider,
		ClientID:            r.URL.Query().Get("client_id"),
		RedirectURI:         r.URL.Query().Get("redirect_uri"),
		State:               r.URL.Query().Get("state"),
		CodeChallenge:       r.URL.Query().Get("code_challenge"),
		CodeChallengeMethod: r.URL.Query().Get("code_challenge_method"),
	}

	result, err := h.socialBeginUseCase.Execute(r.Context(), in)
	if err != nil {
		stackerror.HttpResponse(w, "SocialBeginHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func NewSocialBeginHandler(socialBeginUseCase contract.SocialBeginUseCase) contract.SocialBeginHandler {
	return &socialBeginHandler{
		socialBeginUseCase: socialBeginUseCase,
	}
}
