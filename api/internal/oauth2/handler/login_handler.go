package handler

import (
	"encoding/json"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/input"
)

type loginHandler struct {
	loginUseCase contract.LoginUseCase
}

// Execute implements contract.LoginHandler.
func (h *loginHandler) Execute(w http.ResponseWriter, r *http.Request) {
	var req input.Login
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	login, err := h.loginUseCase.Execute(r.Context(), req)
	if err != nil {
		stackerror.HttpResponse(w, "LoginHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(login)
}

func NewLoginHandler(
	loginUseCase contract.LoginUseCase,
) contract.LoginHandler {
	return &loginHandler{
		loginUseCase: loginUseCase,
	}
}
