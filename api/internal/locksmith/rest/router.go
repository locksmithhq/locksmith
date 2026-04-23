package rest

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	"github.com/locksmithhq/locksmith/api/internal/locksmith/di"
)

func InitializeLocksmithRouter(router chi.Router) {
	router.Get("/locksmith/callback", di.NewCallbackHandler().Execute)
	router.Get("/locksmith/status", di.NewStatusHandler().Execute)

	router.With(httprate.LimitByIP(20, time.Minute)).Post("/locksmith/logout", di.NewLogoutHandler().Execute)
	router.With(httprate.LimitByIP(20, time.Minute)).Post("/locksmith/r", di.NewRefreshTokenHandler().Execute)
}
