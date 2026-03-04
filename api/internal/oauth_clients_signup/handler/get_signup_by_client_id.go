package handler

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/locksmith/api/internal/oauth_clients_signup/contract"
	"github.com/go-chi/chi/v5"
)

type getSignupByClientIDHandler struct {
	getSignupByClientIDUseCase contract.GetSignupByClientIDUseCase
}

func (h *getSignupByClientIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "id")

	signup, err := h.getSignupByClientIDUseCase.Execute(r.Context(), clientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(signup)
}

func NewGetSignupByClientIDHandler(
	getSignupByClientIDUseCase contract.GetSignupByClientIDUseCase,
) contract.GetSignupByClientIDHandler {
	return &getSignupByClientIDHandler{
		getSignupByClientIDUseCase: getSignupByClientIDUseCase,
	}
}
