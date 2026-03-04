package handler

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/contract"
	"github.com/go-chi/chi/v5"
)

type getLoginByClientIDHandler struct {
	getLoginByClientIDUseCase contract.GetLoginByClientIDUseCase
}

func (h *getLoginByClientIDHandler) Execute(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "id")

	login, err := h.getLoginByClientIDUseCase.Execute(r.Context(), clientID)
	if err != nil {
		stackerror.HttpResponse(w, "GetLoginByClientIDHandler", err)
		return
	}

	json.NewEncoder(w).Encode(login)
}

func NewGetLoginByClientIDHandler(
	getLoginByClientIDUseCase contract.GetLoginByClientIDUseCase,
) contract.GetLoginByClientIDHandler {
	return &getLoginByClientIDHandler{
		getLoginByClientIDUseCase: getLoginByClientIDUseCase,
	}
}
