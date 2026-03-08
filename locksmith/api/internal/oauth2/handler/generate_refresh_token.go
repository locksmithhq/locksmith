package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
)

type generateRefreshTokenHandler struct {
	generateRefreshTokenUseCase contract.GenerateRefreshTokenUseCase
}

type generateRefreshTokenRequest struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

// Execute implements contract.GenerateRefreshTokenHandler.
func (h *generateRefreshTokenHandler) Execute(w http.ResponseWriter, r *http.Request) {
	var req generateRefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		stackerror.HttpResponse(w, "GenerateRefreshTokenHandler", err)
		return
	}

	// Validate grant_type
	if req.GrantType != "refresh_token" {
		stackerror.HttpResponse(w, "GenerateRefreshTokenHandler", errors.New("unsupported grant_type"))
		return
	}

	tokens, err := h.generateRefreshTokenUseCase.Execute(r.Context(), req.RefreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}

func NewGenerateRefreshTokenHandler(
	generateRefreshTokenUseCase contract.GenerateRefreshTokenUseCase,
) contract.GenerateRefreshTokenHandler {
	return &generateRefreshTokenHandler{
		generateRefreshTokenUseCase: generateRefreshTokenUseCase,
	}
}
