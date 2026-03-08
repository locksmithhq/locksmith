package contract

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/dashboard/types/output"
)

type GetDashboardStatsRepository interface {
	Execute(context.Context) (output.DashboardStats, error)
}
