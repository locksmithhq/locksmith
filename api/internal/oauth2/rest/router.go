package rest

import (
	"github.com/locksmithhq/locksmith/api/internal/oauth2/di"
	"github.com/go-chi/chi/v5"
)

func InitializeOAuth2Router(router chi.Router) {
	router.Group(func(router chi.Router) {
		router.Post("/oauth2/authorize", di.NewAuthorizeClientHandler().Execute)
		router.Post("/oauth2/login", di.NewLoginHandler().Execute)
		router.Post("/oauth2/register", di.NewRegisterHandler().Execute)
		router.Post("/oauth2/access-token", di.NewGenerateAccessTokenHandler().Execute)
		router.Post("/oauth2/refresh-token", di.NewGenerateRefreshTokenHandler().Execute)
	})
}
