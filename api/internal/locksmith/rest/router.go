package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/locksmithhq/locksmith/api/internal/locksmith/di"
)

func InitializeLocksmithRouter(router chi.Router) {
	router.Get("/locksmith/callback", di.NewCallbackHandler().Execute)
	router.Post("/locksmith/logout", di.NewLogoutHandler().Execute)
	router.Get("/locksmith/status", di.NewStatusHandler().Execute)
	router.Post("/locksmith/r", di.NewRefreshTokenHandler().Execute)
}
