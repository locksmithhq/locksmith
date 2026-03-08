package rest

import (
	"github.com/locksmithhq/locksmith/api/internal/dashboard/di"
	"github.com/go-chi/chi/v5"
	locksmith "github.com/locksmithhq/locksmith-go"
)

func InitializeDashboardRouter(router chi.Router) {
	router.Group(func(r chi.Router) {
		r.Use(locksmith.AuthMiddlewareCookie("LOCKSMITHACCESSTOKEN"))

		r.Get("/dashboard/stats", di.NewGetDashboardStatsHandler().Execute)
	})
}
