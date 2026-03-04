package handler

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_signup/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_signup/types/input"
	"github.com/go-chi/chi/v5"
)

type createSignupByClientIDHandler struct {
	createSignupByClientIDUseCase contract.CreateSignupByClientIDUseCase
}

// Execute implements contract.CreateSignupByClientIDHandler.
func (h *createSignupByClientIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "id")

	var in input.Signup
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		stackerror.HttpResponse(w, "CreateSignupByClientIDHandler", err)
		return
	}

	if err := h.createSignupByClientIDUseCase.Execute(r.Context(), clientID, in); err != nil {
		stackerror.HttpResponse(w, "CreateSignupByClientIDHandler", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewCreateSignupByClientIDHandler(createSignupByClientIDUseCase contract.CreateSignupByClientIDUseCase) contract.CreateSignupByClientIDHandler {
	return &createSignupByClientIDHandler{
		createSignupByClientIDUseCase: createSignupByClientIDUseCase,
	}
}
