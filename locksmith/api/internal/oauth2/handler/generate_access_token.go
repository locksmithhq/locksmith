package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/locksmithhq/locksmith-go"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/types/input"
)

type generateAccessTokenHandler struct {
	generateAccessTokenUseCase contract.GenerateAccessTokenUseCase
}

// Execute implements contract.GenerateAccessTokenHandler.
func (h *generateAccessTokenHandler) Execute(w http.ResponseWriter, r *http.Request) {
	var req input.AccessToken
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		stackerror.HttpResponse(w, "GenerateAccessTokenHandler", err)
		return
	}

	if req.GrantType != "authorization_code" {
		stackerror.HttpResponse(w, "GenerateAccessTokenHandler", errors.New("unsupported grant_type"))
		return
	}

	fp := locksmith.Parse(r)
	req.DeviceID = r.Header.Get("X-Device-ID")
	req.IPAddress = fp.IPAddress
	req.UserAgent = fp.UserAgent
	req.DeviceType = fp.DeviceType
	req.Browser = fp.Browser
	req.BrowserVersion = fp.BrowserVersion
	req.OS = fp.OS
	req.OSVersion = fp.OSVersion
	req.LocationCountry = fp.LocationCountry
	req.LocationRegion = fp.LocationRegion
	req.LocationCity = fp.LocationCity

	tokens, err := h.generateAccessTokenUseCase.Execute(r.Context(), req)
	if err != nil {
		stackerror.HttpResponse(w, "GenerateAccessTokenHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}

func NewGenerateAccessTokenHandler(
	generateAccessTokenUseCase contract.GenerateAccessTokenUseCase,
) contract.GenerateAccessTokenHandler {
	return &generateAccessTokenHandler{
		generateAccessTokenUseCase: generateAccessTokenUseCase,
	}
}
