package rest

import (
	"github.com/go-chi/chi/v5"
	locksmith "github.com/locksmithhq/locksmith-go"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/di"
)

func InitializeOauthClientSocialProvidersRouter(router chi.Router) {
	router.Group(func(r chi.Router) {
		r.Use(locksmith.AuthMiddlewareCookie("LOCKSMITHACCESSTOKEN"))
		r.Get("/projects/{project_id}/clients/{client_id}/social-providers", di.NewGetSocialProvidersHandler().Execute)
		r.Put("/projects/{project_id}/clients/{client_id}/social-providers/{provider}", di.NewUpsertSocialProviderHandler().Execute)
	})
}
