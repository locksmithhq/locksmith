package handler

import (
	"encoding/json"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
)

type pwaManifestHandler struct {
	getPWAManifestUseCase contract.GetPWAManifestUseCase
}

func (h *pwaManifestHandler) Execute(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	if clientID == "" {
		http.Error(w, "client_id is required", http.StatusBadRequest)
		return
	}

	scheme := r.Header.Get("X-Forwarded-Proto")
	if scheme == "" {
		if r.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}
	}

	host := r.Header.Get("X-Forwarded-Host")
	if host == "" {
		host = r.Host
	}

	baseURL := scheme + "://" + host

	manifest, err := h.getPWAManifestUseCase.Execute(r.Context(), clientID, baseURL)
	if err != nil {
		stackerror.HttpResponse(w, "HANDLER: PWAManifestHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/manifest+json")
	json.NewEncoder(w).Encode(manifest)
}

func NewPWAManifestHandler(
	getPWAManifestUseCase contract.GetPWAManifestUseCase,
) contract.PWAManifestHandler {
	return &pwaManifestHandler{
		getPWAManifestUseCase: getPWAManifestUseCase,
	}
}
