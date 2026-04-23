package rest

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/di"
)

func InitializeOAuth2Router(router chi.Router) {
	router.Group(func(router chi.Router) {
		// 5 attempts/min per IP — brute force protection for login
		router.With(httprate.LimitByIP(5, time.Minute)).Post("/oauth2/login", di.NewLoginHandler().Execute)

		// 10 req/min per IP
		router.With(httprate.LimitByIP(10, time.Minute)).Post("/oauth2/register", di.NewRegisterHandler().Execute)
		router.With(httprate.LimitByIP(10, time.Minute)).Post("/oauth2/access-token", di.NewGenerateAccessTokenHandler().Execute)
		router.With(httprate.LimitByIP(10, time.Minute)).Post("/oauth2/refresh-token", di.NewGenerateRefreshTokenHandler().Execute)
		router.With(httprate.LimitByIP(10, time.Minute)).Get("/oauth2/social/{provider}/begin", di.NewSocialBeginHandler().Execute)

		// 20 req/min per IP
		router.With(httprate.LimitByIP(20, time.Minute)).Post("/oauth2/authorize", di.NewAuthorizeClientHandler().Execute)
		router.With(httprate.LimitByIP(20, time.Minute)).Get("/oauth2/social/{provider}/callback", di.NewSocialCallbackHandler().Execute)

		// Unthrottled — read-only, low risk
		router.Get("/oauth2/manifest", di.NewPWAManifestHandler().Execute)
		router.Get("/oauth2/favicon", di.NewFaviconHandler().Execute)
	})
}
