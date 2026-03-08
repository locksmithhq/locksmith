package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/dashboard/contract"
	"github.com/locksmithhq/locksmith/api/internal/dashboard/handler"
	"github.com/locksmithhq/locksmith/api/internal/dashboard/repository"
	"github.com/locksmithhq/locksmith/api/internal/dashboard/usecase"
)

func NewGetDashboardStatsHandler() contract.GetDashboardStatsHandler {
	conn := database.GetConnection()
	return handler.NewGetDashboardStatsHandler(
		usecase.NewGetDashboardStatsUseCase(
			repository.NewGetDashboardStatsRepository(conn),
		),
	)
}
