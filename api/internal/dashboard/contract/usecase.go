package contract

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/dashboard/types/output"
)

type GetDashboardStatsUseCase interface {
	Execute(context.Context) (output.DashboardStats, error)
}
