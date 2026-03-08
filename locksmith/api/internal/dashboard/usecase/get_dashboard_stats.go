package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/dashboard/contract"
	"github.com/locksmithhq/locksmith/api/internal/dashboard/types/output"
)

type getDashboardStatsUseCase struct {
	getDashboardStatsRepository contract.GetDashboardStatsRepository
}

func (u *getDashboardStatsUseCase) Execute(ctx context.Context) (output.DashboardStats, error) {
	stats, err := u.getDashboardStatsRepository.Execute(ctx)
	if err != nil {
		return output.DashboardStats{}, stackerror.NewUseCaseError("GetDashboardStatsUseCase", err)
	}

	return stats, nil
}

func NewGetDashboardStatsUseCase(
	getDashboardStatsRepository contract.GetDashboardStatsRepository,
) contract.GetDashboardStatsUseCase {
	return &getDashboardStatsUseCase{
		getDashboardStatsRepository: getDashboardStatsRepository,
	}
}
