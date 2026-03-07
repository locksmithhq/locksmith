package handler

import (
	"encoding/json"
	"net/http"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/dashboard/contract"
)

type getDashboardStatsHandler struct {
	getDashboardStatsUseCase contract.GetDashboardStatsUseCase
}

func (h *getDashboardStatsHandler) Execute(w http.ResponseWriter, r *http.Request) {
	stats, err := h.getDashboardStatsUseCase.Execute(r.Context())
	if err != nil {
		stackerror.HttpResponse(w, "GetDashboardStatsHandler", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func NewGetDashboardStatsHandler(
	getDashboardStatsUseCase contract.GetDashboardStatsUseCase,
) contract.GetDashboardStatsHandler {
	return &getDashboardStatsHandler{
		getDashboardStatsUseCase: getDashboardStatsUseCase,
	}
}
