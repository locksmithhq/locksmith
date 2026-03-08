package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/dashboard/contract"
	"github.com/locksmithhq/locksmith/api/internal/dashboard/types/output"
)

type getDashboardStatsRepository struct {
	database types.Database
}

func (r *getDashboardStatsRepository) Execute(ctx context.Context) (output.DashboardStats, error) {
	var stats output.DashboardStats

	err := r.database.QueryRowContext(ctx, `SELECT COUNT(*) FROM projects WHERE deleted_at IS NULL`).Scan(&stats.TotalProjects)
	if err != nil {
		return output.DashboardStats{}, stackerror.NewRepositoryError("GetDashboardStatsRepository.TotalProjects", err)
	}

	err = r.database.QueryRowContext(ctx, `SELECT COUNT(*) FROM accounts WHERE deleted_at IS NULL`).Scan(&stats.TotalUsers)
	if err != nil {
		return output.DashboardStats{}, stackerror.NewRepositoryError("GetDashboardStatsRepository.TotalUsers", err)
	}

	err = r.database.QueryRowContext(ctx, `SELECT COUNT(*) FROM user_sessions WHERE revoked = false AND expires_at > NOW()`).Scan(&stats.ActiveSessions)
	if err != nil {
		return output.DashboardStats{}, stackerror.NewRepositoryError("GetDashboardStatsRepository.ActiveSessions", err)
	}

	err = r.database.QueryRowContext(ctx, `SELECT COUNT(*) FROM oauth_clients`).Scan(&stats.TotalClients)
	if err != nil {
		return output.DashboardStats{}, stackerror.NewRepositoryError("GetDashboardStatsRepository.TotalClients", err)
	}

	loginsRows, err := r.database.QueryContext(ctx, `
		SELECT DATE(created_at)::text AS day, COUNT(*) AS count
		FROM user_sessions
		WHERE created_at >= NOW() - INTERVAL '30 days'
		GROUP BY DATE(created_at)
		ORDER BY day
	`)
	if err != nil {
		return output.DashboardStats{}, stackerror.NewRepositoryError("GetDashboardStatsRepository.LoginsPerDay", err)
	}
	defer loginsRows.Close()

	stats.LoginsPerDay = []output.DayCount{}
	for loginsRows.Next() {
		var dc output.DayCount
		if err := loginsRows.Scan(&dc.Day, &dc.Count); err != nil {
			return output.DashboardStats{}, stackerror.NewRepositoryError("GetDashboardStatsRepository.LoginsPerDay.Scan", err)
		}
		stats.LoginsPerDay = append(stats.LoginsPerDay, dc)
	}

	deviceRows, err := r.database.QueryContext(ctx, `
		SELECT COALESCE(device_type, 'unknown') AS device, COUNT(*) AS count
		FROM user_sessions
		WHERE device_type IS NOT NULL
		GROUP BY device_type
	`)
	if err != nil {
		return output.DashboardStats{}, stackerror.NewRepositoryError("GetDashboardStatsRepository.SessionsByDevice", err)
	}
	defer deviceRows.Close()

	stats.SessionsByDevice = []output.DeviceCount{}
	for deviceRows.Next() {
		var dc output.DeviceCount
		if err := deviceRows.Scan(&dc.Device, &dc.Count); err != nil {
			return output.DashboardStats{}, stackerror.NewRepositoryError("GetDashboardStatsRepository.SessionsByDevice.Scan", err)
		}
		stats.SessionsByDevice = append(stats.SessionsByDevice, dc)
	}

	return stats, nil
}

func NewGetDashboardStatsRepository(database types.Database) contract.GetDashboardStatsRepository {
	return &getDashboardStatsRepository{
		database: database,
	}
}
