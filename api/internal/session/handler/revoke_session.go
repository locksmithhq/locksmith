package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
)

type revokeSessionHandler struct {
	revokeSessionUseCase contract.RevokeSessionUseCase
}

func (h *revokeSessionHandler) Execute(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "project_id")
	sessionID := chi.URLParam(r, "session_id")
	if projectID == "" || sessionID == "" {
		stackerror.HttpResponse(w, "RevokeSessionHandler", errors.New("project_id and session_id are required"))
		return
	}

	if err := h.revokeSessionUseCase.Execute(r.Context(), projectID, sessionID); err != nil {
		stackerror.HttpResponse(w, "RevokeSessionHandler", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NewRevokeSessionHandler(revokeSessionUseCase contract.RevokeSessionUseCase) contract.RevokeSessionHandler {
	return &revokeSessionHandler{revokeSessionUseCase: revokeSessionUseCase}
}
