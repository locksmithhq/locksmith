package handler

import (
	"encoding/json"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/input"
)

type authorizeHandler struct {
	authorizeClientUseCase contract.AuthorizeClient
}

// Execute implements contract.AuthorizeHandler.
func (h *authorizeHandler) Execute(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	input := input.Authorization{
		ClientID:            queryParams.Get("client_id"),
		RedirectURI:         queryParams.Get("redirect_uri"),
		ResponseType:        queryParams.Get("response_type"),
		State:               queryParams.Get("state"),
		CodeChallenge:       queryParams.Get("code_challenge"),
		CodeChallengeMethod: queryParams.Get("code_challenge_method"),
	}

	client, err := h.authorizeClientUseCase.Execute(r.Context(), input)
	if err != nil {
		stackerror.HttpResponse(w, "HANDLER: AuthorizeClientHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

func NewAuthorizeClientHandler(
	authorizeClientUseCase contract.AuthorizeClient,
) contract.AuthorizeClientHandler {
	return &authorizeHandler{
		authorizeClientUseCase: authorizeClientUseCase,
	}
}
